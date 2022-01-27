package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var yml Yaml
	var service Service

	if len(os.Args) != 2 {
		log.Fatal("Tenes que pasarle un yaml")
	}

	paths := yml.getPaths(os.Args[1])
	for i := 0; i < len(paths); i++ {
		secrets := service.getEnvVars(paths[i])
		fmt.Println(paths[i])
		fmt.Println(secrets)
		fmt.Println("---------------------------------------------------------------")
	}
}
