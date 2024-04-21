package models

import (
	"log"
	"os"

	"github.com/hpinc/go3mf"
	"github.com/hpinc/go3mf/importer/stl"
)

func ConvertToSTL(filepath string) (err error) {
	// Create a new 3mf container
	model := new(go3mf.Model)

	// Open and decode the STL into the container
	r, err := os.Open(filepath)
	if err != nil {
		// log.Fatalf("Failed to open file with %v", err.Error())
		return
	}

	d := stl.NewDecoder(r)
	d.Decode(model)
	if err != nil {
		// log.Fatalf("Failed to decode with %v", err.Error())
		return
	}
	r.Close()

	log.Printf("Writing %s ...\n", filepath+".3mf")

	// Write a new file out from the decoded information
	w, _ := go3mf.CreateWriter(filepath + ".3mf")
	w.Encode(model)
	w.Close()

	return
}
