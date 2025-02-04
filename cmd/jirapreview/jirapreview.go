package jirapreview

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func New() *JiraPreview {
	return &JiraPreview{
		config: *NewConfig(),
		logger: logrus.New(),
	}
}

// Process JiraPreview event
func (j *JiraPreview) Handler(body io.ReadCloser) ([]byte, error) {
	req := &Request{}
	resp := &Response{
		LinkPreviews: make(map[string]LinkPreview),
	}
	issue := &Issue{}
	linkprev := &LinkPreview{}

	if err := json.NewDecoder(body).Decode(&req); err != nil {
		return nil, err
	}
	j.logger.Info("Processing new link: " + req.Links[0].Url)

	var issueId string = ""

	if strings.Contains(req.Links[0].Url, "browse") {
		issueId = strings.SplitN(req.Links[0].Url, "/", -1)[4]
	} else {
		issueId = strings.SplitN(req.Links[0].Url, "=", 2)[1]
	}

	iss, err := j.JiraGetIssue(issueId)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(iss, &issue)
	if err != nil {
		return nil, err
	}

	linkprev.Title = fmt.Sprintf("%s: %s", issue.Key, issue.Fields.Summary)
	linkprev.Description = fmt.Sprintf("Issue type: %s \t Priority: %s \n  Status: %s \n Assignee: %s ", issue.Fields.IssueType.Name, issue.Fields.Priority.Name, issue.Fields.Status.Name, issue.Fields.Assignee.DisplayName)
	resp.LinkPreviews[req.Links[0].Url] = *linkprev

	respJson, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	respUpd, err := j.UpdateJiraPreview(respJson, req.Message_id)
	if err != nil {
		return nil, err
	}
	return respUpd, err
}

// Get Issue from Jira
func (j *JiraPreview) JiraGetIssue(issueId string) ([]byte, error) {
	j.logger.Info("Get issue from Jira " + issueId)
	requestURL := j.config.JiraAPIUrl + issueId
	login := j.config.JiraLogin
	token := j.config.JiraToken
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(login, token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}

// Upload previews to Pachca
func (j *JiraPreview) UpdateJiraPreview(request []byte, messageId int) ([]byte, error) {
	j.logger.Info("Updating message...")
	requestURL := j.config.PachcaAPIUnfurl + strconv.Itoa(messageId) + "/link_previews"
	// userId := j.config.PachcaUserId
	accessToken := "Bearer " + j.config.PachcaToken

	req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewReader(request))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", accessToken)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	j.logger.Info("Update preview ststus: ", strconv.Itoa(res.StatusCode))
	j.logger.Info((string(resp)))
	defer res.Body.Close()

	return resp, nil
}
