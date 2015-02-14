tmux-top
========

Monitoring information for your tmux status line.

`tmux-top` allows you to see your load, memory usage and network information in status of [http://tmux.sourceforge.net/](tmux).


Installation
------------

This tool is written in [http://golang.org/](Go). If you want to compile it, you have [http://golang.org/doc/install](setup your Go environment) first.

### Go distribution

```
go get github.com/TomasTomecek/tmux-top/cli
```

When the command succeeds, `tmux-top` is placed in directory `${GOPATH}/bin` and named `cli`. You can rename it easily:

```
mv ${GOPATH}/bin/{cli,tmux-top}
```

### Manual installation

```
git clone https://github.com/TomasTomecek/tmux-top.git
```

Once cloned, compile it using well-known process:

```
make
sudo make install
```

Usage
-----

There are three subcommands at the moment:

 1. `tmux-top load` — load of your workstation
 2. `tmux-top mem` — actual memry usage and total memory
 3. `tmux-top net` — network statistics: IP address, network interface and current bandwidth

Configuration
-------------

[https://github.com/TomasTomecek/tmux-top/blob/master/conf/default_json.go](This json) contains default configuration. If you want to change something, just override the json and store it in `~/.tmux-top`. You can change whatever you want. If the value is not found in your configuration file, it's loaded from default oone.
