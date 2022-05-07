package server

import (
	"time"
)

func (s *TwinsServer) runAsLittle() {
	s.Logger.Debug("I'm little.")
	ticker := time.NewTicker(1 * time.Second)
	waitElder := 0
	for {
		select {
		case <-ticker.C:
			if s.RpcServer.Checkpoint == 0 {
				waitElder++
				if waitElder < 5 {
					continue
				}
			}
			if time.Now().UTC().UnixNano()-s.RpcServer.Checkpoint > 5*1e9 {
				s.Logger.Warn("Cannot recv heartbeat in 5 secs.")
				s.SwitchCh <- true
				return
			}
		}
	}
}
