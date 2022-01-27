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

func (yml *Yaml) getClusters() int { return len(yml.Clusters) }

func (yml *Yaml) getNamespaces(i int) int { return len(yml.Clusters[i].Namespaces) }

func (yml Yaml) getPaths(file string) []string {
	yml.readYaml(file)
	clusters := len(yml.Clusters)
	paths := make([]string, 0, 100)
	for i := 0; i < clusters; i++ {
		ns := len(yml.Clusters[i].Namespaces)
		for j := 0; j < ns; j++ {
			files := yml.Clusters[i].Namespaces[j].Services
			for k := 0; k < len(files); k++ {
				auxPath := fmt.Sprint(files[k].File)
				paths = append(paths, auxPath)
			}
		}
	}
	/*
		files := yml.Clusters[0].Namespaces[0].Services
		paths := make([]string, len(files))
		for i := 0; i < len(files); i++ {
			paths[i] = fmt.Sprint(files[i].File)
		}
	*/
	return paths
}
