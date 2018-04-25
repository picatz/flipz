# flipz 
> Table flipping command-line utility.

## Install

```
$ go get github.com/picatz/flipz
```

### Build from Source

```shell
$ go get github.com/urfave/cli
$ git clone github.com/picatz/flipz
$ cd flipz
$ go build flipz
```

## Command-line Options

The `flip` command flips a table. This is also the default command.

```shell
$ flipz flip
[ ╯ ' □']╯︵┻━┻)
```

The `put` command puts a table back (to be flipped again!).

```shell
$ flipz put
┬──┬ノ[･ω･ ノ ]
```

## Help Menu

```
NAME:
   flipz - command-line flip table utility

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     flip, f  random flip table
     put, p   random put table
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
