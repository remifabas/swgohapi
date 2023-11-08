package guild

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type DefaultService struct {
	filePath string
}

func (s *DefaultService) GetPlayers() ([]Player, error) {
	// Read the YAML file
	yamlFile, err := os.ReadFile(s.filePath)
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return nil, err
	}

	var items []Player

	// Unmarshal the YAML data into the Go struct
	if err := yaml.Unmarshal(yamlFile, &items); err != nil {
		fmt.Println("Error unmarshaling YAML:", err)
		return nil, err
	}

	return items, nil
}
