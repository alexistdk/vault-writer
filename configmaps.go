package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
)

type ConfigMap struct {
	Metadata struct {
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

func (cm *ConfigMap) getValue(requestedKey string) string {
	data := cm.getData()
	dict := reflect.ValueOf(data)
	if dict.Kind() == reflect.Map {
		for _, key := range dict.MapKeys() {
			if key.Interface() == requestedKey {
				value := dict.MapIndex(key).Interface()
				return fmt.Sprint(value)
			}
		}
	}
	return ""
}

func (cm *ConfigMap) getKeys() []reflect.Value {
	data := cm.getData()
	dict := reflect.ValueOf(data)
	return dict.MapKeys()
}

func (cm *ConfigMap) getValues(configmaps []string) []map[string]string {
	sliceEnvVars := make([]map[string]string, 0, 200)
	for i := 0; i < len(configmaps); i++ {
		config := cm.getConfigMap(configmaps[i])
		value := config.getValue("police_registration.exchange")
		fmt.Println(value)
	}
	return sliceEnvVars
}
