summary: Basics of writing a Go CLI tool
id: codelab
categories: golang
tags: cli
status: Published
authors: Nicolas Lepage
feedback link: https://github.com/Zenika/catption/issues

# Basics of writing a Go CLI tool

## Introduction

In this codelab you will learn the basics of writing a Go CLI tool.

### What you'll learn
 - Setup a development environment
 - Discover `os`, `os/exec` and `flag` packages
 - Discover `github.com/spf13/cobra` CLI library
 - Create commands and subcommands
 - Read command flags and args
 - Discover `github.com/spf13/viper` config library
 - Read and write a config file
 - Put `cobra` and `viper` together
 - Read environment variables
 - Discover `github.com/sirupsen/logrus` logging library
 - Use build time variable injection
 - Use conditional compilation and build tags

The steps marked with a 🎁 are optional.

## Ch.1: Introduction

### What you'll learn
 - Setup a development environment
 - Read args (package `os`)
 - Bonus: Read flags (package `flag`)

## Ch.1: Setup environment

In order to go through this codelab, you are going to need a working Go development environment.

The minimum required version is Go 1.12.

Positive
: Already have Go installed?
Make sure you are running a version >= 1.12 by running `go version`.
If it is the case you may proceed to the next step.

### 🐧 Linux

Negative
: Do not use `apt` (old versions of Go)

#### snap

Run:

```sh
sudo snap install go --classic
```

#### tarbal

Follow the instructions at [https://golang.org/doc/install#tarball](https://golang.org/doc/install#tarball)

### 🍏 macOS

#### brew

Run:
```sh
brew install go
```

#### tarbal

Download the package file at [https://golang.org/dl/](https://golang.org/dl/), open it, and follow the prompts.

### 🏁 Windows

Download the MSI file at [https://golang.org/dl/](https://golang.org/dl/), open it, and follow the prompts.

Positive
: Check your installation by running `go version` and `go env`.

## Ch.1: Download codelab

There are two ways of downloading the codelab contents.
The prefered way is `git`, which will allow you to keep track of your work and revert things if needed.

### git

Run:

```sh
git clone https://github.com/Zenika/catption.git
```

### zip

Download [https://github.com/Zenika/catption/archive/master.zip](https://github.com/Zenika/catption/archive/master.zip) and unzip it.

Positive
: Each chapter of the codelab has its own directory:
```
📂 catption
|-📂 codelab
| |-📁 chapter1
| |-📁 chapter2
```
Run `cd catption/codelab/chapter1` to go to chapter 1.

## Ch.1: Choose an IDE

The last thing you need is a Go friendly IDE.

If you don't already have one, here are some popular IDEs for Go:
 - [Visual Studio Code](https://code.visualstudio.com/) + [vscode-go](https://github.com/microsoft/vscode-go)
 - [Goland](https://www.jetbrains.com/go/)
 - vim + [vim-go](https://github.com/fatih/vim-go)

Now open the codelab contents and you are ready 👷, let's Go!

## Ch.1: Read args

### Run `hello.go`

In `📂catption/codelab/chapter1` you will find a classic `hello.go`:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
```

⌨ Execute this program by running `go run hello.go`.

### Format the message

We would like to replace `World` by a variable in our message.

⌨ Create a new string variable:
```go
var recipient = "Gopher"
```

⌨ Use [`fmt.Printf()`](https://pkg.go.dev/fmt?tab=doc#Printf) to format the message with `recipient`.

Negative
: Unlike `fmt.Println()`, `fmt.Printf()` does not add a new line at the end of the string.
You must add it by appending `\n` at the end of the message.

### Read command line arguments

As you can see the main function of a Go program has no parameters.

The command line arguments are available in the [`Args` variable of the `os` package](https://pkg.go.dev/os?tab=doc#pkg-variables).

Positive
: `os.Args` has the type `[]string` (slice of string).
A slice is a variable length array.

⌨ Use `os.Args` to fill the `recipient` variable.

Positive
: [`strings.Join`](https://pkg.go.dev/strings?tab=doc#Join) concatenates the elements of a slice of strings.

Positive
: To extract a subset of a slice, use the slice operator.
Having `var ii = []int{1, 2, 3, 4}`, `ii[2:]` will give you the slice `[3, 4]`

## Ch.1: 🎁 Interpret flags

Flags allow to change the behavior of commands, like the `-r` flag of `rm` which enables recursive removal.

The [`flag` package](https://pkg.go.dev/flag?tab=doc) allows to parse the flags contained in `os.Args`.

We would like our command to have a `-u` flag which uppercases the message:
```sh
$ hello -u capslock
HELLO CAPSLOCK!
```

⌨ Explore the [`flag` package](https://pkg.go.dev/flag) and parse the `-u` flag in `hello.go`.

Positive
: [`flag.Args`](https://pkg.go.dev/flag?tab=doc#Args) returns the non-flag command-line arguments.

Positive
: [`strings.ToUpper`](https://pkg.go.dev/strings?tab=doc#ToUpper) returns an upper case copy of a string.

Positive
: [`fmt.Sprintf](https://pkg.go.dev/fmt?tab=doc#Sprintf) returns a formatted string

## Ch.1: End

🎉 Congratulations! You have completed chapter 1.

### What we've covered
 - Setup a development environment
 - Read args (package `os`)
 - 🎁 Read flags (package `flag`)

## Ch.2: Introduction

### What you'll learn
 - Discover `github/spf13/cobra`
 - Create a cobra command
 - 🎁 Validate arguments

## Ch.2: Discover cobra

> Cobra is a library for creating powerful modern CLI applications.

![Cobra](codelab/assets/cobra.png)

Cobra provides:
* Easy subcommand-based CLIs: `app server`, `app fetch`, etc.
* Fully POSIX-compliant flags (including short & long versions)
* Nested subcommands
* Global, local and cascading flags
* Easy generation of applications & commands with `cobra init appname` & `cobra add cmdname`
* Intelligent suggestions (`app srver`... did you mean `app server`?)
* Automatic help generation for commands and flags
* Automatic help flag recognition of `-h`, `--help`, etc.
* Automatically generated bash autocomplete for your application
* Automatically generated man pages for your application
* Command aliases so you can change things without breaking them
* The flexibility to define your own help, usage, etc.
* Optional tight integration with [viper](https://pkg.go.dev/github.com/spf13/viper?tab=overview) for 12-factor apps

👀 Explore cobra's [documentation](https://github.com/spf13/cobra/blob/master/README.md) and [API](https://pkg.go.dev/github.com/spf13/cobra?tab=doc).

## Ch.2: Create a command

Let's see how to recreate our hello command using Cobra.

In `📂catption/codelab/chapter2` you will find a new `hello.go` with the skeleton of a cobra app:
```go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	RunE: func(_ *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func sayHello(args []string) error {
	if _, err := fmt.Printf("Hello %s!\n", strings.Join(args, " ")); err != nil {
		return err
	}
	return nil
}
```

### Describe the command

⌨ Fill the `Use` and `Long` fields of the `cmd` [Command struct](https://pkg.go.dev/github.com/spf13/cobra?tab=doc#Command), then execute `go run hello.go -h` to see the result.

### Implement the command

⌨ Call `sayHello` in the `RunE` function of `cmd` in order to have a working hello command, execute `go run hello.go cobra` to see the result.

Negative
: `sayHello` may return an error, you may forward this error to the caller of `RunE`.

### Version the command

⌨ Finally fill the `Version` field of `cmd`, then execute `go run hello-go --version` to see the result.

## Ch.2: 🎁 Validate args

Our hello command needs at least one command line argument.

⌨ Fill the `Args` field of `cmd` with the correct value in order to raise an error if hello doesn't receive any arguments.

Positive
: The type of `Args` is [`cobra.PositionalArgs`](https://pkg.go.dev/github.com/spf13/cobra?tab=doc#PositionalArgs), which is a function type.
You could implement your own command-line arguments validator (this is not the goal here).

## Ch.2: End

🎉 Congratulations! You have completed chapter 2.

### What we've covered
 - Discover `github/spf13/cobra`
 - Create a cobra command
 - 🎁 Validate arguments

## Ch.3: Introduction

### What you'll learn
 - Interpret flags
 - 🎁 Flag shorthand

## Ch.3: Interpret flags

Enough of hello messages, let's start writing our cat caption CLI 🐱

In `📂catption/codelab/chapter3` you will find a `catption.go` with a new command:

```go
var (
	top, bottom            string
	size, fontSize, margin float64

	cmd = &cobra.Command{
		Use:     "catption",
		Long:    "Cat caption generator CLI",
		Args:    cobra.ExactArgs(1),
		Version: "chapter3",
		RunE: func(_ *cobra.Command, args []string) error {
			var name = args[0]

			cat, err := catption.LoadJPG(name)
			if err != nil {
				return err
			}

			cat.Top, cat.Bottom = top, bottom
			cat.Size, cat.FontSize, cat.Margin = size, fontSize, margin

			return cat.SaveJPG("out.jpg")
		},
	}
)
```

This command does 3 things:
1. Create a catption by loading a JPEG file
2. Setup the catption's parameters
3. Write the catption to `out.jpg`

However the variables used to setup the catption have not been initialized.

### Define flags

⌨ In the `init` function, setup `cmd`'s flags:
 - `top` and `bottom` string flags
 - `size`, `fontSize` and `margin` float flags (Use `catption.DefaultSize`, `catption.DefaultFontSize` and `catption.DefaultMargin` as default values)

Positive
: [`Command.Flags`](https://pkg.go.dev/github.com/spf13/cobra?tab=doc#Command.Flags) returns the [`FlagSet`](https://pkg.go.dev/github.com/spf13/pflag?tab=doc#FlagSet) of a command.
 The `FlagSet` allows to setup the flags of a command.

Positive
: Some methods of `FlagSet`, such as [`IntVar`](https://pkg.go.dev/github.com/spf13/pflag?tab=doc#IntVar), expect a pointer as first argument.
Having `var i = 42`, use `&i` to get a pointer to `i`, `&i` has the type `*int`.

⌨ Play around with your new command, some pictures are available in `📂cats/`

## Ch.3 🎁 Flags shorthand

Flags shorthands allow users to type more concise commands.

⌨ Add some shorthands to `cmd`:
 - `-t` for `--top`
 - `-b` for `--bottom`
 - `-s` for `--size`

Positive
: All `FlagSet` methods have a shorthand variant.
To add a shorthand to an int flag, use [`IntVarP`](https://pkg.go.dev/github.com/spf13/pflag?tab=doc#FlagSet.IntVarP) instead of `IntVar`.

## Ch.3: End

🎉 Congratulations! You have completed chapter 3.

### What we've covered
 - Interpret flags
 - 🎁 Flag shorthand

## Ch.4: Introduction

### What you'll learn
 - Discover `github.com/spf13/viper`
 - Read a config file
 - 🎁 Access user's config dir

## Ch.4: Discover viper

> Viper is a complete configuration solution for Go applications including 12-Factor apps.
> It is designed to work within an application, and can handle all types of configuration needs and formats.

![Viper](codelab/assets/viper.png)

It supports:

* setting defaults
* reading from JSON, TOML, YAML, HCL, envfile and Java properties config files
* live watching and re-reading of config files (optional)
* reading from environment variables
* reading from remote config systems (etcd or Consul), and watching changes
* reading from command line flags
* reading from buffer
* setting explicit values

👀 Explore viper's [documentation](https://github.com/spf13/viper/blob/master/README.md) and [API](https://pkg.go.dev/github.com/spf13/viper?tab=doc).

## Ch.4: Read config file

Specifying the full path to the input JPEG file is not very userfriendly...

Let's use a config file to define directories where catption should look for JPEG files.

In `📂catption/codelab/chapter4` the catption command now has a `PreRunE` function:

```go
PreRunE: func(_ *cobra.Command, _ []string) error {
	viper.SetConfigName("catption")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	return nil
},
```

This function tries to load a `catption.*` config file in the current directory.

### Define default config values

⌨ Before the call to `ReadInConfig`, define the default value for the `"dirs"` config key (use the value of the `dirs` var).

Positive
: [`viper.SetDefault`](https://pkg.go.dev/github.com/spf13/viper?tab=doc#SetDefault) allows to define default values for config keys.

### Read config values

⌨ After the call to `ReadInConfig`, set the value of the `dirs` var using the `"dirs"` config key.

Positive
: viper has all kinds of getters for reading config keys.
[`viper.GetIntSlice`](https://pkg.go.dev/github.com/spf13/viper?tab=doc#GetIntSlice) reads a config key into a slice of ints (`[]int`).

### Create a config file

⌨ Create a `catption.*` config file with the directories where you want catption to look for JPEG files.

Example `catption.yaml`:

```yaml
dirs:
  - "."
  - "../../cats"
```

You can now try your configuration: `go run catption.go -t "Hello" -b "World" dinner.jpg`

## Ch.4: 🎁 Config dir

Many applications read there config file from the user's config directory (`$HOME/Library/Application Support` on macOS for example).

⌨ Call [`viper.AddConfigPath`](https://pkg.go.dev/github.com/spf13/viper?tab=doc#AddConfigPath) a second time to read catption's config file from the user's config directory, in addition of current the directory.

Positive
: Package `os` has some useful helpers such as [`UserHomeDir`](https://pkg.go.dev/os?tab=doc#UserHomeDir) to read platform dependent environment variables.

## Ch.4: End

🎉 Congratulations! You have completed chapter 4.

### What we've covered
 - Discover `github.com/spf13/viper`
 - Read a config file
 - 🎁 Access user's config dir

## Ch.5: Introduction

### What you'll learn
 - Connect cobra and viper
 - 🎁 Read environment variables

## Ch.5: cobra 🔌 viper

Some of our users don't want to use config files.

We would like to offer them the possibility to override the `dirs` config key with a flag.

Luckily viper has the ability to read config values from cobra!

Negative
: When connecting cobra and viper, you must read config values from viper.
viper reads values from cobra, but not the other way around.

⌨ Create a new `dir` flag with the type slice of strings.

⌨ Bind the `dir` flag to viper's `dirs` config key.

Positive
: [`FlagSet.Lookup`](https://pkg.go.dev/github.com/spf13/pflag?tab=doc#FlagSet.Lookup) returns the `*pflag.Flag` for a previously created flag's name.

Positive
: [`viper.BindPFlag`](https://pkg.go.dev/github.com/spf13/viper?tab=doc#BindPFlag) binds a config key to a `*pflag.Flag`.

Try it out: `go run catption.go -t "Hello" -b "World" --dir "../../cats" --dir "." dinner.jpg`

## Ch.5: 🎁 Read env vars

One of our users would like to deploy catption on a kubernetes cluster.

The easiest way for him/her to specify the input files directories is to use an environment variable.

⌨ Use viper's API to read the `dirs` config key from a `CATPTION_DIRS` environment variable.

Try it out: `CATPTION_DIRS="../../cats" go run catption.go -t "Hello" -b "World" dinner.jpg`

## Ch.5: End

🎉 Congratulations! You have completed chapter 5.

### What we've covered
 - Connect cobra and viper
 - 🎁 Read environment variables

## Ch.6: Introduction

### What you'll learn
 - Create a subcommand
 - 🎁 Inject compile time variables

## Ch.6: Subcommands

Some of our users don't know how to create a config file and add directories to it.

Let's help them by adding a new `dir` subcommand to catption, which will add a directory to the config file.

In `📂catption/codelab/chapter6` we now have a `dirCmd` command, and a `addDir` function which implements adding a new directory to the config file.

⌨ Fill the fields of `dirCmd`: `Use`, `Long`, `Args` and `RunE`

⌨ In the `init` function, add `dirCmd` as a subcommand to `cmd`

Positive
: `cmd`'s `RunE` function is now a `PersistentPreRunE`.
It will be executed for `cmd` and it's subcommands.

Positive
: [`Command.AddComand`](https://pkg.go.dev/github.com/spf13/cobra?tab=doc#Command.AddCommand) adds a subcommand to a parent command

## Ch.6: 🎁 Compile vars

Using a constant value for `cmd`'s `Version` field is not very useful.

It would be nice to set this variable at compile time, with a git tag or commit hash.

⌨ Create a `version` variable at package level, and set `cmd.Version`'s value with this variable.

⌨ Try changing the binary's version with build flags: `go build -ldflags "-X main.version=1.0.0"`

## Ch.6: End

🎉 Congratulations! You have completed chapter 6.

### What we've covered
 - Create a subcommand
 - 🎁 Inject compile time variables

## Ch.7: Introduction

### What you'll learn
 - Interpret custom flags
 - 🎁 Discover `github.com/sirupsen/logrus`

## Ch.7: Custom flags

We've added some logs to catption using a library called [logrus](https://pkg.go.dev/github.com/sirupsen/logrus?tab=doc).

However we would like to be able to set the log level using a flag.

In `📂catption/codelab/chapter7` we now have a `logLevel` variable used to set the log level.
This variable has the type [`logrus.Level`](https://pkg.go.dev/github.com/sirupsen/logrus?tab=doc#Level).

In order to create a flag with a custom type, you must implement pflag's [`Value` interface](https://pkg.go.dev/github.com/spf13/pflag?tab=doc#Value).

This is already done by the type `logLevelValue`:

```go
type logLevelValue logrus.Level

var _ pflag.Value = new(logLevelValue)

func (l *logLevelValue) Set(value string) error {
	lvl, err := logrus.ParseLevel(value)
	if err != nil {
		return err
	}
	*l = logLevelValue(lvl)
	return nil
}

func (l *logLevelValue) String() string {
	return logrus.Level(*l).String()
}

func (l *logLevelValue) Type() string {
	return "string"
}
```

⌨ In the `init` function, create a new `--logLevel` flag for the `logLevel` variable.

Positive
: [`Command.PersistentFlags`]() returns a `FlagSet` used for the current command and its subcommands.

Positive
: [`FlagSet.Var`](https://pkg.go.dev/github.com/spf13/pflag?tab=doc#FlagSet.Var) defines a custom typed flag.

It is possible to perform a type cast between pointer types, here is an example:

```go
type Celsius float64

func example() {
	var temperature float64
	measureTemperature((*Celsius)(&temperature))
	fmt.Println("temp:", temperature)
}

// measureTemperature stores a new measure in the t pointer
func measureTemperature(t *Celsius)
```

## Ch.7: 🎁 Discover logrus

👀 Have a look at logrus's [documentation](https://github.com/sirupsen/logrus/blob/master/README.md) and [API](https://pkg.go.dev/github.com/sirupsen/logrus?tab=doc)

⌨ Add some new logs in catption.

## Ch.7: End

🎉 Congratulations! You have completed chapter 7.

### What we've covered
 - Interpret custom flags
 - 🎁 Discover `github.com/sirupsen/logrus`

## Ch.8: Introduction

### What you'll learn
 - Discover `os/exec` package
 - Use conditional compilation
 - 🎁  Use build tags

## Ch.8: Discover os/exec

We would like catption to open an image viewer as soon as the image has been written to disk.

Most operating systems have commands to open the appropriate viewer for a file:
 - The `xdg-open` command on 🐧 Linux
 - The `open` command on 🍏 macOS
 - The `start` command on 🏁 Windows

⌨ Use the [`os/exec` package](https://pkg.go.dev/os/exec?tab=doc) to execute the appropriate command for your OS and display the image.

Positive
: [`Cmd.Run`](https://pkg.go.dev/os/exec?tab=doc#Cmd.Run) starts a command and waits for it to complete.

## Ch.8: Conditional compile

Some users don't have the same OS as you.

We would like to cross-compile catption to other systems, but the command for opening a viewer is system dependent!

The go compiler is able to include/exclude source files, based on their suffix.
`source_darwin.go` will only be compiled when targeting macOS systems.

⌨ Create 3 files with each an `openCmd` string const:
 - `open_linux.go` for Linux
 - `open_darwin.go` for macOS
 - `open_windows.go` for Windows

⌨ Use `openCmd` to call [`exec.Command`](https://pkg.go.dev/os/exec?tab=doc#Command)

## Ch.8: 🎁 Build tags

One of our users would like to run catption on a FreeBSD system.

`xdg-open` is also available on this system, it would be nice to use the same `openCmd` const for Linux and FreeBSD.

⌨ Rename `open_linux.go` to `open_xdg.go`.

⌨ Add [build tags](https://golang.org/pkg/go/build/#hdr-Build_Constraints) to `open_xdg.go` in order to target Linux and FreeBSD.

## Ch.8: End

🎉 Congratulations! You have completed chapter 8.

### What we've covered
 - Discover `os/exec` package
 - Use conditional compilation
 - 🎁  Use build tags

## End

🎉 Congratulations! You have completed the codelab!

You now know the basics to build you own CLI with Go.

### What we've covered
 - Setup a development environment
 - Discover `os`, `os/exec` and `flag` packages
 - Discover `github.com/spf13/cobra` CLI library
 - Create commands and subcommands
 - Read command flags and args
 - Discover `github.com/spf13/viper` config library
 - Read and write a config file
 - Put `cobra` and `viper` together
 - Read environment variables
 - Discover `github.com/sirupsen/logrus` logging library
 - Use build time variable injection
 - Use conditional compilation and build tags

The fully working catption CLI source is available at the repositories root.
