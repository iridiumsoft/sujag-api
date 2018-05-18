package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Main include info about API settings.
type Main struct {
	Port          string        `json:"port"`
	Debug         Debug         `json:"debug"`
	Auth          Auth          `json:""`
	DB            DB            `json:"database"`
	Security      Security      `json:"security"`
	Poloniex      Poloniex      `json:"poloniex"`
	MiningHamster MiningHamster `json:"MiningHamster"`
	Binance       Binance       `json:"binance"`
}

// Auth includes all authorization settings
type Auth struct {
	AccessToken  string `json:"accessToken"`         // CMS token
	AuthTokenTTL int    `json:"authTokenTtlSeconds"` // Time to life of user authentication link, seconds
	JWTSecret    string `json:"jwtSecret"`           // Secret key for generating user token
	JwtTtl       int    `json:"jwtTokenTtlMinutes"`  // Time to life of user session token
}

// DB include info about database connection.
type DB struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

type MiningHamster struct {
	ApiKey string `json:"api_key"`
	URI    string `json:"uri"`
}

type Poloniex struct {
	ApiKey  string `json:"apiKey"`
	Secrete string `json:"secret"`
}

type Binance struct {
	ApiKey  string `json:"apiKey"`
	Secrete string `json:"secret"`
}

// Debug include all debugging flags
type Debug struct {
	Log bool `json:"log"`
	DB  bool `json:"db"`
}

type Security struct {
	AllowedAdminIP []string `json:"allowed_admin_ip"`
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

func Save(path string, config Main) error {
	configJson, _ := json.MarshalIndent(config, "", "  ")
	if err := ioutil.WriteFile(path, configJson, 0644); err != nil {
		return err
	}

	return nil
}

// DbConnURL create url to connect with database.
func (c *Main) DbConnURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.DB.User,
		c.DB.Password,
		c.DB.Host,
		c.DB.Port,
		c.DB.Name)
}
