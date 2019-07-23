FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/heaptracetechnology/jira-go-sdk

#RUN go get github.com/andygrunwald/go-jira

WORKDIR /go/src/github.com/heaptracetechnology/jira

ADD . /go/src/github.com/heaptracetechnology/jira

RUN go install github.com/heaptracetechnology/jira

ENTRYPOINT jira

EXPOSE 3000