package business

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Ship struct {
	Name string `yaml:"ship_name"`
}

func (s *Service) getShips() ([]Ship, error) {
	// Read the YAML file
	yamlFile, err := os.ReadFile("ships.yaml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return nil, err
	}

	var ships []Ship

	// Unmarshal the YAML data into the Go struct
	if err := yaml.Unmarshal(yamlFile, &ships); err != nil {
		fmt.Println("Error unmarshaling YAML:", err)
		return nil, err
	}

	return ships, nil

}
