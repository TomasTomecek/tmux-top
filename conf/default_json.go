package conf

var default_conf string = `
{
	"load": {
		"intervals": [{
			"to": 0.1,
			"bg_color": "default",
			"fg_color": "colour10"
		}, {
			"from": 0.1,
			"to": 0.25,
			"bg_color": "default",
			"fg_color": "green"
		}, {
			"from": 0.25,
			"to": 0.7,
			"bg_color": "default",
			"fg_color": "colour166"
		}, {
			"from": 0.7,
			"to": 5.0,
			"bg_color": "default",
			"fg_color": "colour1"
		}, {
			"from": 5.0,
			"bg_color": "colour1",
			"fg_color": "white"
		}]
	},
	"mem": {
		"intervals": [{
			"to": 618076569.6,
			"bg_color": "default",
			"fg_color": "colour10"
		}, {
			"from": 618076569.6,
			"to": 1545191424.0,
			"bg_color": "default",
			"fg_color": "green"
		}, {
			"from": 1545191424.0,
			"to": 3090382848.0,
			"bg_color": "default",
			"fg_color": "colour166"
		}, {
			"from": 3090382848.0,
			"to": 5562689126.4,
			"bg_color": "default",
			"fg_color": "colour1"
		}, {
			"from": 5562689126.4,
			"to": 6180765696.0,
			"bg_color": "colour166",
			"fg_color": "white"
		}, {
			"from": 6180765696.0,
			"bg_color": "colour1",
			"fg_color": "white"
		}],
		"separator": "/",
		"separator_bg": "default",
		"separator_fg": "white",
		"total_bg": "default",
		"total_fg": "colour14"
	},
	"net": {
		"interfaces": {
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
		},
		"threshold": 0.0,
		"upload_label": "⬆",
		"upload_label_bg": "default",
		"upload_label_fg": "white",
		"download_label": "⬇",
		"download_label_bg": "default",
		"download_label_fg": "white",
		"intervals": [{
			"to": 1024.0,
			"bg_color": "default",
			"fg_color": "colour10"
		}, {
			"from": 1024.0,
			"to": 102400.0,
			"bg_color": "default",
			"fg_color": "green"
		}, {
			"from": 102400.0,
			"to": 524288.0,
			"bg_color": "default",
			"fg_color": "colour166"
		}, {
			"from": 524288.0,
			"to": 1048576.0,
			"bg_color": "default",
			"fg_color": "colour1"
		}, {
			"from": 1048576.0,
			"to": 5242880.0,
			"bg_color": "colour166",
			"fg_color": "white"
		}, {
			"from": 5242880.0,
			"bg_color": "colour1",
			"fg_color": "white"
		}]
	}

}`
