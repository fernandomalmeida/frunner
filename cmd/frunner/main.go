package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fernandomalmeida/frunner"
	"gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Use with DIR argument: %s DIR\n", os.Args[0])
		os.Exit(1)
	}

	dir := os.Args[1]

	pipelineBytes, err := ioutil.ReadFile(filepath.Join(dir, ".frunner.yaml"))
	if err != nil {
		log.Fatalf("error on read yaml file: %s", err)
	}

	var pipeline frunner.Pipeline
	err = yaml.Unmarshal(pipelineBytes, &pipeline)
	if err != nil {
		log.Fatalf("error on read yaml file: %s", err)
	}

	pipeline.FillDir(dir)

	for i := 0; i < len(pipeline.Steps); i++ {
		step := pipeline.Steps[i]
		err := step.Run()
		if err != nil {
			log.Fatalf("error on step %s: %s", step.Name, err)
		}
	}

}
