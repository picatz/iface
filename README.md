# iface
> Network interface command-line utility.

Tools like `ifconfig`, `ip` and `ipconfig` are awesome! But, sometimes they're just a little too much to parse -- or you just don't want to fuss with all of that. With `iface`, you get a clean, scriptable command-line utility that's lightweight, cross-platform, and easy to learn. 

## Build from Source

Until I've figured out a better solution:

```
git clone https://github.com/picatz/iface.git
cd isit
go build isit.go
mv isit /usr/bin/local
```

## Example Usage

Get the first `default` ( non-loopback ) network interface:
> This sort of simulates the LibPcap `lookupdev` function to get the default interface to capture packets on.

```
$ iface default
```

Get all of the network interface names available on the system:

```
$ iface names
```

Get all of the network interface names available on the system with their mac addresses:

```
$ iface macs 
```

## Command-line Interface

```
NAME:
   iface - network interface command-line utility

USAGE:
   iface [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     names    all network interface names
     macs     all network interfaces with their mac addresses
     mtus     all network interfaces with their maximum transmission unit (mtu)s
     default  the first non-loopback network interface
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

