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
  password:  passwordd
  dbname: godb
  path: "/kjasd/sd"
server:
  globalauthcode: "hello"
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
	// if yS.Server.Globalauthcodes[0] != []string{"hello", "sadjksad"}[0] || yS.Server.Globalauthcodes[1] != []string{"hello", "sadjksad"}[1] {
	// 	t.Error("Couldn't read server authcodes type")
	// }
	if yS.Server.Usessl != true {
		t.Error("Couldn't read usessl")
	}

}
