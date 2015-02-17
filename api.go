package main

import (
	"bufio"
	"code.google.com/p/gorest"
	"github.com/gorilla/schema"
	"log"
	"os/exec"
	"strings"
)

type ApiService struct {
	gorest.RestService `root:"/v1/" consumes:"application/json" produces:"application/json"`
	//Routes
	cpu gorest.EndPoint `method:"GET" path:"/cpu/" output:"[]Processor"`
}

func (serv ApiService) Cpu() []Processor {
	cmd := exec.Command("cat", "/proc/cpuinfo")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	buf := bufio.NewScanner(stdout)
	var values map[string]string
	processors := make([]Processor, 1)
	decoder := schema.NewDecoder()

	for buf.Scan() {
		line := buf.Text()
		if line == "" {
			//procs[proc["processor"].(string)] = proc
			proc := new(Processor)
			decoder.Decode(proc, values)
			processors = append(processors, proc)
			values := make(map[string]string)
		} else {
			split := strings.Split(line, ":")
			if len(split) > 1 {
				//proc[split[0]] = split[1]
				values[strings.Trim(split[0], "\t ")] = strings.Trim(split[1], "\t ")
			}
		}
		proc := new(Processor)
		decoder.Decode(proc, values)
		processors = append(processors, proc)
	}
	return processors
}
