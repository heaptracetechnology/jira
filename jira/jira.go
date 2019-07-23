package jira

import (
	"encoding/json"
	//"fmt"
	//jira "github.com/andygrunwald/go-jira"
	jira "github.com/heaptracetechnology/jira-go-sdk"
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

//CreateProject Jira
func CreateProject(responseWriter http.ResponseWriter, request *http.Request) {

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

	projectDetails := jira.CreateProject{
		Key:            "AD",
		Name:           "AD Soft-ware",
		Description:    "Example Project description",
		ProjectTypeKey: "software",
		Lead:           "admin",
	}

	project, _, projectErr := client.Project.Create(&projectDetails)
	if projectErr != nil {
		result.WriteErrorResponseString(responseWriter, projectErr.Error())
		return
	}

	bytes, _ := json.Marshal(project)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
