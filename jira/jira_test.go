package jira

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
	apiKey     = os.Getenv("GODADDY_API_KEY")
	apiSecret  = os.Getenv("GODADDY_API_SECRET")
	domainName = os.Getenv("GODADDY_DOMAIN_NAME")
)

var _ = Describe("Check Domain Availability with invalid param", func() {

	godaddy := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(godaddy)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/checkDomain", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CheckDomainAvailability)
	handler.ServeHTTP(recorder, request)

	Describe("Check Availability", func() {
		Context("availability", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Check Domain Availability without domain name", func() {

	os.Setenv("API_KEY", apiKey)
	os.Setenv("API_SECRET", apiSecret)

	godaddy := Arguments{}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(godaddy)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/checkDomain", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CheckDomainAvailability)
	handler.ServeHTTP(recorder, request)

	Describe("Check Availability", func() {
		Context("availability", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Check Domain Availability with all required param ", func() {

	os.Setenv("API_KEY", apiKey)
	os.Setenv("API_SECRET", apiSecret)

	godaddy := Arguments{Domain: domainName}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(godaddy)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/checkDomain", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CheckDomainAvailability)
	handler.ServeHTTP(recorder, request)

	Describe("Check Availability", func() {
		Context("availability", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
