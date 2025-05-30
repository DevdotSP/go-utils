package main

import (
	"fmt"
	"os"

	command "github.com/DevdotSP/go-utils/boilerplate/cmd"
)

func main() {
	//go run package/boilerplate/tool.go module customer
	//["gcloud", "database", "firebase"]
	//go run package/boilerplate/tool.go config database

	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  go run tool.go module [ModuleName]")
		fmt.Println("  go run tool.go config [ComponentName]")
		return
	}

	cmd := os.Args[1]
	arg := os.Args[2]

	switch cmd {
	case "module":
		command.GenerateModule(arg) // go run package/boilerplate/tool.go module customer
	case "config":
		command.GenerateConfig(arg) // go run package/boilerplate/tool.go config database
	default:
		fmt.Println("Unknown command:", cmd)
	}
}
