package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/nicklaw5/helix"
)

func healthCheckHandler(res http.ResponseWriter, req *http.Request) {
	// returns 200 OK by default
}

func twitchWebhookHandler(res http.ResponseWriter, req *http.Request) {
	if verbose {
		dumpReq, err := httputil.DumpRequest(req, true)
		if err != nil {
			panic(err)
		}
		ldebugf("%+v", string(dumpReq))
	}

	if req.Method == http.MethodGet {
		handleWebhookGetRequest(res, req)
	}

	if req.Method == http.MethodPost {
		handleWebhookPostRequest(res, req)
	}
}

func handleWebhookGetRequest(res http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["hub.challenge"]

	if !ok || len(keys[0]) < 1 {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if verbose {
		ldebugf("hub.challenge value: %s", keys[0])
	}

	// TODO: Verify subscription request is valid
	// See https://github.com/chatstatz/webhooks/issues/3
	res.Write([]byte(keys[0]))
}

func handleWebhookPostRequest(res http.ResponseWriter, req *http.Request) {
	linkHeader := req.Header.Get("link")
	ldebugf("Link Header: %s", linkHeader)

	webhookTopic := helix.GetWebhookTopicFromRequest(req)
	ldebugf("Topic: %+v", webhookTopic)

	webhookTopicValues := helix.GetWebhookTopicValuesFromRequest(req, -1)
	ldebugf("Topic Values: %+v", webhookTopicValues)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	event := &webhookEvent{
		Topic:       webhookTopic,
		TopicValues: webhookTopicValues,
		Payload:     string(body),
	}

	eventBytes, err := event.ToBytes()
	if err != nil {
		panic(err)
	}

	err = service.PublishMessage(eventBytes)
	if err != nil {
		panic(err)
	}

	ldebugf("%s: %s", "Message published", string(eventBytes))
}
