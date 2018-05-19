package conf

import (
	"encoding/json"
	"io/ioutil"
)

// Main include info about API settings.
type Main struct {
	Port  string `json:"port"`
	Debug Debug  `json:"debug"`
	DB    DB     `json:"database"`
}

// DB include info about database connection.
type DB struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

// Debug include all debugging flags
type Debug struct {
	Log bool `json:"log"`
	DB  bool `json:"db"`
}

// FromFile read config from path and create configuration struct.
func FromFile(path string) Main {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic("error reading config " + path + ": " + err.Error())
	}
	var conf Main
	if err := json.Unmarshal(bytes, &conf); err != nil {
		panic("error parsing config " + path + ": " + err.Error())
	}
	return conf
}
