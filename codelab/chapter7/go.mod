module github.com/nlepage/catption/codelab/chapter7

go 1.13

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gobuffalo/here v0.6.2 // indirect
	github.com/markbates/pkger v0.17.0 // indirect
	github.com/mitchellh/mapstructure v1.3.2 // indirect
	github.com/nlepage/catption/lib v0.0.0-00010101000000-000000000000
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	golang.org/x/sys v0.0.0-20200610111108-226ff32320da // indirect
	gopkg.in/ini.v1 v1.57.0 // indirect
)

replace (
	github.com/nlepage/catption/font/impact => ../../font/impact
	github.com/nlepage/catption/lib => ../../lib
)
