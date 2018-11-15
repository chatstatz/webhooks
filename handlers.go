package main

import (
	"io/ioutil"
	"net/http"
)

func healthCheckHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte(`{"success":false,"message":"method not allowed"}`))
		return
	}

	res.Write([]byte(`{"success":true}`))
	res.WriteHeader(http.StatusOK)
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
		lerror(err)

		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"success":false,"message":"failed to read request body"}`))
		return
	}
	defer req.Body.Close()

	err = service.PublishMessage(body)
	if err != nil {
		lerror(err)

		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"success":false,"message":"failed to publish message"}`))
		return
	}

	ldebugf("%s: %s", "Message published", string(body))

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(`{"success":true}`))
}
