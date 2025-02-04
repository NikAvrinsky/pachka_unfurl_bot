package apiserver

import (
	"pachca-bot/cmd/jirapreview"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config      *Config
	logger      *logrus.Logger
	router      *mux.Router
	jirapreview *jirapreview.JiraPreview
}
