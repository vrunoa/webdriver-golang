# webdriver-golang
A simple example running webdriver tests using golang client

## Requirements
* golang
* or docker

## Setup
In order to run tests, we need to setup our sauce credentials and environment
```
export SAUCE_USERNAME=<your_username>
export SAUCE_ACCESS_KEY=<your_access_key>
export SAUCE_REGION=<sauce_region>
```

Possible regions:
* us-west-1
* us-east-1
* eu-central-1

## Running tests
* Using docker
```
docker-compose -f docker-compose.yml up --build --exit-code-from simple-golang-webdriver-test
```

* Using golang
```
go run selenium_simple.go
```
