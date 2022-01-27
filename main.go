package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var yml Yaml
	var service Service
	var cm ConfigMap

	if len(os.Args) != 2 {
		log.Fatal("Tenes que pasarle un yaml")
	}

	paths, configMaps := yml.getYamls(os.Args[1])
	sliceOfScrets := make([]string, 0, 200)
	for i := 0; i < len(paths); i++ {
		secrets := service.getEnvVars(paths[i])
		sliceOfScrets = append(sliceOfScrets, secrets)
		fmt.Println(paths[i])
		fmt.Println(secrets)
		fmt.Println("---------------------------------------------------------------")
	}

	for i := 0; i < len(configMaps); i++ {
		config := cm.getConfigMap(configMaps[i])
		value := config.getValue("police_registration.exchange")
		fmt.Println(value)
	}

}
