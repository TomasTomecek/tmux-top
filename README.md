[![Build Status](https://drone.io/github.com/TomasTomecek/tmux-top/status.png)](https://drone.io/github.com/TomasTomecek/tmux-top/latest)

tmux-top
========

Monitoring information for your [tmux](http://tmux.sourceforge.net/) status line.

`tmux-top` allows you to see your:

 * load
 * memory usage
 * network information
 * I/O

![tmux-top sample](https://raw.githubusercontent.com/TomasTomecek/tmux-top/master/docs/tmux_top_example.png)
![tmux-top sample](https://raw.githubusercontent.com/TomasTomecek/tmux-top/master/docs/tmux_top_example2.png)


Installation
------------

This tool is written in [Go](http://golang.org/). If you want to compile it, you have to [setup your Go environment](http://golang.org/doc/install) first.

#### Supported platforms

 * linux

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
 4. `tmux-top io` — I/O statistics: current reads and writes

Configuration
-------------

[This json](https://github.com/TomasTomecek/tmux-top/blob/master/conf/default_json.go) contains default configuration. If you want to change something, just override the json and store it in `~/.tmux-top`. You can change whatever you want. If the value is not found in your configuration file, it's loaded from default one.

Your configuration may look like this:

```json
{
  "net": {
    "interfaces": {
      "enp0s25": {
        "alias": "E",
        "label_color_fg": "white",
        "label_color_bg": "default",
        "address_color_fg": "colour4",
        "address_color_bg": "default"
      },
      "enp5s0": {
        "alias": "E",
        "label_color_fg": "white",
        "label_color_bg": "default",
        "address_color_fg": "colour4",
        "address_color_bg": "default"
      },
      "wlp3s0": {
        "alias": "W",
        "label_color_fg": "white",
        "label_color_bg": "default",
        "address_color_fg": "green",
        "address_color_bg": "default"
      },
      "tun0": {
        "alias": "V",
        "label_color_fg": "white",
        "label_color_bg": "default",
        "address_color_fg": "colour3",
        "address_color_bg": "default"
      }
    }
  }
}
```

and tmux configuration:

```shell
set -g status-left "#(tmux-top n)"
set -g status-right "#(tmux-top m) #[fg=white]:: #(tmux-top l)"
```

Layout inspiration from [this blog post](http://zanshin.net/2013/09/05/my-tmux-configuration/ ).

Other goodies for tmux
----------------------

 * [tmux-mem-cpu-load](https://github.com/thewtex/tmux-mem-cpu-load)
 * [powerline](https://github.com/powerline/powerline)
 * [rainbarf](https://github.com/creaktive/rainbarf)
 * [Battery](https://github.com/Goles/Battery)
