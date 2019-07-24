package service

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	email         = os.Getenv("JIRA_EMAIL")
	apiToken      = os.Getenv("JIRA_API_TOKEN")
	baseURL       = os.Getenv("JIRA_BASE_URL")
	username      = os.Getenv("JIRA_USERNAME")
	issueID       = os.Getenv("JIRA_ISSUE_ID")
	assigneeName  = os.Getenv("JIRA_ASSIGNEE_NAME")
	assigneeEmail = os.Getenv("JIRA_ASSIGNEE_EMAIL")
	projectKey    = os.Getenv("JIRA_PROJECT_KEY")
)

var _ = Describe("Get Account Details without environment variables", func() {

	jira := JiraArguments{Username: username}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(jira)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/getAccount", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAccountDetails)
	handler.ServeHTTP(recorder, request)

	Describe("Get Account", func() {
		Context("get account details", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List all projects without environment variables", func() {

	jira := JiraArguments{}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(jira)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/listProject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListProject)
	handler.ServeHTTP(recorder, request)

	Describe("List Project", func() {
		Context("list all project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get issue by ID without environment variables", func() {

	jira := JiraArguments{IssueID: issueID}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(jira)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/getIssue", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetIssue)
	handler.ServeHTTP(recorder, request)

	Describe("Get Issue", func() {
		Context("get issue by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create issue without environment variables", func() {

	jira := JiraArguments{AssigneeName: assigneeName}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(jira)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/createIssue", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateIssue)
	handler.ServeHTTP(recorder, request)

	Describe("Create Issue", func() {
		Context("create issue in project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get Account Details with invalid param", func() {

	os.Setenv("EMAIL", email)
	os.Setenv("API_TOKEN", apiToken)
	os.Setenv("BASE_URL", baseURL)

	jira := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(jira)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/getAccount", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAccountDetails)
	handler.ServeHTTP(recorder, request)

	Describe("Get Account", func() {
		Context("get account details", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get Account Details", func() {

	os.Setenv("EMAIL", email)
	os.Setenv("API_TOKEN", apiToken)
	os.Setenv("BASE_URL", baseURL)

	jira := JiraArguments{Username: username}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(jira)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/getAccount", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAccountDetails)
	handler.ServeHTTP(recorder, request)

	Describe("Get Account", func() {
		Context("get account details", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List all projects with invalid param", func() {

	os.Setenv("EMAIL", email)
	os.Setenv("API_TOKEN", apiToken)
	os.Setenv("BASE_URL", baseURL)

	jira := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(jira)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/listProject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListProject)
	handler.ServeHTTP(recorder, request)

	Describe("List Project", func() {
		Context("list all project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List all projects", func() {

	os.Setenv("EMAIL", email)
	os.Setenv("API_TOKEN", apiToken)
	os.Setenv("BASE_URL", baseURL)

	jira := JiraArguments{}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(jira)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/listProject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListProject)
	handler.ServeHTTP(recorder, request)

	Describe("List Project", func() {
		Context("list all project", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get issue by ID without ", func() {

	os.Setenv("EMAIL", email)
	os.Setenv("API_TOKEN", apiToken)
	os.Setenv("BASE_URL", baseURL)

	jira := JiraArguments{IssueID: issueID}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(jira)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/getIssue", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetIssue)
	handler.ServeHTTP(recorder, request)

	Describe("Get Issue", func() {
		Context("get issue by id", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get issue by ID without environment variables", func() {

	jira := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(jira)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/getIssue", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetIssue)
	handler.ServeHTTP(recorder, request)

	Describe("Get Issue", func() {
		Context("get issue by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create issue with invalid param", func() {

	jira := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(jira)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/createIssue", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateIssue)
	handler.ServeHTTP(recorder, request)

	Describe("Create Issue", func() {
		Context("create issue in project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create issue with valid param", func() {

	os.Setenv("EMAIL", email)
	os.Setenv("API_TOKEN", apiToken)
	os.Setenv("BASE_URL", baseURL)

	jira := JiraArguments{AssigneeName: assigneeName, AssigneeEmail: assigneeEmail, Description: "Omg created this Test Task", Type: "Bug", ProjectKey: projectKey, Summary: "Test Task"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(jira)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/createIssue", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateIssue)
	handler.ServeHTTP(recorder, request)

	Describe("Create Issue", func() {
		Context("create issue in project", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
