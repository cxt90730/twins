package server

func (s *TwinsServer) runSwitcher() {
	for {
		select {
		case <-s.SwitchCh:
			if AsElder {
				s.Logger.Debug("switch to little")
				AsElder = false
				go s.runAsLittle()
			} else {
				s.Logger.Debug("switch to elder")
				AsElder = true
				go s.runAsElder()
			}
		}
	}
}
