package server

import "github.com/BurntSushi/toml"

type TwinsConf struct {
	InitAsElder bool
	Bind        string
	TwinsBind   string
	LogPath     string
	LogLevel    string
	TaskConfDir string
	TaskLogDir  string
}

func readConfig(path string) (conf TwinsConf, err error) {
	_, err = toml.DecodeFile(path, &conf)
	return conf, err
}
