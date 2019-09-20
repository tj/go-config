// Package config provides super simple JSON configuration for command-line programs.
package config

import (
	"encoding/json"
	"io/ioutil"
)

// Load configuration from path, into the struct pointer provided.
func Load(path string, v interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

// Save saves configuration to path.
func Save(path string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, b, 0600)
}
