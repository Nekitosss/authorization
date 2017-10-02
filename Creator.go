package authorization

import (
	"github.com/Nekitosss/authorization/network"
	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)


func PrepareToWork(database *gorm.DB, configuration network.EmailConfiguration, router *mux.Router) {

	network.PrepareExecutor(database, configuration)

	addRoutes(router)
}
