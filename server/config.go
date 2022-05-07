package server

import "github.com/BurntSushi/toml"

const DefaultConfPath = "/etc/twins/twins.toml"

type conf struct {
	InitAsElder bool
	Bind        string
	TwinsBind   string
	LogPath     string
	LogLevel    string
}

func readConfig(path string) (conf conf, err error) {
	_, err = toml.DecodeFile(path, &conf)
	return conf, err
}
