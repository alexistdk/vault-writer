package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

type Service struct {
	// para hacer el struct del yaml use https://zhwt.github.io/yaml-to-go/
	Spec struct {
		Template struct {
			Spec struct {
				Containers []struct {
					Env []struct {
						Name  string `yaml:"name"`
						Value string `yaml:"value"`
					} `yaml:"env"`
				} `yaml:"containers"`
			} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
}

func (service *Service) getService(filePath string) *Service {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, service)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return service
}

func (service Service) getEnvVars(filePath string) string {
	service.getService(filePath)
	containers := service.Spec.Template.Spec.Containers[0]
	env := containers.Env
	var str strings.Builder
	str.WriteString("{\n")
	for i := 0; i < len(env); i++ {
		if env[i].Value != "" {
			str.WriteString("\t\"" + env[i].Name + "\": \"" + env[i].Value + "\"\n")
		}
	}
	str.WriteString("}")
	return str.String()
}
