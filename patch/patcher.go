package patch

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/structtag"
	"github.com/sudhiaithal/protopatch/patch/gopb"
	"golang.org/x/tools/go/ast/astutil"
	"google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/sudhiaithal/protopatch/lint"
	"github.com/sudhiaithal/protopatch/patch/ident"
)

// Patcher patches a set of generated Go Protobuf files with additional features:
// - (go.message).name overrides the name of a message’s synthesized struct.
// - (go.field).name overrides the name of a synthesized struct field and getters.
// - (go.field).tags lets you add additional struct tags to a field.
// - (go.oneof).name overrides the name of a oneof field, including wrapper types and getters.
// - (go.oneof).tags lets you specify additional struct tags on a oneof field.
// - (go.enum).name overrides the name of an enum type.
// - (go.value).name overrides the name of an enum value.
type Patcher struct {
	gen               *protogen.Plugin
	fset              *token.FileSet
	filesByName       map[string]*ast.File
	info              *types.Info
	packages          []*Package
	packagesByPath    map[string]*Package
	packagesByName    map[string]*Package
	renames           map[protogen.GoIdent]string
	typeRenames       map[protogen.GoIdent]string
	valueRenames      map[protogen.GoIdent]string
	fieldRenames      map[protogen.GoIdent]string
	methodRenames     map[protogen.GoIdent]string
	objectRenames     map[types.Object]string
	tags              map[protogen.GoIdent]string
	fieldTags         map[types.Object]string
	embeds            map[protogen.GoIdent]string
	nonNullables      map[protogen.GoIdent]string
	parentMap         map[protogen.GoIdent]string
	fieldEmbeds       map[types.Object]string
	fieldNonNullables map[types.Object]string
	fieldParentMap    map[types.Object]string
	types             map[protogen.GoIdent]string
	fieldTypes        map[types.Object]string
	processedMessages map[protogen.GoIdent]bool
}

// NewPatcher returns an initialized Patcher for gen.
func NewPatcher(gen *protogen.Plugin) (*Patcher, error) {
	p := &Patcher{
		gen:               gen,
		packagesByPath:    make(map[string]*Package),
		packagesByName:    make(map[string]*Package),
		renames:           make(map[protogen.GoIdent]string),
		typeRenames:       make(map[protogen.GoIdent]string),
		valueRenames:      make(map[protogen.GoIdent]string),
		fieldRenames:      make(map[protogen.GoIdent]string),
		methodRenames:     make(map[protogen.GoIdent]string),
		objectRenames:     make(map[types.Object]string),
		tags:              make(map[protogen.GoIdent]string),
		fieldTags:         make(map[types.Object]string),
		embeds:            make(map[protogen.GoIdent]string),
		nonNullables:      make(map[protogen.GoIdent]string),
		fieldEmbeds:       make(map[types.Object]string),
		parentMap:         make(map[protogen.GoIdent]string),
		fieldParentMap:    make(map[types.Object]string),
		fieldNonNullables: make(map[types.Object]string),
		types:             make(map[protogen.GoIdent]string),
		fieldTypes:        make(map[types.Object]string),
		processedMessages: make(map[protogen.GoIdent]bool),
	}
	return p, p.scan()
}

func getExtensionDesc(pb proto.Message, extname string) (*proto.ExtensionDesc, error) {
	desc := proto.RegisteredExtensions(pb)
	for _, d := range desc {
		if d.Name == extname {
			return d, nil
		}
	}
	return nil, fmt.Errorf("ExtensionDesc not found")
}

func getExtension(pb proto.Message, extname string) (interface{}, error) {
	d, err := getExtensionDesc(pb, extname)
	if err != nil {
		return nil, err
	}
	e, err := proto.GetExtension(pb, d)
	if err != nil {
		return nil, err
	}
	return e, err
}

func (p *Patcher) scan() error {
	for _, f := range p.gen.Files {
		p.scanFile(f)
	}

	for _, f := range p.gen.Request.ProtoFile {
		found := false
		mident := protogen.GoIdent{GoName: "", GoImportPath: ""}
		fident := protogen.GoIdent{GoName: "", GoImportPath: ""}
		for _, genFile := range p.gen.Files {
			if *f.Name == genFile.Desc.Path() {
				found = true
				mident = protogen.GoIdent{GoName: "", GoImportPath: genFile.GoImportPath}
				fident = protogen.GoIdent{GoName: "", GoImportPath: genFile.GoImportPath}
				break
			}
		}
		if !found {
			panic("Not found")
		}
		for _, m := range f.MessageType {
			mident.GoName = *m.Name
			if _, ok := p.processedMessages[mident]; ok {
				continue
			}
			for _, msgfield := range m.Field {
				fident.GoName = *msgfield.Name
				p.scanProtoField(mident, fident, msgfield)
			}
		}
	}

	return nil
}

func (p *Patcher) scanProtoField(mident protogen.GoIdent, fident protogen.GoIdent, f *descriptorpb.FieldDescriptorProto) {
	//m := f.Parent
	//o := f.Oneof

	if f.TypeName == nil {
		return
	}
	fi, err := getExtension(f.GetOptions(), "go.field")
	if err != nil {
		return
	}
	opts := fi.(*gopb.Options)

	log.Printf("Parent Message %v (%v), opts %v", *f.Name, *f.TypeName, opts)
	// Embed field ?
	embed := false
	newName := ""
	if opts.GetEmbed() {
		switch {
		//case f.Message == nil:
		//TODO
		//log.Printf("Warning: embed declared for non-message field: %s", f.Desc.Name())
		//case f.Oneof != nil:
		//log.Printf("Warning: embed declared for oneof field: %s", f.Desc.Name())
		default:
			embed = true
			log.Printf("Embed Set %v ", *f.Name, *f.TypeName)
			splitStrings := strings.Split((*f.TypeName)[1:], ".")
			newName = splitStrings[len(splitStrings)-1]
			// use the embed field message type's go name or rename option if defined
			/*if mOpts := messageOptions(f.Message); mOpts.GetName() != "" {
				newName = mOpts.GetName()
			} else {
				newName = f.Message.GoIdent.GoName
			}*/
		}
	}
	// Nullable  field ?
	nullable := true
	if opts != nil && opts.Nullable != nil {
		switch {
		//case f.Message == nil:
		//log.Printf("Warning: nullable declared for non-message field: %s", f.Desc.Name())
		//case f.Oneof != nil:
		//log.Printf("Warning: nullable declared for oneof field: %s", f.Desc.Name())
		default:
			if !*opts.Nullable {
				nullable = false
			}
		}
	}
	if newName != "" {
		//if o != nil {
		if false {
			p.RenameType(fident, p.nameFor(mident)+"_"+newName)                          // Oneof wrapper struct
			p.RenameField(ident.WithChild(fident, fident.GoName), newName, false, false) // Oneof wrapper field (not embeddable)
			//ifName := ident.WithPrefix(o.GoIdent, "is")
			//p.RenameMethod(ident.WithChild(f.GoIdent, ifName.GoName), p.nameFor(ifName)) // Oneof interface method
			//childID := ident.WithChild(f.GoIdent, ifName.GoName)
			//p.parentMap[childID] = string(m.Desc.Name())
		} else {
			p.RenameField(ident.WithChild(mident, fident.GoName), newName, embed, nullable) // Field
			childID := ident.WithChild(mident, fident.GoName)
			p.parentMap[childID] = string(mident.GoName)
			log.Printf("child %v parent %v", childID, mident.GoName)
			//childID := ident.WithChild(mident, fident.GoName)
			//p.parentMap[childID] = string(m.Desc.Name())
		}
		p.RenameMethod(ident.WithChild(mident, "Get"+fident.GoName), "Get"+newName) // Getter
	} else {
		p.RenameField(ident.WithChild(mident, fident.GoName), newName, embed, nullable) // Field
		childID := ident.WithChild(mident, fident.GoName)
		p.parentMap[childID] = string(mident.GoName)
	}
}

func (p *Patcher) scanFile(f *protogen.File) {
	log.Printf("\nScan proto:\t%s", f.Desc.Path())

	// Locally generate Go from the source proto file.
	// This is equivalent to running the go protoc plugin, but in-process.
	if f.Generate {
		log.Printf("Generating:\t%s", f.Desc.Path())
		internal_gengo.GenerateFile(p.gen, f)
	}

	_ = p.getPackage(string(f.GoImportPath), string(f.GoPackageName), true)

	for _, e := range f.Enums {
		p.scanEnum(e, nil)
	}

	for _, m := range f.Messages {
		p.scanMessage(m, nil)
	}

	for _, e := range f.Extensions {
		p.scanExtension(e)
	}

	// TODO: scan gRPC services
}

func (p *Patcher) scanEnum(e *protogen.Enum, parent *protogen.Message) {
	opts := enumOptions(e)
	lints := fileLintOptions(e.Desc)

	// Rename enum?
	newName := opts.GetName()
	if newName == "" && parent != nil && p.isRenamed(parent.GoIdent) {
		newName = replacePrefix(e.GoIdent.GoName, parent.GoIdent.GoName, p.nameFor(parent.GoIdent))
		log.Printf("•••• %s → newName: %s", e.GoIdent.GoName, newName)
	}
	if lints.GetEnums() || lints.GetAll() {
		if newName == "" {
			newName = e.GoIdent.GoName
		}
		newName = lint.Name(newName, lints.InitialismsMap())
	}
	if newName != "" {
		p.RenameType(e.GoIdent, newName)                                       // Enum type
		p.RenameValue(ident.WithSuffix(e.GoIdent, "_name"), newName+"_name")   // Enum name map
		p.RenameValue(ident.WithSuffix(e.GoIdent, "_value"), newName+"_value") // Enum value map
	}

	// Rename String method?
	newStringer := opts.GetStringer()
	// TODO: remove StringerName in two minor versions (~0.3.0)
	if newStringer == "" {
		newStringer = opts.GetStringerName()
		if newStringer != "" {
			log.Printf("Warning: stringer_name is deprecated and will be removed in a future version. Please use stringer.")
		}
	}
	if newStringer != "" {
		p.RenameMethod(ident.WithChild(e.GoIdent, "String"), newStringer)
	}

	// Scan enum values.
	for _, v := range e.Values {
		p.scanEnumValue(v, parent)
	}
}

func (p *Patcher) scanEnumValue(v *protogen.EnumValue, parent *protogen.Message) {
	// Enum values are prefixed with the parent *message* name if it exists.
	// https://github.com/protocolbuffers/protobuf-go/blob/160c7477e0e899d5072bb25635f46053df619fbf/compiler/protogen/protogen.go#L640-L643
	parentIdent := v.Parent.GoIdent
	if parent != nil {
		parentIdent = parent.GoIdent
	}
	opts := valueOptions(v)
	lints := fileLintOptions(v.Desc)

	// Rename enum value?
	newName := opts.GetName()
	if newName == "" {
		newName = replacePrefix(v.GoIdent.GoName, parentIdent.GoName, p.nameFor(parentIdent))
	}
	if lints.GetValues() || lints.GetAll() {
		vname := string(v.Desc.Name())
		if vname == strings.ToUpper(vname) && strings.HasSuffix(newName, vname) {
			newName = strings.TrimSuffix(newName, vname) + "_" + strings.ToLower(vname)
		}

		newName = lint.Name(newName, lints.InitialismsMap())

		// Remove type name prefix stutter, e.g. FooFooUnknown → FooUnknown
		pname := p.nameFor(parentIdent)
		pfx := pname + pname
		if len(newName) > len(pfx) && strings.HasPrefix(newName, pfx) {
			newName = strings.TrimPrefix(newName, pname)
		}
	}
	if newName != "" {
		p.RenameValue(v.GoIdent, newName)
	}
}

func (p *Patcher) scanMessage(m *protogen.Message, parent *protogen.Message) {
	opts := messageOptions(m)
	lints := fileLintOptions(m.Desc)

	p.processedMessages[m.GoIdent] = true
	// Rename message?
	newName := opts.GetName()
	if newName == "" && parent != nil && p.isRenamed(parent.GoIdent) {
		newName = replacePrefix(m.GoIdent.GoName, parent.GoIdent.GoName, p.nameFor(parent.GoIdent))
	}
	if lints.GetMessages() || lints.GetAll() {
		log.Printf("Linting: %q.%s", m.GoIdent.GoImportPath, m.GoIdent.GoName)
		if newName == "" {
			newName = m.GoIdent.GoName
		}
		newName = lint.Name(newName, lints.InitialismsMap())
	}
	if newName != "" {
		p.RenameType(m.GoIdent, newName) // Message struct
	}

	// Scan message oneof fields.
	for _, o := range m.Oneofs {
		p.scanOneof(o)
	}

	// Scan message fields.
	for _, f := range m.Fields {
		p.scanField(f)
	}

	// Scan nested enums.
	for _, e := range m.Enums {
		p.scanEnum(e, m)
	}

	// Scan nested messages.
	for _, mm := range m.Messages {
		p.scanMessage(mm, m)
	}
}

func replacePrefix(s, prefix, with string) string {
	if !strings.HasPrefix(s, prefix) {
		return s
	}
	return with + strings.TrimPrefix(s, prefix)
}

func (p *Patcher) scanOneof(o *protogen.Oneof) {
	m := o.Parent
	opts := oneofOptions(o)
	lints := fileLintOptions(o.Desc)

	// Rename oneof field?
	newName := opts.GetName()
	if newName == "" && p.isRenamed(m.GoIdent) {
		// Implicitly rename this oneof field because its parent message was renamed.
		newName = o.GoName
	}
	if lints.GetFields() || lints.GetAll() {
		if newName == "" {
			newName = o.GoIdent.GoName
		}
		newName = lint.Name(newName, lints.InitialismsMap())
	}
	if newName != "" {
		p.RenameField(ident.WithChild(m.GoIdent, o.GoName), newName, false, false) // Oneof
		p.parentMap[ident.WithChild(m.GoIdent, o.GoName)] = string(m.Desc.Name())
		p.RenameMethod(ident.WithChild(m.GoIdent, "Get"+o.GoName), "Get"+newName) // Getter
		ifName := ident.WithPrefix(o.GoIdent, "is")
		newIfName := "is" + p.nameFor(m.GoIdent) + "_" + newName
		p.RenameType(ifName, newIfName)                                   // Interface type (e.g. isExample_Person)
		p.RenameMethod(ident.WithChild(ifName, ifName.GoName), newIfName) // Interface method
	}

	// Add or replace any struct tags?
	tags := opts.GetTags()
	if tags != "" {
		p.Tag(ident.WithChild(m.GoIdent, o.GoName), tags)
	}
}

func (p *Patcher) scanField(f *protogen.Field) {
	m := f.Parent
	o := f.Oneof
	if f.Desc.HasOptionalKeyword() {
		o = nil
	}
	opts := fieldOptions(f)
	lints := fileLintOptions(f.Desc)

	// Rename message field?
	newName := opts.GetName()
	if newName == "" && o != nil && (p.isRenamed(m.GoIdent) || p.isRenamed(o.GoIdent)) {
		// Implicitly rename this oneof field because its parent(s) were renamed.
		newName = f.GoName
	}
	// Embed field ?
	embed := false
	if opts.GetEmbed() {
		switch {
		case f.Message == nil:
			//TODO
			log.Printf("Warning: embed declared for non-message field: %s", f.Desc.Name())
		case f.Oneof != nil:
			log.Printf("Warning: embed declared for oneof field: %s", f.Desc.Name())
		default:
			embed = true
			// use the embed field message type's go name or rename option if defined
			if mOpts := messageOptions(f.Message); mOpts.GetName() != "" {
				newName = mOpts.GetName()
			} else {
				newName = f.Message.GoIdent.GoName
			}
		}
	}
	// Nullable  field ?
	nullable := true
	if opts != nil && opts.Nullable != nil {
		switch {
		case f.Message == nil:
			log.Printf("Warning: nullable declared for non-message field: %s", f.Desc.Name())
		case f.Oneof != nil:
			log.Printf("Warning: nullable declared for oneof field: %s", f.Desc.Name())
		default:
			if !*opts.Nullable {
				nullable = false
			}
		}
	}
	if lints.GetFields() || lints.GetAll() {
		if newName == "" {
			newName = f.GoName
		}
		newName = lint.Name(newName, lints.InitialismsMap())
	}
	log.Printf("Child %s, parent %v", f.Desc.Name(), m.Desc.Name())
	if newName != "" {
		if o != nil {
			p.RenameType(f.GoIdent, p.nameFor(m.GoIdent)+"_"+newName)                  // Oneof wrapper struct
			p.RenameField(ident.WithChild(f.GoIdent, f.GoName), newName, false, false) // Oneof wrapper field (not embeddable)
			ifName := ident.WithPrefix(o.GoIdent, "is")
			p.RenameMethod(ident.WithChild(f.GoIdent, ifName.GoName), p.nameFor(ifName)) // Oneof interface method
			childID := ident.WithChild(f.GoIdent, ifName.GoName)
			p.parentMap[childID] = string(m.Desc.Name())
		} else {
			p.RenameField(ident.WithChild(m.GoIdent, f.GoName), newName, embed, nullable) // Field
			childID := ident.WithChild(m.GoIdent, f.GoName)
			p.parentMap[childID] = string(m.Desc.Name())
		}
		p.RenameMethod(ident.WithChild(m.GoIdent, "Get"+f.GoName), "Get"+newName) // Getter
	} else {
		p.RenameField(ident.WithChild(m.GoIdent, f.GoName), newName, embed, nullable) // Field
		childID := ident.WithChild(m.GoIdent, f.GoName)
		p.parentMap[childID] = string(m.Desc.Name())
	}

	// check type
	if fieldType := opts.GetType(); fieldType != "" {
		switch {
		case f.Message != nil && !f.Desc.IsList():
			log.Printf("Warning: type declared for message field: %s", f.Desc.Name())
		case f.Oneof != nil && !f.Desc.HasOptionalKeyword():
			p.Type(ident.WithChild(f.GoIdent, f.GoName), fieldType)
			p.Type(ident.WithChild(m.GoIdent, "Get"+f.GoName), fieldType)
		default:
			p.Type(ident.WithChild(m.GoIdent, f.GoName), fieldType)
			p.Type(ident.WithChild(m.GoIdent, "Get"+f.GoName), fieldType)
		}
	}

	// Add or replace any struct tags?
	tags := opts.GetTags()
	if tags != "" {
		if o != nil {
			p.Tag(ident.WithChild(f.GoIdent, f.GoName), tags) // Oneof wrapper field tags
		} else {
			p.Tag(ident.WithChild(m.GoIdent, f.GoName), tags) // Field tags
		}
	}
}

func (p *Patcher) scanExtension(f *protogen.Field) {
	opts := fieldOptions(f)
	lints := fileLintOptions(f.Desc)

	// Rename extension?
	newName := opts.GetName()
	if lints.GetExtensions() || lints.GetAll() {
		if newName == "" {
			// Idiomatic Go values are prefixed with some flavor of the type, in this case Ext.
			newName = "Ext" + f.GoName
		}
		newName = lint.Name(newName, lints.InitialismsMap())
	}
	if newName != "" {
		id := f.GoIdent
		id.GoName = "E_" + f.GoName
		p.RenameValue(id, newName)
	}
}

// RenameType renames the Go type specified by id to newName.
// The id argument specifies a GoName from GoImportPath, e.g.: "github.com/org/repo/example".FooMessage
// To rename a package-level identifier such as a type, var, or const, specify just the name, e.g. "Message" or "Enum_VALUE".
// newName should be the unqualified name.
// The value of id.GoName should be the original generated type name, not a renamed type.
func (p *Patcher) RenameType(id protogen.GoIdent, newName string) {
	p.renames[id] = newName
	p.typeRenames[id] = newName
	log.Printf("Rename type:\t%s.%s → %s", id.GoImportPath, id.GoName, newName)
}

// RenameValue renames the Go value (const or var) specified by id to newName.
// The id argument specifies a GoName from GoImportPath, e.g.: "github.com/org/repo/example".FooValue
// newName should be the unqualified name.
// The value of id.GoName should be the original generated type name, not a renamed type.
func (p *Patcher) RenameValue(id protogen.GoIdent, newName string) {
	p.renames[id] = newName
	p.valueRenames[id] = newName
	log.Printf("Rename value:\t%s.%s → %s", id.GoImportPath, id.GoName, newName)
}

// RenameField renames the Go struct field specified by id to newName.
// The id argument specifies a GoName from GoImportPath, e.g.: "github.com/org/repo/example".FooMessage.BarField
// newName should be the unqualified name (after the dot).
// The value of id.GoName should be the original generated identifier name, not a renamed identifier.
func (p *Patcher) RenameField(id protogen.GoIdent, newName string, embed, nullable bool) {
	p.renames[id] = newName
	p.fieldRenames[id] = newName
	if embed {
		p.embeds[id] = newName
	}
	if !nullable {
		p.nonNullables[id] = newName
	}
	log.Printf("Rename field:\t%s.%s → %s", id.GoImportPath, id.GoName, newName)
}

// RenameMethod renames the Go struct or interface method specified by id to newName.
// The id argument specifies a GoName from GoImportPath, e.g.: "github.com/org/repo/example".FooMessage.GetBarField
// The new name should be the unqualified name (after the dot).
// The value of id.GoName should be the original generated identifier name, not a renamed identifier.
func (p *Patcher) RenameMethod(id protogen.GoIdent, newName string) {
	p.renames[id] = newName
	p.methodRenames[id] = newName
	log.Printf("Rename method:\t%s.%s → %s", id.GoImportPath, id.GoName, newName)
}

func (p *Patcher) isRenamed(id protogen.GoIdent) bool {
	_, ok := p.renames[id]
	return ok
}

func (p *Patcher) nameFor(id protogen.GoIdent) string {
	if name, ok := p.renames[id]; ok {
		return name
	}
	return ident.LeafName(id)
}

// Type casts the Go struct field as the desired type
// The typeName value must be a named type, e.g.: "type String string"
func (p *Patcher) Type(id protogen.GoIdent, typeName string) {
	if isTypeValid(typeName) {
		log.Printf("Warning: field %s has invalid type option: %s", id.GoImportPath, typeName)
		return
	}
	p.types[id] = typeName
	log.Printf("Cast type:\t%s.%s → %s", id.GoImportPath, id.GoName, typeName)
}

// Tag adds the specified struct tags to the field specified by selector,
// in the form of "Message.Field". The tags argument should omit outer backticks (`).
// The value of id.GoName should be the original generated identifier name, not a renamed identifier.
// The struct tags will be applied when Patch is called.
func (p *Patcher) Tag(id protogen.GoIdent, tags string) {
	p.tags[id] = tags
	log.Printf("Tags:\t%s.%s `%s`", id.GoImportPath, id.GoName, tags)
}

// Patch applies the patch(es) in p the Go files in res.
// Clone res before calling Patch if you want to retain an unmodified copy.
// The behavior of calling Patch multiple times is currently undefined.
func (p *Patcher) Patch(res *pluginpb.CodeGeneratorResponse) error {
	p.reset()

	if err := p.parseGoFiles(res); err != nil {
		return err
	}

	// Inject default generated Go code from protoc-gen-go.
	// FIXME: should this be here?
	res2 := p.gen.Response()
	if err := p.parseGoFiles(res2); err != nil {
		return err
	}

	if err := p.checkGoFiles(); err != nil {
		return err
	}

	if err := p.patchGoFiles(); err != nil {
		return err
	}

	return p.serializeGoFiles(res)
}

func (p *Patcher) reset() {
	p.fset = token.NewFileSet()
	p.filesByName = make(map[string]*ast.File)
}

func (p *Patcher) parseGoFiles(res *pluginpb.CodeGeneratorResponse) error {
	for _, rf := range res.File {
		if rf.Name == nil || !strings.HasSuffix(*rf.Name, ".go") || rf.Content == nil {
			continue
		}

		if p.filesByName[*rf.Name] != nil {
			log.Printf("Skipping duplicate file:\t%s", *rf.Name)
			continue
		}

		f, err := p.parseGoFile(*rf.Name, *rf.Content)
		if err != nil {
			return err
		}

		// TODO: should we cache these?
		p.filesByName[*rf.Name] = f

		// FIXME: this will break if the package’s implicit name differs from the types.Package name.
		if pkg, ok := p.packagesByName[f.Name.Name]; ok {
			pkg.AddFile(*rf.Name, f)
		} else {
			return fmt.Errorf("unknown package: %s", f.Name.Name)
		}
	}

	return nil
}

func (p *Patcher) checkGoFiles() error {
	// Type-check Go packages first to find any missing identifiers.
	if err := p.checkPackages(); err != nil {
		return err
	}

	var recheck bool

	// Find missing type declarations.
	for id := range p.typeRenames {
		if obj, _ := p.find(id); obj != nil {
			continue
		}
		if err := p.synthesize(id); err != nil {
			return err
		}
		recheck = true
	}

	// Find missing field declarations.
	for id := range p.fieldRenames {
		if obj, _ := p.find(id); obj != nil {
			continue
		}
		if err := p.synthesize(id); err != nil {
			return err
		}
		recheck = true
	}

	// Find missing method declarations.
	for id := range p.methodRenames {
		if obj, _ := p.find(id); obj != nil {
			continue
		}
		if err := p.synthesize(id); err != nil {
			return err
		}
		recheck = true
	}

	// Find missing value declarations.
	for id := range p.valueRenames {
		if obj, _ := p.find(id); obj != nil {
			continue
		}
		if err := p.synthesize(id); err != nil {
			return err
		}
		recheck = true
	}

	// Re-type-check if necessary.
	if recheck {
		if err := p.checkPackages(); err != nil {
			return err
		}
	}

	// Map renames.
	for id, name := range p.renames {
		obj, _ := p.find(id)
		if obj == nil {
			log.Printf("Did not find the object for id %+v:%v %v", id.GoImportPath, id.GoName, name)
			continue
		}
		p.objectRenames[obj] = name
		if _, ok := p.parentMap[id]; ok {
			//splits := strings.Split(obj.Type().String(), ".")
			parentNames := strings.Split(id.GoName, ".")
			if len(parentNames) == 2 {
				//parentNames = parentNames[0:len(parentNames)-1]
				p.fieldParentMap[obj] = parentNames[0]
			} else {
				p.fieldParentMap[obj] = p.parentMap[id]
			}
			log.Printf("Adding parent reference (%v) %v fo field  %v parents %v  TS %v Parent %v ", id, p.fieldParentMap[obj], obj, parentNames, obj.Type().String(), p.fieldParentMap[obj])

			//p.fieldParentMap[obj] = p.parentMap[id]
			//log.Printf("Adding parent reference %v fo field  %v", p.parentMap[id], obj)
		} else {
			log.Printf("Did not find parent map for id %v:%v", id, obj)
		}
		if _, ok := p.embeds[id]; ok {
			p.fieldEmbeds[obj] = name
		}
		if _, ok := p.nonNullables[id]; ok {
			p.fieldNonNullables[obj] = name
		}
	}

	for id, name := range p.nonNullables {
		if _, ok := p.renames[id]; ok {
			continue
		}
		obj, _ := p.find(id)
		if obj == nil {
			continue
		}
		p.objectRenames[obj] = name
		if _, ok := p.parentMap[id]; ok {
			splits := strings.Split(obj.Type().String(), ".")
			parentNames := strings.Split(splits[len(splits)-1], "_")
			if len(parentNames) > 2 {
				parentNames = parentNames[0 : len(parentNames)-1]
				p.fieldParentMap[obj] = strings.Join(parentNames, "_")
			} else {
				p.fieldParentMap[obj] = p.parentMap[id]
			}
			log.Printf("Adding Non nullable parent reference %v fo field  %v", p.parentMap[id], obj)
		}
		if _, ok := p.embeds[id]; ok {
			p.fieldEmbeds[obj] = name
		}
		if _, ok := p.nonNullables[id]; ok {
			p.fieldNonNullables[obj] = name
		}
	}

	// Map cast types
	for id, typ := range p.types {
		obj, _ := p.find(id)
		if obj == nil {
			continue
		}
		p.fieldTypes[obj] = typ
	}

	// Map struct tags.
	for id, tags := range p.tags {
		obj, _ := p.find(id)
		if obj == nil {
			continue
		}
		p.fieldTags[obj] = tags
	}

	return nil
}

func (p *Patcher) parseGoFile(filename string, src interface{}) (*ast.File, error) {
	f, err := parser.ParseFile(p.fset, filename, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	log.Printf("\nParse Go:\t%s\n", filename)
	return f, nil
}

func (p *Patcher) checkPackages() error {
	p.info = &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	for _, pkg := range p.packagesByName {
		pkg.Reset()
	}

	// Iterate packages deterministically
	for _, pkg := range p.packages {
		if len(pkg.files) == 0 {
			continue
		}
		// Resolve symbols defined in this package across all files
		_, _ = ast.NewPackage(p.fset, pkg.filesByName, nil, nil)
		err := pkg.Check(basicImporter{p}, p.fset, p.info)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Patcher) synthesize(id protogen.GoIdent) error {
	pkg := p.getPackage(string(id.GoImportPath), id.GoName, true)

	// Already synthesized?
	filename := pkg.pkg.Path() + "/" + id.GoName + ".synthetic.go"
	if f := pkg.File(filename); f != nil {
		return nil
	}

	// Synthesize a Go file for this identifier.
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "package %s\n\n", pkg.pkg.Name())
	names := strings.Split(id.GoName, ".")
	if len(names) == 1 {
		// Type or value.
		// Synthesize a Go type as a map so subscript access works, e.g.: foo[key]
		fmt.Fprintf(b, "type %s map[interface{}]interface{}\n", names[0])
	} else {
		// Field or method.
		// Synthesize a Go method so a non-call expr works, e.g.: foo.Method
		fmt.Fprintf(b, "func (%s) %s() {}\n", names[0], names[1])
	}
	log.Printf("\nGenerated Go code: %s\n\n%s\n", filename, b.String())

	// Parse and add it to pkg.
	f, err := p.parseGoFile(filename, b)
	if err != nil {
		return err
	}
	return pkg.AddFile(filename, f)
}

// find finds id in all parsed Go packages, along with any ancestor(s),
// or nil if the id is not found.
func (p *Patcher) find(id protogen.GoIdent) (obj types.Object, ancestors []types.Object) {
	pkg := p.getPackage(string(id.GoImportPath), "", true)
	if pkg == nil {
		return
	}
	return pkg.Find(id)
}

// getPackage finds a getPackage with path, or creates it if it doesn’t exist.
// If name is empty, getPackage will use the last path element as the package name.
func (p *Patcher) getPackage(path, name string, create bool) *Package {
	if pkg, ok := p.packagesByPath[path]; ok {
		return pkg
	}
	if !create {
		return nil
	}
	if name == "" {
		name = filepath.Base(path)
	}
	pkg := NewPackage(path, name)
	name = pkg.pkg.Name() // Get real name
	p.packagesByPath[path] = pkg
	p.packagesByName[name] = pkg
	p.packages = append(p.packages, pkg)
	return pkg
}

func (p *Patcher) serializeGoFiles(res *pluginpb.CodeGeneratorResponse) error {
	for _, rf := range res.File {
		if rf.Name == nil || !strings.HasSuffix(*rf.Name, ".go") || rf.Content == nil {
			continue
		}
		log.Printf("\nSerialize:\t%s\n", *rf.Name)

		f := p.filesByName[*rf.Name]
		if f == nil {
			continue // Should never happen
		}

		var b strings.Builder
		err := format.Node(&b, p.fset, f)
		if err != nil {
			return err
		}

		content := b.String()
		rf.Content = &content
	}
	return nil
}

type funcDecls struct {
	decls     []*ast.FuncDecl
	filename  string
	className string
}

func (p *Patcher) getFuncDecls(name string) []funcDecls {

	funcDeclsSlice := []funcDecls{}
	for fileName, f := range p.filesByName {
		var decls []*ast.FuncDecl
		for _, decl := range f.Decls {
			switch decl.(type) {
			case *ast.FuncDecl:
				if decl.(*ast.FuncDecl).Name.Name == name {
					decls = append(decls, decl.(*ast.FuncDecl))
				}
			}
			funcDeclsSlice = append(funcDeclsSlice, funcDecls{decls: decls, filename: fileName})
		}
	}
	return funcDeclsSlice

}

func (p *Patcher) addFuncDecls(fileName string, addDecl *ast.FuncDecl) {

	f := p.filesByName[fileName]
	for _, decl := range f.Decls {
		switch decl.(type) {
		case *ast.FuncDecl:
			if decl.(*ast.FuncDecl).Name.Name == addDecl.Name.Name &&
				decl.(*ast.FuncDecl).Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name == addDecl.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name {
				return
			}
		}
	}
	f.Decls = append(f.Decls, addDecl)
}

func getPointerFunction(className string) *ast.FuncDecl {

	funcDecl := &ast.FuncDecl{
		Name: ast.NewIdent("Pointer"),
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{ast.NewIdent("x")},
					Type: &ast.StarExpr{
						X: ast.NewIdent(className),
					},
				},
			},
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("data interface{}"),
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("interface{}"),
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{ast.NewIdent("a")},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{ast.NewIdent(fmt.Sprintf("data.(%s)", className))},
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						ast.NewIdent("&a"),
					},
				},
			},
		},
	}

	return funcDecl
}

func (p *Patcher) patchGoFiles() error {
	log.Printf("\nDefs")
	for id, obj := range p.info.Defs {
		p.patchTypeDef(id, obj)
		p.patchIdent(id, obj, true)
		p.patchTags(id, obj)
		// if id.IsExported() {
		// 	f := p.fset.File(id.NamePos)
		// 	log.Printf("Ident %s:\t%s @ %s", typeString(obj), id.Name, f.Name())
		// }
	}

	log.Printf("\nUses\n")
	for id, obj := range p.info.Uses {
		p.patchTypeUsage(id, obj)
		p.patchIdent(id, obj, false)
	}

	log.Printf("\nUnresolved\n")
	for fileName, f := range p.filesByName {
		for _, id := range f.Unresolved {
			p.patchIdent(id, nil, false)
		}

		ast.Inspect(f, func(n ast.Node) bool {
			// try to convert n to ast.TypeSpec
			if typeSpec, isTypeSpec := n.(*ast.TypeSpec); isTypeSpec {
				_, isStructType := typeSpec.Type.(*ast.StructType)
				// check if conversion was successful
				if !isStructType {
					return true
				}
				log.Printf("Struct found %v\n", typeSpec.Name.Name)
				if typeSpec.Name.Name != "x" {
					p.addFuncDecls(fileName, getPointerFunction(typeSpec.Name.Name))
				}

			}
			return true
		})
	}

	return nil
}

func (p *Patcher) patchIdent(id *ast.Ident, obj types.Object, isDecl bool) {

	name := p.objectRenames[obj]
	setNonNullable := func(name string) {
		if _, ok := p.fieldNonNullables[obj]; ok && isDecl {
			log.Printf("Renamed %s: → %s (non-nullable) %v", typeString(obj), id.Name, id.Obj)
			switch t := id.Obj.Decl.(type) {
			case *ast.Field:
				log.Printf("Matched Renamed %s: → %s (non-nullable) %v", typeString(obj), id.Name, id.Obj)
				switch t1 := t.Type.(type) {
				case *ast.StarExpr:
					switch t2 := t1.X.(type) {
					case *ast.Ident:
						name = t2.Name
					case *ast.SelectorExpr:
						name = fmt.Sprintf("%v", t2.X) + "." + t2.Sel.Name
					default:
						log.Printf("Defaiult Renamed %s: → %s (non-nullable) %v %T", typeString(obj), id.Name, id.Obj, t2)
					}
					if name == "" {
						name = t.Type.(*ast.StarExpr).X.(*ast.Ident).Name
					}
					t.Type = &ast.Ident{
						Name: name}
				case *ast.ArrayType:
					switch t3 := t1.Elt.(type) {
					case *ast.StarExpr:
						switch t4 := t3.X.(type) {
						case *ast.Ident:
							name = t4.Name
						case *ast.SelectorExpr:
							name = fmt.Sprintf("%v", t4.X) + "." + t4.Sel.Name
						default:
							log.Printf("Array  Renamed %s: → %s (non-nullable) %v %T", typeString(obj), id.Name, id.Obj, t3)
						}
						if name == "" {
							name = t1.Elt.(*ast.StarExpr).X.(*ast.Ident).Name
						}
						//t1.Elt = &ast.Ident{
						//	Name: name}
						t.Type = &ast.ArrayType{
							Elt: &ast.Ident{Name: name}}
					}
					/*			case *ast.Field:
								switch t2 := t1.Type.(type) {
								default:
								log.Printf("Defaiult Renamed %s: → %s (non-nullable) %v %T", typeString(obj), id.Name, id.Obj, t2)
								} */
				}
			default:
				log.Printf("Defaiult Renamed %s: → %s (non-nullable) %v", typeString(obj), id.Name, id.Obj)

			}
		}
	}

	patchIdentGetFunc := func() {

		funcDecls := p.getFuncDecls(id.Name)
		if funcDecls == nil || len(funcDecls) == 0 {
			return
		}
		fieldName := strings.TrimPrefix(id.Name, "Get")
		for idx := range p.fieldNonNullables {
			_, ok := obj.(*types.Func)
			if !ok {
				continue
			}
			if idx.Name() == fieldName && idx.Pkg().Name() == obj.(*types.Func).Pkg().Name() {
				for _, funcDeclData := range funcDecls {
					for _, funcDecl := range funcDeclData.decls {
						className := ""
						for _, param := range funcDecl.Recv.List {
							className = fmt.Sprintf("%s", param.Type.(*ast.StarExpr).X)
						}
						if p.fieldParentMap[idx] != className {
							log.Printf("Skipping Function as class name %v does not match of function %v does not match to rename field's class name %v", className, obj.(*types.Func).Name(), p.fieldParentMap[idx])
							continue
						}
						for _, field := range funcDecl.Type.Results.List {
							switch t1 := field.Type.(type) {
							case *ast.StarExpr:

								log.Printf("1 Changed return for method %v from *%s : → %s (non-nullable)",
									idx.Name, t1.X, t1.X)
								switch t3 := t1.X.(type) {
								case *ast.Ident:
									field.Type = &ast.Ident{
										Name: t3.Name}
								case *ast.SelectorExpr:
									field.Type = &ast.Ident{
										Name: fmt.Sprintf("%v", t3.X) + "." + t3.Sel.Name}
								}
								for _, stmt := range funcDecl.Body.List {
									switch t2 := stmt.(type) {
									case *ast.ReturnStmt:
										for i := range t2.Results {
											t2.Results[i].(*ast.Ident).Name = fmt.Sprintf("%v{}", fmt.Sprintf("%v", field.Type))
										}
									}
								}
							case *ast.ArrayType:
								switch t3 := t1.Elt.(type) {
								case *ast.StarExpr:
									switch t4 := t3.X.(type) {
									case *ast.Ident:
										field.Type = &ast.ArrayType{
											Elt: &ast.Ident{Name: t4.Name}}
									case *ast.SelectorExpr:
										field.Type = &ast.ArrayType{
											Elt: &ast.Ident{Name: fmt.Sprintf("%v", t4.X) + "." + t4.Sel.Name}}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	patchIdentGetFunc()
	if name == "" {
		setNonNullable(name)
		return
	}
	p.patchComments(id, name)
	if _, ok := p.fieldEmbeds[obj]; ok && isDecl {
		log.Printf("Renamed %s:\t%s → %s (embedded)", typeString(obj), id.Name, name)
		id.Name = ""
		setNonNullable(name)
	} else {
		log.Printf("Renamed %s:\t%s → %s", typeString(obj), id.Name, name)
		id.Name = name
	}
}

func (p *Patcher) nodeToString(n ast.Node) string {
	b := &bytes.Buffer{}
	if err := printer.Fprint(b, p.fset, n); err != nil {
		log.Fatal(err)
	}
	return b.String()
}

func (p *Patcher) findParentNode(id ast.Node) ast.Node {
	var node ast.Node
	astutil.Apply(p.fileOf(id), func(cursor *astutil.Cursor) bool {
		if cursor.Node() == id {
			node = cursor.Parent()
			return false
		}
		return true
	}, nil)
	return node
}

func (p *Patcher) patchTags(id *ast.Ident, obj types.Object) {
	fieldTags := p.fieldTags[obj]
	if fieldTags == "" || id.Obj == nil {
		return
	}

	v, ok := id.Obj.Decl.(*ast.Field)
	if !ok {
		log.Printf("Warning: struct tags declared for non-field object: %v `%s`", obj, fieldTags)
		return
	}

	if v.Tag == nil {
		v.Tag = &ast.BasicLit{}
	}

	tags, err := structtag.Parse(strings.Trim(v.Tag.Value, "`"))
	if err != nil {
		log.Printf("Error: parsing struct tags for %q.%s: %s", obj.Pkg().Path(), id.Name, err)
		return
	}

	newTags, err := structtag.Parse(fieldTags)
	if err != nil {
		log.Printf("Error: parsing struct tags for %q.%s: %s", obj.Pkg().Path(), id.Name, err)
		return
	}

	for _, tag := range newTags.Tags() {
		tags.Set(tag)
	}

	v.Tag.Value = "`" + tags.String() + "`"

	log.Printf("Add tags:\t%q.%s `%s`", obj.Pkg().Path(), id.Name, newTags.String())
}

func (p *Patcher) patchComments(id *ast.Ident, repl string) {
	doc, comment := p.findCommentGroups(id)
	if doc == nil && comment == nil {
		return
	}
	x, err := regexp.Compile(`\b` + regexp.QuoteMeta(id.Name) + `\b`)
	if err != nil {
		return
	}
	log.Printf("Comment:\t%v → %s", x, repl)
	patchCommentGroup(doc, x, repl)
	patchCommentGroup(comment, x, repl)
}

// Borrowed from https://github.com/golang/tools/blob/HEAD/refactor/rename/rename.go#L543
func (p *Patcher) findCommentGroups(id *ast.Ident) (doc *ast.CommentGroup, comment *ast.CommentGroup) {
	f := p.fileOf(id)
	if f == nil {
		return
	}
	nodes, _ := astutil.PathEnclosingInterval(f, id.Pos(), id.End())
	for _, node := range nodes {
		switch decl := node.(type) {
		case *ast.FuncDecl:
			return decl.Doc, nil
		case *ast.Field:
			return decl.Doc, decl.Comment
		case *ast.GenDecl:
			return decl.Doc, nil
		// For {Type,Value}Spec, if the doc on the spec is absent, search for the enclosing GenDecl.
		case *ast.TypeSpec:
			if decl.Doc != nil {
				return decl.Doc, decl.Comment
			}
		case *ast.ValueSpec:
			if decl.Doc != nil {
				return decl.Doc, decl.Comment
			}
		case *ast.Ident:
		default:
			return
		}
	}
	return
}

func (p *Patcher) fileOf(node ast.Node) *ast.File {
	tf := p.fset.File(node.Pos())
	if tf == nil {
		return nil
	}
	return p.filesByName[tf.Name()]
}

func patchCommentGroup(c *ast.CommentGroup, x *regexp.Regexp, repl string) {
	if c == nil {
		return
	}
	for _, c := range c.List {
		c.Text = x.ReplaceAllString(c.Text, repl)
	}
}

func typeString(obj types.Object) string {
	switch obj.(type) {
	case *types.PkgName:
		return "package name"
	case *types.TypeName:
		return "type"
	case *types.Var:
		if obj.Parent() == nil {
			return "field"
		}
		return "var"
	case *types.Const:
		return "const"
	case *types.Func:
		if obj.Parent() == nil {
			return "method"
		}
		return "func"
	case nil:
		return "(nil)"
	}
	return obj.Type().String()
}
