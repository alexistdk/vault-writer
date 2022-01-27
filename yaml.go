package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Yaml struct {
	URL      string `yaml:"url"`
	Clusters []struct {
		Name       string `yaml:"name"`
		Configmaps []struct {
			File string `yaml:"file"`
			Path string `yaml:"path"`
		} `yaml:"configmaps"`
		Namespaces []struct {
			Name     string `yaml:"name"`
			Services []struct {
				Path string `yaml:"path"`
				File string `yaml:"file"`
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

func (yml *Yaml) getFiles(i int, j int, pathsSlice *[]string) {
	keys := yml.Clusters[i].Namespaces[j].Services
	for k := 0; k < len(keys); k++ {
		if keys[k].File != "" {
			auxPath := fmt.Sprint(keys[k].File)
			*pathsSlice = append(*pathsSlice, auxPath)
		} else {
			auxPath := fmt.Sprint(keys[k].Path)
			getFilesInPath(auxPath, pathsSlice)
		}
	}
}

func getFilesInPath(path string, pathsSlice *[]string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		pathFile := filepath.Join(path, f.Name())
		*pathsSlice = append(*pathsSlice, pathFile)
	}
}

func (yml Yaml) getYamls(file string) []string {
	yml.readYaml(file)
	clusters := len(yml.Clusters)
	paths := make([]string, 0, 1000)
	for i := 0; i < clusters; i++ {
		ns := len(yml.Clusters[i].Namespaces)
		for j := 0; j < ns; j++ {
			yml.getFiles(i, j, &paths)
		}
	}
	return paths
}
