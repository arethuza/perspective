package misc

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"path"
)

type Config struct {
	DbConnection   	string
	Port           	int
	BcryptCost     	int
	PasswordLength 	int
	TokenLength    	int
	RedisHost 	string
	RedisPort	int
	RedisExpiration	string
}

func LoadConfig() (*Config, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	filepath := path.Join(dir, "config.tml")
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	s := string(data)
	config := &Config{}
	_, err = toml.Decode(s, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
