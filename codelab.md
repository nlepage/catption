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

### 🍏 Mac OS

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
| |-📁 chapter2
...
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

## Chapitre 2
   * Découverte de cobra
   * Création d'une commande
   * Bonus : Validation des arguments

## Chapitre 3
   * Interprétation des flags
   * Bonus : TODO

## Chapitre 4
   * Découverte de viper
   * Lecture d'un fichier de config
   * Bonus : TODO

## Chapitre 5
   * Connexion cobra/viper
   * Bonus : Lecture variable d'environnement

## Chapitre 6
   * Création d'une sous-commande
   * Écriture d'un fichier de config
   * Bonus : Injection de variable à la compilation

## Chapitre 7
   * Interprétation d'un flag custom
   * Bonus : Utilisation de logrus

## Chapitre 8
   * Découverte du package `os/exec`
   * Découverte de la compilation conditionnelle
   * Bonus : Découverte des build tags

## FAQ

ah!