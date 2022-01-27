package main

import "fmt"

func main() {
	// var service Service
	// secrets := service.getEnvVars(filePath)
	// fmt.Println(secrets)
	var yml Yaml
	var service Service

	paths := yml.getPaths()
	for i := 0; i < len(paths); i++ {
		secrets := service.getEnvVars(paths[i])
		fmt.Println(paths[i])
		fmt.Println(secrets)
		fmt.Println("---------------------------------------------------------------")
	}
}
