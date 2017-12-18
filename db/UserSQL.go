package db




const (
	modelTable = " auth.user_model_models "
	
	modelColums = "identifier, email, regID"
	
	insertSQL = "INSERT INTO " + modelTable + " (" + modelColums + ") VALUES ($1, $2, $3)"
	
	updateSQL = "UPDATE " + modelTable + " SET regID = $1 WHERE identifier = $2"
	
	selectAndFrom = "SELECT " + modelColums + " FROM " + modelTable
	
	selectUserByLoginSQL = selectAndFrom + " WHERE email = $1"
	
	selectUseByRegIDSQL = selectAndFrom + " WHERE regID = $1"
	
	existsLoginValidationSQL = "SELECT EXISTS(SELECT 1 FROM " + modelTable + " WHERE email = $1)"

)