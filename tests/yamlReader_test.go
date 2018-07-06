package tests

import (
	"testing"

	"../persistence"
)

var yamlString = `
database:
  adapter: mysql
  host: localhost
  port: 1111
  user: david
  password:  password
  dbname: godb
  path: "/kjasd/sd"
server:
  globalauthcodes: ["gs123","sadjksad"]
  usessl: true
`

func TestParseYAML(t *testing.T) {
	yS, err := persistence.ParseYAML(yamlString)
	//  fmt.Printf("--- t:\n%v\n\n", yS)
	if err != nil {
		t.Error("Couldn't unmarshal yaml", err)
	}
	if yS.Database.Adapter != "mysql" {
		t.Error("Couldn't read database type")
	}
	if yS.Server.Globalauthcodes[0] != []string{"gs123", "sadjksad"}[0] || yS.Server.Globalauthcodes[1] != []string{"gs123", "sadjksad"}[1] {
		t.Error("Couldn't read server authcodes type")
	}
	if yS.Server.Usessl != true {
		t.Error("Couldn't read usessl")
	}

}
