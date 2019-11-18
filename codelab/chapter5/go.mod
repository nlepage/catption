module github.com/Zenika/catption/codelab/chapter5

go 1.12

require (
	github.com/Zenika/catption/codelab/chapter3 v0.0.0-20191117233821-d7ab6bbc1439
	github.com/Zenika/catption/font/impact v0.0.0-00010101000000-000000000000
	github.com/fogleman/gg v1.3.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.3.2
)

replace github.com/Zenika/catption/font/impact => ../../font/impact
