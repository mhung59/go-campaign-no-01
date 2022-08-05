package main

import (
	"fmt"

	services "github.com/mhung59/go-campaign/pkg/services"
)

func main() {
	fmt.Print(services.Hello("Hung"))
	services.GenAndExportUser(1, "Hung", 23, "Fresher", "Dak Nong")
}
