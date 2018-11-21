package main

import (
	"errors"
	"io/ioutil"
	"os"
	//	"path/filepath"
	//	"sort"
	//	"strings"

	"gopkg.in/yaml.v2"
)

var (
	configModtime  int64
	errNotModified = errors.New("Not modified")
)

type Host struct {
	Hostgroup []string `yaml:"group"`
	Retention string   `yaml:"retention"`
}

type Config struct {
	Hosts            []Host `yaml:"hosts"`
	Aggregator       string `yaml:"aggregator"`
	DefaultRetention string `yams:"retention"`
	LogPath          string `yaml:"logpath"`
	LogLevel         string `yaml:"loglevel"`
	LogFilename      string `yaml:"log_filename"`
	LogMaxSize       int    `yaml:"log_max_size"`
	LogMaxBackups    int    `yaml:"log_max_backups"`
	LogMaxAge        int    `yaml:"log_max_age"`
}

func readConfig(ConfigName string) (x *Config, err error) {
	var file []byte
	if file, err = ioutil.ReadFile(ConfigName); err != nil {
		return nil, err
	}
	x = new(Config)
	if err = yaml.Unmarshal(file, x); err != nil {
		return nil, err
	}
	if x.LogLevel == "" {
		x.LogLevel = "Debug"
	}
	return x, nil
}

func reloadConfig(configName string) (cfg *Config, err error) {
	info, err := os.Stat(configName)
	if err != nil {
		return nil, err
	}
	if configModtime != info.ModTime().UnixNano() {
		configModtime = info.ModTime().UnixNano()
		cfg, err = readConfig(configName)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}
	return nil, errNotModified
}
