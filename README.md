# _Jira OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/godaddy.svg?branch=master)](https://travis-ci.com/heaptracetechnology/godaddy)
[![codecov](https://codecov.io/gh/heaptracetechnology/godaddy/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/godaddy)


An OMG service for godaddy, check availability for specific domain name.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Check domain availability
```coffee
godaddy checkDomainAvailability domain:'domain.com'
{"available": true,"currency": "USD","definitive": true,"domain": "domain.com"}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Check domain availability
```shell
$ omg run checkDomainAvailability -a domain=<DOMAIN_NAME> -e API_KEY=<API_KEY> -e API_SECRET=<API_SECRET>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/godaddy/blob/master/LICENSE).
