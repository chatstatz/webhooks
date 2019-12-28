package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/nicklaw5/helix"
)

func (s *Server) healthCheckHandler(res http.ResponseWriter, req *http.Request) {
	// returns 200 OK by default
}

func (s *Server) twitchWebhookHandler(res http.ResponseWriter, req *http.Request) {
	dumpReq, err := httputil.DumpRequest(req, true)
	if err != nil {
		panic(err)
	}
	s.ctx.Logger.Debugf("%+v", string(dumpReq))

	if req.Method == http.MethodGet {
		s.handleWebhookGetRequest(res, req)
	}

	if req.Method == http.MethodPost {
		s.handleWebhookPostRequest(res, req)
	}
}

func (s *Server) handleWebhookGetRequest(res http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["hub.challenge"]

	if !ok || len(keys[0]) < 1 {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	s.ctx.Logger.Debugf("hub.challenge value: %s", keys[0])

	// TODO: Verify subscription request is valid
	// See https://github.com/chatstatz/webhooks/issues/3
	res.Write([]byte(keys[0]))
}

func (s *Server) handleWebhookPostRequest(res http.ResponseWriter, req *http.Request) {
	linkHeader := req.Header.Get("link")
	s.ctx.Logger.Debugf("Link Header: %s", linkHeader)

	webhookTopic := helix.GetWebhookTopicFromRequest(req)
	s.ctx.Logger.Debugf("Topic: %+v", webhookTopic)

	webhookTopicValues := helix.GetWebhookTopicValuesFromRequest(req, -1)
	s.ctx.Logger.Debugf("Topic Values: %+v", webhookTopicValues)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	event := &WebhookEvent{
		Topic:       webhookTopic,
		TopicValues: webhookTopicValues,
		Payload:     string(body),
	}

	eventBytes, err := event.ToBytes()
	if err != nil {
		panic(err)
	}

	err = s.ctx.Producer.PublishMessage(eventBytes)
	if err != nil {
		panic(err)
	}

	s.ctx.Logger.Debugf("%s: %s", "Message published", string(eventBytes))
}
