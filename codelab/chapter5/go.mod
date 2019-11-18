module github.com/Zenika/catption/codelab/chapter5

go 1.12

require (
	github.com/Zenika/catption/lib v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.3.2
)

replace github.com/Zenika/catption/font/impact => ../../font/impact

replace github.com/Zenika/catption/lib => ../../lib
