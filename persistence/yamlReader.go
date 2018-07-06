package persistence

import (
	"gopkg.in/yaml.v2"
)

//YAMLReader is used to unmarshal yaml file.
type YAMLReader struct {
	Database struct {
		Adapter  string
		Host     string
		Port     string
		User     string
		Password string
		Dbname   string
		Path     string
	}
	Server struct {
		Globalauthcodes []string `yaml:",flow"`
		Urls            []string `yaml:",flow"`
		Usessl          bool
	}
}

// ParseYAML parses a string of yaml.
func ParseYAML(yamlS string) (*YAMLReader, error) {
	yamlreader := YAMLReader{}

	err := yaml.Unmarshal([]byte(yamlS), &yamlreader)

	return &yamlreader, err
}
