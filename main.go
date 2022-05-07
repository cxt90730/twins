package main

import (
	"flag"
	. "twins/server"
)


func main() {
	confPath := flag.String("c", DefaultConfPath, "config file path")
	flag.Parse()
	server := NewTwinsServer(*confPath)
	err := server.Run()
	if err != nil {
		panic(err)
	}
}


