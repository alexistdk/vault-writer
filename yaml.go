package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Yaml struct {
	URL      interface{} `yaml:"url"`
	Clusters []struct {
		Name       interface{} `yaml:"name"`
		Namespaces []struct {
			Name     interface{} `yaml:"name"`
			Services []struct {
				File interface{} `yaml:"file"`
			} `yaml:"services"`
		} `yaml:"namespaces"`
	} `yaml:"clusters"`
}

func (yml *Yaml) readYaml(file string) *Yaml {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, yml)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return yml
}

func (yml Yaml) getPaths(file string) []string {
	yml.readYaml(file)
	files := yml.Clusters[0].Namespaces[0].Services
	paths := make([]string, len(files))
	for i := 0; i < len(files); i++ {
		paths[i] = fmt.Sprint(files[i].File)
	}
	return paths
}
