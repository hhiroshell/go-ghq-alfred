# go-ghq-alfred

![](https://github.com/hhiroshell/go-ghq-alfred/workflows/test/badge.svg?branch=master)

Search local repos with ghq in Alfred Workflow.

This repository is forked from [pddg/go-ghq-alfred](https://github.com/pddg/go-ghq-alfred). Please see also original one.

## Environment

* Alfred 4.1 (or later)
* [ghq](https://github.com/motemen/ghq)

## Usage

This tool is a CLI tool. Output JSON strings.

```bash
$ ./go-alfred-ghq '{query}' $(ghq list -p)
```

## In Alfred

This workflow start with `ghq {query}` in alfred.  

Filtering the result of `ghq list -p` with `{query}` and show them.

### Preparing

You should specify path to `ghq`. Open this workflow settings, and edit environment variables. Default is `/usr/local/bin/ghq`.  

### Modifier key options

* **Enter**: Open repository in the Visual Studio Code.
* **Command + Enter**: Open repository in the Goland.
* **Control + Enter**: Open repository in the Intellij IDEA.
* **Shift + Enter**: Open repository in your default browser.
* **Option + Enter**: Open repository in your terminal.
* **Fn + Enter**: Search "user/repo" in google.

## Build

My environment is as follows.

* Go 1.15

```bash
$ go build .
```

## Attributes

Icons provided by www.flaticon.com.

### github and git logo

Icon made by Freepik from www.flaticon.com

### bitbucket logo

Icon made by Swifticons from www.flaticon.com

## Author

Customized by hhiroshell, based on [pudding's ghq-alfred](https://github.com/pddg/go-ghq-alfred)

## License

MIT