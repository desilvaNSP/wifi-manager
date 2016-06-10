package handlers

import (
    "wislabs.wifi.manager/commons"
    "wislabs.wifi.manager/authenticator"
    "encoding/json"
    "net/http"
    "github.com/Sirupsen/logrus"
)

func Login(w http.ResponseWriter, r *http.Request) {
    requestUser := new(commons.SystemUser)
    decoder := json.NewDecoder(r.Body)
    decoder.Decode(&requestUser)

    responseStatus, token := authenticator.Login(requestUser)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(responseStatus)
    w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    requestUser := new(commons.SystemUser)
    decoder := json.NewDecoder(r.Body)
    decoder.Decode(&requestUser)

    w.Header().Set("Content-Type", "application/json")
    w.Write(authenticator.RefreshToken(requestUser))
}

func Logout(w http.ResponseWriter, r *http.Request) {
    err := authenticator.Logout(r)
    if err != nil {
	logrus.Error(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
    } else {
	w.WriteHeader(http.StatusOK)
    }
}
