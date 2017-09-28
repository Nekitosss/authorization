package authorization

import (
	"authorization/network"
	"database/sql"

	"github.com/gorilla/mux"
)


func PrepareToWork(database *sql.DB, configuration network.EmailConfiguration, router *mux.Router) {

	network.PrepareExecutor(database, configuration)

	addRoutes(router)
}
