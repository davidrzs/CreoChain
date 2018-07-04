package tests

import (
	"fmt"
	"testing"

	"../persistence"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

var yamlString = `
database: memory
server:
  authcodes: ["gs123","sadjksad"]
  usessl: true
`

func TestParseYAML(t *testing.T) {
	yS, err := persistence.ParseYAML(yamlString)
	fmt.Printf("--- t:\n%v\n\n", yS)
	if err != nil {
		t.Error("Couldn't unmarshal yaml", err)
	}
	if yS.Database != "memory" {
		t.Error("Couldn't read database type")
	}
	if yS.Server.Authcodes[0] != []string{"gs123", "sadjksad"}[0] || yS.Server.Authcodes[1] != []string{"gs123", "sadjksad"}[1] {
		t.Error("Couldn't read server authcodes type")
	}
	if yS.Server.Usessl != true {
		t.Error("Couldn't read usessl")
	}

}
