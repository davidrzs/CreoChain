package persistence

import (
	"gopkg.in/yaml.v2"
)

//YAMLReader is used to unmarshal yaml file.
type YAMLReader struct {
	Database string
	Server   struct {
		Authcodes []string `yaml:",flow"`
		Usessl    bool
	}
}

// ParseYAML parses a string of yaml.
func ParseYAML(yamlS string) (*YAMLReader, error) {
	yamlreader := YAMLReader{}

	err := yaml.Unmarshal([]byte(yamlS), &yamlreader)
	if err != nil {
		return &yamlreader, err
	}

	return &yamlreader, nil
}
