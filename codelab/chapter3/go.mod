module github.com/nlepage/catption/codelab/chapter3

go 1.13

require (
	github.com/gobuffalo/here v0.6.2 // indirect
	github.com/markbates/pkger v0.17.0 // indirect
	github.com/nlepage/catption/lib v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5 // indirect
)

replace (
	github.com/nlepage/catption/font/impact => ../../font/impact
	github.com/nlepage/catption/lib => ../../lib
)
