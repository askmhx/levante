package main

import (
	"flag"
	"iosxc.com/levante/app"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "NOT SET", "app cfg path")
	flag.Parse()
	app.Run(cfgPath)
}
