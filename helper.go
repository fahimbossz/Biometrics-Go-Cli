package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const FILE_NAME = "STORAGE.txt"

func getFileContents() string {
	dat, err := ioutil.ReadFile(FILE_NAME)

	if err != nil {
		fmt.Printf("Error: %v", err.Error())
	}

	dataString := string(dat)

	return dataString
}

func getPresents() []Present {

	var presents []Present
	fileContentsArray := split(getFileContents(), "\n")

	for _, content := range fileContentsArray {
		if len(content) > 1 {
			// content: 0;1-January-2000;20:42:27;2015338501
			// fmt.Println("Content: " + content)
			// fmt.Printf("Present: %+v\n", getPresentModelFromString(content))
			presents = append(presents, getPresentModelFromString(content))
		}
	}
	return presents
}



func getPresentModelFromString(presentString string) Present {
	var present Present

	presentStringArr := split(presentString, ";")

	for i, presentValue := range presentStringArr {
		if i == 0 {
			present.CourseName = "EEE 101"
		} else if i == 1 {
			present.Date = presentValue
		} else if i == 2 {
			present.Time = presentValue
		} else if i == 3 {
			present.Roll = presentValue
		}
		// fmt.Printf("i: %v, val: %v\n", i, presentValue)
	}

	return present
}


func split(arr string, separator string) []string {
	slc := strings.Split(arr, separator)
	for i := range slc {
		slc[i] = strings.TrimSpace(slc[i])
	}
	return slc
}