package main

import (
	"fmt"
	"os"
)

func main() {

	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("DOMAIN")
	processorArch := os.Getenv("PROCESSOR_ARCHITECTURE")
	processorIdentifier := os.Getenv("PROCESSOR_IDENTIFIER")
	processorLevel := os.Getenv("PROCESSOR_LEVEL")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	fmt.Println(uName, uDomain, processorArch, processorIdentifier, processorLevel, goRoot, goPath, homePath)

}
