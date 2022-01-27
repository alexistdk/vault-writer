package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ConfigMap struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Data interface{} `yaml:"data"`
}

func (cm *ConfigMap) getConfigMap(filePath string) *ConfigMap {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, cm)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return cm
}

func (cm *ConfigMap) getName() string { return cm.Metadata.Name }

func (cm *ConfigMap) getData() interface{} { return cm.Data }
