module github.com/sudhiaithal/protopatch

go 1.13

replace golang.org/x/tools => golang.org/x/tools v0.0.0-20210106214847-113979e3529a

replace golang.org/x/crypto v0.12.0 => golang.org/x/crypto v0.0.0-20181015023909-0c41d7ab0a0e

replace golang.org/x/net => golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20190429190828-d89cdac9e872

replace golang.org/x/sync v0.3.0 => golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f


require (
	github.com/envoyproxy/protoc-gen-validate v1.0.2
	github.com/fatih/structtag v1.2.0
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/stretchr/testify v1.8.4
	golang.org/x/tools v0.15.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/lyft/protoc-gen-star/v2 v2.0.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
