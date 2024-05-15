package main

import (
	"fmt"

	"github.com/RichSvK/Stock_Holder_Composition_Go/configs"
	"github.com/RichSvK/Stock_Holder_Composition_Go/helpers"
	"github.com/RichSvK/Stock_Holder_Composition_Go/utilities"
)

func init() {
	configs.MakeFolder("output")
	utilities.LoginMenu()
}

func main() {
	// Close database when the main function end
	defer configs.PoolDB.Close()

	var choice int = 0
	for choice != 3 {
		choice = utilities.MainMenu()
		switch choice {
		case 1:
			utilities.InsertMenu()
		case 2:
			utilities.ExportMenu()
		default:
			fmt.Println("Program finished")
			return
		}
		helpers.PressEnter()
		helpers.ClearScreen()
	}
}
