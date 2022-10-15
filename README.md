# sample-go-cli

Sample Go-CLI build with cobra.

# How to setup

* https://go.dev/doc/install
* go mod init github.com/jfeddern/samplegocli
* touch main.go
* go get -u github.com/spf13/cobra@latest (https://github.com/spf13/cobra)
* go install github.com/spf13/cobra-cli@latest
* cobra-cli init
* cobra-cli add list

Basic setup done in order to start with the real implementation.
In order to verify if everything works, build the go app locally and check if you can already see the created commands.

* go build
* ./samplegocli -h

This should output the following context:
```
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  samplegocli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        A brief description of your command

Flags:
  -h, --help     help for samplegocli
  -t, --toggle   Help message for toggle

Use "samplegocli [command] --help" for more information about a command.
```

# About Go
Common package structure https://github.com/golang-standards/project-layout#go-directories