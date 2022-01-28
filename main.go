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

	yml.readYaml(os.Args[1])
	clusters := len(yml.Clusters)
	paths := make([]string, 0, 1000)

	for i := 0; i < clusters; i++ {
		ns := len(yml.Clusters[i].Namespaces)
		for j := 0; j < ns; j++ {
			namespaces := yml.Clusters[i].Namespaces[j].Name
			pathsConfigMaps := make([]string, 0, 10)
			yml.getFiles(i, j, &paths)
			yml.getConfigMaps(i, &pathsConfigMaps)
			cmValues := cm.getValues(pathsConfigMaps)
			sliceOfScrets := make([]string, 0, 200)
			for k := 0; k < len(paths); k++ {
				// los secrets son strings
				secrets := service.getEnvVars(paths[k], cmValues)
				sliceOfScrets = append(sliceOfScrets, secrets)
				fmt.Println(namespaces)
				/*
					fmt.Println(paths[k])
					fmt.Println(secrets)
					fmt.Println("---------------------------------------------------------------")
				*/
			}
		}
	}

}
