package globalvariables

import (
	"sync"

	"../persistence"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ServerManager stores all relevant data of our runtime
type ServerManager struct {
	Mutex    *sync.Mutex
	Name     string
	Database *gorm.DB
	Config   *persistence.YAMLReader
}

func DatabaseConnectionString(config *persistence.YAMLReader) (string, string) {
	if config.Database.Adapter == "mysql" {
		return "mysql", config.Database.User + ":" + config.Database.Password + "@/" + config.Database.Dbname + "?charset=utf8&parseTime=True&loc=Local"
	} else if config.Database.Adapter == "sqlite" {
		return "sqlite3", config.Database.Path
	} else {
		panic("No database adapter specified or an unavailable database adapter has been set. Set the database adapter in the config.yml file.")
	}

}
