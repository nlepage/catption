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
 - Bonus: Reading flags (package `flag`)

## Ch.1: Setup a development environment

In order to go through this codelab, you are going to need a working Go development environment.

The minimum required Go version is 1.12.

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

Check your installation by running `go version` and `go env`.

#### tarbal

Follow the instructions at https://golang.org/doc/install#tarball

Check your installation by running `go version` and `go env`.

### 🍏 Mac OS

Download the package file at https://golang.org/dl/, open it, and follow the prompts.

Check your installation by running `go version` and `go env`.

### 🏁 Windows

Download the MSI file at https://golang.org/dl/, open it, and follow the prompts.

Check your installation by running `go version` and `go env`.

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