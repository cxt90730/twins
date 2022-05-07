package task

import (
	"fmt"
	"time"
)

func init() {
	RegisterTask(ExampleTask{
		stopCh: make(chan struct{}),
	})
}

type ExampleTask struct{
	stopCh chan struct{}
}

func (e ExampleTask) Name() string {
	return "ExampleTask"
}

func (e ExampleTask) Run() error {
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

func (e ExampleTask) Stop() {
	fmt.Println("ExampleTask is stopping...")
	e.stopCh <- struct{}{}
}
