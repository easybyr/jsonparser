package jsonparser

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// Config struct is to read from json file to maps.
type Config struct {
	file string
	maps map[string]interface{}
}


// New to init Config struct.
func New(file string) *Config {
	return &Config{file: file}
}


// Get to get key from json.
func (c *Config) Get(key string) interface{} {
	if c.maps == nil {
		c.load()
	}
	if c.maps == nil {
		return nil
	}

	var keys []string = strings.Split(key, ".")
	idx := len(keys)

	var ret interface{}
	for i := 0; i < idx; i++ {
		if i == 0 {
			ret = c.maps[keys[i]]
			if ret == nil {
				return nil
			}
		} else {
			if m, ok := ret.(map[string]interface{}); ok {
				ret = m[keys[i]]
			} else {
				return nil
			}
		}
	}

	return ret
}

// Load to read json file
func (c *Config) load() {
	data, err := ioutil.ReadFile(c.file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &c.maps)
	if err != nil {
		panic(err)
	}
}
