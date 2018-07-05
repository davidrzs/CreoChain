package globalvariables

import (
	"sync"

	"github.com/jinzhu/gorm"
)

// ServerManager stores all relevant data of our runtime
type ServerManager struct {
	Mutex       *sync.Mutex
	Name        string
	BlockChains *gorm.DB
}
