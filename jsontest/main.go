package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Address represents the address information in the JSON.
type Address struct {
	StreetAddress string `json:"streetAddress"`
	City          string `json:"city"`
	State         string `json:"state"`
	PostalCode    string `json:"postalCode"`
}

// Grid represents a grid in the JSON.
type Grid struct {
	GridName  string  `json:"gridName"`
	GridCoord float64 `json:"gridCoord"`
	GridAngle float64 `json:"gridAngle"`
}

// Project represents the overall structure of the JSON.
type Project struct {
	ProjectName string  `json:"projectName"`
	Address     Address `json:"address"`
	Grids       []Grid  `json:"grids"`
}

func main() {
	// Read JSON file
	filePath := "project.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal JSON data into struct
	var project Project
	err = json.Unmarshal(data, &project)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the struct
	fmt.Printf("Project Name: %s\n", project.ProjectName)
	fmt.Printf("Address: %s, %s, %s %s\n", project.Address.StreetAddress, project.Address.City, project.Address.State, project.Address.PostalCode)

	fmt.Println("Grids:")
	for _, grid := range project.Grids {
		fmt.Printf("  Grid Name: %s, Coord: %.3f, Angle: %.3f\n", grid.GridName, grid.GridCoord, grid.GridAngle)
	}
}
