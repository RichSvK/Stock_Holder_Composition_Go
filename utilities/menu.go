package utilities

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/RichSvK/Stock_Holder_Composition_Go/configs"
	"github.com/RichSvK/Stock_Holder_Composition_Go/helpers"
	"github.com/RichSvK/Stock_Holder_Composition_Go/services"
)

func LoginMenu() {
	var (
		username string = ""
		password string = ""
		dbName   string = ""
	)

	var db *sql.DB = nil
	for db == nil {
		fmt.Println("Login Menu")
		username = helpers.ScanInput("Insert username: ")
		password = helpers.ScanInput("Insert password: ")
		dbName = helpers.ScanInput("Insert Database name: ")
		db = configs.GetConnection(username, password, dbName)
		helpers.PressEnter()
		helpers.ClearScreen()
	}
}

func MainMenu() int {
	var userInput string = ""
	var choice int = 0
	var err error = nil
	helpers.ClearScreen()
	fmt.Println("Main Menu")
	fmt.Println("1. Insert Data to Database")
	fmt.Println("2. Export Data to Database")
	fmt.Println("3. Exit")
	for {
		userInput = helpers.ScanInput("Input[1 - 3]: ")
		choice, err = strconv.Atoi(userInput)
		if err == nil && choice >= 1 && choice <= 3 {
			break
		}
	}
	return choice
}

func ExportMenu() {
	var code string = ""

	for len(code) != 4 {
		code = helpers.ScanInput("Input stock name [4 Letter]: ")
	}
	services.Export(code)
}

func InsertMenu() {
	helpers.ClearScreen()
	fmt.Println("Menu Insert")

	// Change this to the directory you want to list files from
	directory := "data/"

	// Initialize an empty slice to hold file paths
	var fileList []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// Add the file path to the slice
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var size int = len(fileList)
	var tempName []string = nil
	for i := 0; i < size; i++ {
		tempName = strings.Split(fileList[i], "data\\")
		fmt.Printf("%d. %s from %s\n", (i + 1), tempName[1], fileList[i])
	}

	var choice int = 0
	var userInput string = ""
	promptString := fmt.Sprintf("Input [1 - %d]: ", size)
	for {
		userInput = helpers.ScanInput(promptString)
		choice, err = strconv.Atoi(userInput)
		if err == nil && choice >= 1 && choice <= size {
			break
		} else {
			fmt.Println("Invalid Input")
		}
	}
	services.InsertData(fileList[choice-1])
}
