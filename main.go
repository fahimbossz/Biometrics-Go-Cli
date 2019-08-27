package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

// 0;1-January-2000;20:42:27;2015338501
// 0;1-January-2000;20:42:31;2015338504

func main() {

	presents := getPresents()

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()


	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, present := range presents {
		row = sheet.AddRow()

		cell = row.AddCell()
		cell.Value = present.CourseName

		cell = row.AddCell()
		cell.Value = present.Date

		cell = row.AddCell()
		cell.Value = present.Time

		cell = row.AddCell()
		cell.Value = present.Roll
	}


	err = file.Save("Exported/MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("Presents: %+v\n", presents)

	// fmt.Println(fileContentsArray)


}
