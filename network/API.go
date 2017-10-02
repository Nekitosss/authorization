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


func Register(w http.ResponseWriter, r *http.Request) {

	var params structures.RegisterInfo

	var err = json.NewDecoder(r.Body).Decode(&params)

	if err != nil && err != io.EOF {
		setServerError(w, err)
		return
	}

	err = executor.Register(params)

	if err != nil {
		setServerError(w, err)
		return
	}

	sendSimpleSuccess(w)
}


func VerifyRegistration(w http.ResponseWriter, r *http.Request) {
	
	var registrationIDString = mux.Vars(r)["id"]
	
	var registrationID, err = uuid.FromString(registrationIDString)
	
	if err != nil {
		setServerError(w, err)
		return
	}
	
	err = executor.VerifyRegister(registrationID)

	if err != nil {
		setServerError(w, err)
		return
	}

	sendSimpleSuccess(w)
}



func Login(writer http.ResponseWriter, request *http.Request) {

	var params structures.LoginInfo

	err := json.NewDecoder(request.Body).Decode(&params)

	if err != nil && err != io.EOF {
		setServerError(writer, err)
		return
	}

	session, err := executor.Login(params)

	if err != nil {
		setServerError(writer, err)
		return
	}

	expiration := time.Now().AddDate(0, 1, 0)
	cookie := http.Cookie{Name: "session", Value: session.SessionID.String(), Expires: expiration}

	http.SetCookie(writer, &cookie)

	var result = serverResponse{"userID" : session.UserModelID.String()}
	json.NewEncoder(writer).Encode(result)
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
	successRes, err := json.Marshal(result)
	
	if err != nil {
		setServerError(w, err);
		return
	}
	
	json.NewEncoder(w).Encode(successRes)
}
