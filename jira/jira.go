package jira

import (
	"encoding/json"
	jira "github.com/andygrunwald/go-jira"
	result "github.com/heaptracetechnology/jira/result"
	"net/http"
	"os"
)

//JiraArguments struct
type JiraArguments struct {
	Username       string `json:"username,omitempty"`
	AssigneeName   string `json:"assigneeName,omitempty"`
	AssigneeEmail  string `json:"assigneeEmail,omitempty"`
	Description    string `json:"description,omitempty"`
	Type           string `json:"type,omitempty"`
	ProjectKey     string `json:"projectKey,omitempty"`
	Summary        string `json:"summary,omitempty"`
	ProjectName    string `json:"name,omitempty"`
	ProjectTypeKey string `json:"typeKey,omitempty"`
	ProjectLead    string `json:"lead,omitempty"`
	IssueID        string `json:"issueId,omitempty"`
}

//GetClient Basic Auth Transport
func GetClient(email, apiToken, baseURL string) (*jira.Client, error) {

	basicAuthTransport := jira.BasicAuthTransport{
		Username: email,
		Password: apiToken,
	}

	client, err := jira.NewClient(basicAuthTransport.Client(), baseURL)
	if err != nil {
		return nil, err
	}

	return client, nil
}

//GetAccountDetails Jira
func GetAccountDetails(responseWriter http.ResponseWriter, request *http.Request) {

	email := os.Getenv("EMAIL")
	apiToken := os.Getenv("API_TOKEN")
	baseURL := os.Getenv("BASE_URL")

	decoder := json.NewDecoder(request.Body)
	var param JiraArguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client, err := GetClient(email, apiToken, baseURL)
	if err != nil {
		result.WriteErrorResponseString(responseWriter, err.Error())
		return
	}

	userDetails, _, userErr := client.User.Get(param.Username)
	if userErr != nil {
		result.WriteErrorResponseString(responseWriter, userErr.Error())
		return
	}

	bytes, _ := json.Marshal(userDetails)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//CreateIssue Jira
func CreateIssue(responseWriter http.ResponseWriter, request *http.Request) {

	email := os.Getenv("EMAIL")
	apiToken := os.Getenv("API_TOKEN")
	baseURL := os.Getenv("BASE_URL")

	decoder := json.NewDecoder(request.Body)
	var param JiraArguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client, err := GetClient(email, apiToken, baseURL)
	if err != nil {
		result.WriteErrorResponseString(responseWriter, err.Error())
		return
	}

	issueDetails := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				Name:         param.AssigneeName,
				EmailAddress: param.AssigneeEmail,
			},
			Description: param.Description,
			Type: jira.IssueType{
				Name: param.Type,
			},
			Project: jira.Project{
				Key: param.ProjectKey,
			},
			Summary: param.Summary,
		},
	}

	issue, _, issueErr := client.Issue.Create(&issueDetails)
	if issueErr != nil {
		result.WriteErrorResponseString(responseWriter, issueErr.Error())
		return
	}

	bytes, _ := json.Marshal(issue)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//ListProject Jira
func ListProject(responseWriter http.ResponseWriter, request *http.Request) {

	email := os.Getenv("EMAIL")
	apiToken := os.Getenv("API_TOKEN")
	baseURL := os.Getenv("BASE_URL")

	decoder := json.NewDecoder(request.Body)
	var param JiraArguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client, err := GetClient(email, apiToken, baseURL)
	if err != nil {
		result.WriteErrorResponseString(responseWriter, err.Error())
		return
	}

	req, _ := client.NewRequest("GET", "rest/api/2/project", nil)

	projects := new([]jira.Project)
	_, err = client.Do(req, projects)
	if err != nil {
		panic(err)
	}

	bytes, _ := json.Marshal(projects)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//GetIssue Jira
func GetIssue(responseWriter http.ResponseWriter, request *http.Request) {

	email := os.Getenv("EMAIL")
	apiToken := os.Getenv("API_TOKEN")
	baseURL := os.Getenv("BASE_URL")

	decoder := json.NewDecoder(request.Body)
	var param JiraArguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client, err := GetClient(email, apiToken, baseURL)
	if err != nil {
		result.WriteErrorResponseString(responseWriter, err.Error())
		return
	}

	issue, _, _ := client.Issue.Get(param.IssueID, nil)

	bytes, _ := json.Marshal(issue)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
