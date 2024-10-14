module github.com/nlepage/catption

go 1.18

require (
	github.com/gorilla/mux v1.7.4
	github.com/nlepage/catption/lib v0.0.0-00010101000000-000000000000
	github.com/nlepage/go-wasm-http-server/v2 v2.0.2
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
)

require (
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gobuffalo/here v0.6.2 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/hack-pad/safejs v0.1.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/markbates/pkger v0.17.0 // indirect
	github.com/mitchellh/mapstructure v1.3.2 // indirect
	github.com/nlepage/catption/font/impact v0.0.0-00010101000000-000000000000 // indirect
	github.com/nlepage/go-js-promise v1.0.0 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	golang.org/x/image v0.0.0-20200609002522-3f4726a040e8 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/ini.v1 v1.57.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/nlepage/catption/font/impact => ./font/impact
	github.com/nlepage/catption/lib => ./lib
)
