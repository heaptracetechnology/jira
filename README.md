# _Jira_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)

An OMG service for jira, to create issue, list project etc. jira workflows control the status of the project.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Get Account
```coffee
jira getAccount username:'username'
```
##### Create Issue
```coffee
jira createIssue assigneeName:'assignee username' assigneeEmail:'assignee email' description:'description' type:'issue type' projectKey:'project key' summary:'project summary'
```
##### List Project
```coffee
jira listProject
```
##### Get Issue By ID
```coffee
jira getIssue issueId:'issue Id'
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Get Account
```shell
$ omg run getAccount -a username=<USERNAME> -e EMAIL=<ACCOUNT_EMAIL> -e API_TOKEN=<API_TOKEN> -e BASE_URL=<BASE_URL>
```
##### Create Issue
```shell
$ omg run createIssue -a assigneeName=<ASSIGNEE_USERNAME> -a assigneeEmail=<ASSIGNEE_EMAIL> -a description=<DESCRIPTION> -a type=<ISSUE_TYPE> -a projectKey=<PROJECT_KEY> -a summary=<PROJECT_SUMMARY> -e EMAIL=<ACCOUNT_EMAIL> -e API_TOKEN=<API_TOKEN> -e BASE_URL=<BASE_URL>
```
##### List Project
```shell
$ omg run listProject -e EMAIL=<ACCOUNT_EMAIL> -e API_TOKEN=<API_TOKEN> -e BASE_URL=<BASE_URL>
```
##### Get Issue By ID
```shell
$ omg run getIssue -a issueId=<ISSUE_ID> -e EMAIL=<ACCOUNT_EMAIL> -e API_TOKEN=<API_TOKEN> -e BASE_URL=<BASE_URL>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/jira/blob/master/LICENSE).
