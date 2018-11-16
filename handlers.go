package main

import (
	"io/ioutil"
	"net/http"
)

func healthCheckHandler(res http.ResponseWriter, req *http.Request) {
}

func twitchWebhookHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte(`{"success":false,"message":"method not allowed"}`))
		return
	}

	if req.Body == nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"success":false,"message":"missing request body"}`))
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	err = service.PublishMessage(body)
	if err != nil {
		panic(err)
	}

	ldebugf("%s: %s", "Message published", string(body))

	res.Write([]byte(`{"success":true}`))
}
