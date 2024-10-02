package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
)

func main() {
	res := folder.GenerateData()
	folder.PrettyPrint(res)
	if err := folder.WriteSampleData(res); err != nil {
		fmt.Printf("Error writing sample data: %v\n", err)
	}
}
