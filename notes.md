# Init

Create dir with the name `wingnut` change into this dir and run:
```shell
go mod init github.com/ocular-d/wingnut
```

You can run this command in the terminal of your project’s working directory to install the Cobra-cli package.

```shell
go install github.com/spf13/cobra-cli@latest
```

After installing the package, you can initialize a Cobra CLI project with the init command of the command line tool.

```shell
cobra-cli init
```

## Commands

Use `cobra add`, the following example creates a new command with the name *foo*.

```shell
cobra-cli add repo
```

### Subcommands

The following example adds the subcommand *bar* to the command *foo*.

```shell
cobra-cli add -p bar fooCmd
```

### Create CLI reference docs from the code



