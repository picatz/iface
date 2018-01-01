# iface
> Network interface command-line utility.

Tools like `ifconfig`, `ip` and `ipconfig` are awesome! But, sometimes they're just a little too much to parse -- or you just don't want to fuss with all of that. With `iface`, you get a clean, scriptable command-line utility that's lightweight, cross-platform, and easy to learn. Learn once, user everywhere!

## Supported Platforms

> Thanks to [`VBaczynski`](https://github.com/VBaczynski) for helping test FreeBSD!

* macOS
* Linux 
* Windows
* FreeBSD 

## Build from Source

```
go get github.com/picatz/iface
```

Build `iface.exe` on macOS or Linux to run on a Windows box:

```
GOOS=windows GOARCH=386 go build iface.go
```

## Example Usage

Get the first `default` ( non-[loopback](https://en.wikipedia.org/wiki/Loopback#Virtual_network_interface) ) network [interface](https://en.wikipedia.org/wiki/Network_interface):
> This sort of simulates the LibPcap [`lookupdev`](https://linux.die.net/man/3/pcap_lookupdev) function to get the default interface to capture packets on.

```
$ iface default
```

Get all of the network interface names available on the system:

```
$ iface names
```

Get network interface names with thier [mac addresses](https://en.wikipedia.org/wiki/MAC_address):

```
$ iface macs 
```

Get network interface names with their [IP addresses](https://en.wikipedia.org/wiki/IP_address):

```
$ iface ips 
```

## Command-line Interface

```
NAME:
   iface - network interface command-line utility

USAGE:
   iface [global options] command [command options] [arguments...]

VERSION:
   1.0.1

COMMANDS:
     names    all network interface names
     macs     network interfaces with their mac addresses
     mtus     network interfaces with their maximum transmission unit (mtu)s
     ips      network interfaces with their (local) ip addresses
     default  the first non-loopback network interface
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
