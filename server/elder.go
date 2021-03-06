package server

import (
	"context"
	"errors"
	"time"
	"twins/task"
)

func (s *TwinsServer) runAsElder() {
	s.Logger.Debug("run as elder now.")
	ticker := time.NewTicker(1 * time.Second)
	err := s.sendHeartbeat()
	if err != nil {
		return
	}
	AsElder = true
	taskMap := task.LoadAllTasks()
	if len(taskMap) == 0 {

	}
	for name, t := range taskMap {
		s.Logger.Info("task", name, "start to run.")
		go func(t task.Task) {
			err := task.ReadTaskConfig(t)
			if err != nil {
				panic("load task config " + t.Name() + " failed:" + err.Error())
			}
			if t.Enabled() {
				err = t.Load()
				if err != nil {
					s.Logger.Error("task", t.Name(), "load error:", err)
					panic(err)
				}
				err = t.Run()
				if err != nil {
					s.Logger.Error("task", t.Name(), "run error:", err)
				}
			}
		}(t)
	}
	for {
		select {
		case <-ticker.C:
			if err := s.sendHeartbeat(); err != nil {
				return
			}
		}
	}
}

func (s *TwinsServer) sendHeartbeat() error {
	resp, err := s.RpcClient.SendHeartBeat(context.Background())
	if err != nil {
		s.Logger.Warn("Send heartbeat to little err:", err)
		return nil
	}
	if resp.IsElderNow {
		for _, t := range task.LoadAllTasks() {
			t.Stop()
		}
		s.SwitchCh <- true
		s.Logger.Warn("another is elder, I'm little now.")
		return errors.New("another is elder")
	}
	return nil
}
