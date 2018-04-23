// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package setting

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

const (
	configFilePath = "./conf/app.yaml"
)

// Dev describe is now running under dev
var Dev bool

// App main app config definition. app.yaml will be mapped to it
var App struct {
	Vk struct {
		AppSecret string `yaml:"app_secret"`
	}
}

// NewContext opent conf file, parse it
func NewContext(configFilePaths ...string) (err error) {
	cnfFilePath := configFilePath
	if len(configFilePaths) > 0 {
		cnfFilePath = configFilePaths[0]
	}
	log.Printf("read config: %v", cnfFilePath)
	f, err := os.Open(cnfFilePath)
	if err != nil {
		return errors.Wrap(err, "open config file")
	}
	dec := yaml.NewDecoder(f)
	err = dec.Decode(&App)
	if err != nil {
		data, err := ioutil.ReadFile(cnfFilePath)
		if err != nil {
			return errors.Wrap(err, "ioutil.ReadFile")
		}
		log.Printf("cnf data: %s", data)
		return errors.Wrap(err, "decode config file")
	}
	return
}
