package apiserver

import (
	"encoding/json"
	"net/http"
	"pachca-bot/cmd/jirapreview"
	"pachca-bot/cmd/reminder"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// New ...
func New(config *Config) *ApiServer {
	return &ApiServer{
		config:      config,
		logger:      logrus.New(),
		router:      mux.NewRouter(),
		jirapreview: jirapreview.New(),
	}
}

// Start ...

func (c *ApiServer) Start() error {
	if err := c.configureLogger(); err != nil {
		return err
	}
	c.configureRouter()
	c.logger.Info("starting server")

	return http.ListenAndServe(c.config.BindAddr, c.router)
}

// Config logger
func (c *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(c.config.LogLevel)
	if err != nil {
		return err
	}
	c.logger.SetLevel(level)
	return nil
}

// Error wrapper
func (c *ApiServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	c.respond(w, r, code, map[string]string{"error": err.Error()})
}

// Respond ...
func (c *ApiServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// Config Router
func (c *ApiServer) configureRouter() {
	webhookPath := "/webhooks/" + "01JACZW1J0WX93SVKD1GAXJP27"
	c.router.HandleFunc(webhookPath, c.handleWebhook())
	reminderPath := "/reminder/" + "01JACZW1J0WX93SVKD1GAXJP27"
	c.router.HandleFunc(reminderPath, c.handleReminder())
}

func (c *ApiServer) handleWebhook() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := c.jirapreview.Handler(r.Body)
		if err != nil {
			c.logger.Error(err.Error())
			c.error(w, r, http.StatusInternalServerError, err)
		}
		c.respond(w, r, http.StatusOK, resp)
	}
}

func (c *ApiServer) handleReminder() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := reminder.Handler(r.Body)
		if err != nil {
			c.logger.Error(err.Error())
			c.error(w, r, http.StatusInternalServerError, err)
		}
		c.respond(w, r, http.StatusOK, nil)
	}
}
