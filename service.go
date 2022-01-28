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
					Name string `yaml:"name"`
					Env  []struct {
						Name      string `yaml:"name"`
						Value     string `yaml:"value,omitempty"`
						ValueFrom struct {
							ConfigMapKeyRef struct {
								Name string `yaml:"name"`
								Key  string `yaml:"key"`
							} `yaml:"configMapKeyRef"`
						} `yaml:"valueFrom,omitempty"`
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

func getEnvVarFromConfigMap(requestedKey string, dataConfigMap []map[string]string) string {
	for cm := range dataConfigMap {
		dict := dataConfigMap[cm]
		if v, found := dict[requestedKey]; found {
			return v
		}
	}
	return ""
}

func (service Service) getEnvVars(filePath string, dataConfigMap []map[string]string) string {
	service.getService(filePath)
	// TODO: identificar cronjobs que no tienen containers
	containers := service.Spec.Template.Spec.Containers[0]
	env := containers.Env
	var str strings.Builder
	str.WriteString("{\n")
	for i := 0; i < len(env); i++ {
		name := env[i].Name
		configmap := env[i].ValueFrom.ConfigMapKeyRef
		if (env[i].ValueFrom.ConfigMapKeyRef) == configmap {
			if configmap.Name != "" && configmap.Key != "" {
				key := getEnvVarFromConfigMap(configmap.Key, dataConfigMap)
				str.WriteString("\t\"" + name + "\": \"" + key + "\"\n")
			}
		}
		if env[i].Value != "" {
			str.WriteString("\t\"" + name + "\": \"" + env[i].Value + "\"\n")
		}

	}
	str.WriteString("}")
	return str.String()
}
