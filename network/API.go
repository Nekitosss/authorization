package network

import (
	"github.com/Nekitosss/authorization/network/controller/structures"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)


type serverResponse map[string]interface{}


func AuthVerifiedUser(writer http.ResponseWriter, request *http.Request) {

	var params structures.AuthInfo

	err := json.NewDecoder(request.Body).Decode(&params)

	if err != nil && err != io.EOF {
		setErrorResult(writer, err)
		return
	}

	session, err := executor.AuthUser(params)

	if err != nil {
		setErrorResult(writer, err)

	} else {
		expiration := time.Now().AddDate(10, 0, 0)
		cookie := http.Cookie{Name: "session", Value: session.SessionID.String(), Expires: expiration}
		http.SetCookie(writer, &cookie)

		var result = serverResponse{"userID" : session.UserModelID.String()}
		json.NewEncoder(writer).Encode(result)
	}
}


func VerifyRegistration(w http.ResponseWriter, r *http.Request) {
	
	var registrationIDString = mux.Vars(r)["id"]
	
	var registrationID, err = uuid.FromString(registrationIDString)
	
	if err != nil {
		setErrorResult(w, err)
		return
	}
	
	err = executor.VerifyRegister(registrationID)

	if err != nil {
		setErrorResult(w, err)
		return
	}

	var html = `
	<!DOCTYPE html>
	<html>
	<head>
	<title>Success confirmation</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<script type="text/javascript">
		function codeAddress() {
			var iOS = !!navigator.platform && /iPad|iPhone|iPod/.test(navigator.platform);
			if (iOS) {
				window.location.replace("fooddly://authorization");
			}
		}
	window.onload = codeAddress;
	</script>
	</head>
	<body>
	<p>Ok result, you should go back to application</p> 
	</body>
	</html>
	`
	
	r.Write(html)
}



func Login(writer http.ResponseWriter, request *http.Request) {

	var params structures.LoginInfo

	err := json.NewDecoder(request.Body).Decode(&params)

	if err != nil && err != io.EOF {
		setErrorResult(writer, err)
		return
	}

	registrationID, err := executor.Login(params)

	if err != nil {
		setErrorResult(writer, err)

	} else {
		result := serverResponse{"registrationID": registrationID.String()}
		json.NewEncoder(writer).Encode(result)
	}
}



func ValidateSession(writer http.ResponseWriter, request *http.Request) {

	var info structures.ValidateSessionInfo
	err := json.NewDecoder(request.Body).Decode(&info)

	if err != nil {
		setErrorResult(writer, err)
		return
	}

	userID, err := executor.ValidateSession(info)

	if err != nil {
		setErrorResult(writer, err)
		return
	}

	var result = serverResponse{"userID": userID.String()}
	json.NewEncoder(writer).Encode(&result)
}



func sendSimpleSuccess(w http.ResponseWriter) {
	result := serverResponse{"success": true}

	json.NewEncoder(w).Encode(result)
}
