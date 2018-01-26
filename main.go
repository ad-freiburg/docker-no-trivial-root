package main

import (
	"log"

	"github.com/docker/go-plugins-helpers/authorization"
)

func main() {

	noroot, err := NewNoRootPlugin()
	if err != nil {
		log.Fatal(err)
	}

	h := authorization.NewHandler(noroot)

	if err := h.ServeUnix("no-trivial-root", 0); err != nil {
		log.Fatal(err)
	}
}
