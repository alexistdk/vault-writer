package main

import (
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
	cmValues := cm.getValues(configMaps)
	sliceOfScrets := make([]string, 0, 200)
	for i := 0; i < len(paths); i++ {
		// los secrets son strings
		secrets := service.getEnvVars(paths[i], cmValues)
		sliceOfScrets = append(sliceOfScrets, secrets)
	}
}
