package route

import (
    "github.com/gorilla/mux"
    jira "github.com/heaptracetechnology/jira/jira"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "GetAccountDetails",
        "POST",
        "/getAccount",
        jira.GetAccountDetails,
    },
    Route{
        "CreateIssue",
        "POST",
        "/createIssue",
        jira.CreateIssue,
    },
    Route{
        "CreateProject",
        "POST",
        "/createProject",
        jira.CreateProject,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
