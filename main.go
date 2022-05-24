package main

import (
	"flag"
	"fmt"
	"time"
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

type ExampleTask struct {
	Enable bool
	stopCh chan struct{}
}

func (e *ExampleTask) Name() string {
	return "example"
}

func (e *ExampleTask) Load() error {
	return nil
}

func (e *ExampleTask) Run() error {
	count := 0
	for {
		select {
		case <-e.stopCh:
			fmt.Println("ExampleTask is stopped.")
			return nil
		default:
			count++
			fmt.Println("ExampleTask is running...")
			time.Sleep(1 * time.Second)
			if count > 30 {
				fmt.Println("ExampleTask finished.")
				return nil
			}
		}
	}
}

func (e *ExampleTask) Stop() {
	fmt.Println("ExampleTask is stopping...")
	e.stopCh <- struct{}{}
}

func (e *ExampleTask) Enabled() bool {
	return e.Enable
}

const DefaultConfPath = "/etc/twins/twins.toml"
