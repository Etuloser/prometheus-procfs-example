package main

import (
	"fmt"

	"github.com/prometheus/procfs"
	log "github.com/sirupsen/logrus"
)

func main() {
	proc, err := procfs.NewProc(846)
	if err != nil {
		log.Error(err)
	}
	maps, err := proc.ProcMaps()
	if err != nil {
		log.Error(err)
	}
	for k, v := range maps {
		fmt.Println(k, v)
	}
}
