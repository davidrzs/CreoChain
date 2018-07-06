package globalvariables

import (
	"sync"

	"../persistence"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ServerManager stores all relevant data of our runtime
type ServerManager struct {
	Mutex       *sync.Mutex
	Name        string
	BlockChains *gorm.DB
}

func DatabaseConnectionString(config *persistence.YAMLReader) {
	config := `
database: memory
server:
  globalauthcodes: ["gs123","sadjksad"]
  usessl: true
`

}
