package api

import (
	"bytes"
	"log"
	"net/http"
)

type API struct {
	URL    string
	Agency string
	Token  string
}

func (a *API) getTokenString() string {
	return "Token " + a.Token
}

func (a *API) getRideURL() string {
	return a.URL + "/v1/ondemand/" + a.Agency + "/rides"
}

func (a *API) PostRide(payload []byte) *http.Response {
	token := a.getTokenString()
	req, err := http.NewRequest("POST", a.getRideURL(), bytes.NewBuffer(payload))
	req.Header.Add("Authorization", token)
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	return resp
}
