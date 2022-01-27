package main

import "fmt"

func main() {
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
