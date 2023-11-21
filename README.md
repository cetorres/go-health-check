# go-health-check
This is a simple CLI program written in Go to check if a website is running and reachable or not.

I used the following library:
- [urfave/cli](https://cli.urfave.org): A simple, fast, and fun package for building command line apps in Go. The goal is to enable developers to write fast and distributable command-line applications in an expressive way.

## Build and usage
To build and run the program, just run:

```
$ go build
$ ./go-health-check
```

## Output

```
NAME:
   Website Health Checker - A small tool to check if a website is running and reachable or not.

USAGE:
   Website Health Checker [global options] command [command options] [arguments...]

VERSION:
   1.0

AUTHOR:
   Carlos E. Torres <cetorres@cetorres.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --domain value, -d value  domain name to check
   --port value, -p value    port number to check
   --help, -h                show help
   --version, -v             print the version

COPYRIGHT:
   2023 Carlos E. Torres
2023/11/21 11:55:32 Required flag "domain" not set
```

Run with a domain to check:

```
$./go-health-check -d cetorres.com
[UP] cetorres.com (138.197.45.25:80) is reachable.
```
