# key-value persistent store

[![Go Report Card](https://goreportcard.com/badge/github.com/eyalgolan/key-value-persistent-store)](https://goreportcard.com/report/github.com/eyalgolan/key-value-persistent-store)

A service that exposes a REST API that allows users to create, update, delete, and retrieve answers as key-value pairs.
The answers are stored in a Postgres DB, so they can handle service restarts.

An answer is defined as:

```json
{
  "key" : "name",
  "value" : "Eyal"
}
```

An event is defined as:
```json
{
  "event" : "create",
  "data" : 
  {
    "key": "name",
    "value": "Eyal"
  }
}
```
### API
The API exposes the following endpoints:

* Get answer for a given key (returns the latest answer for the given key): ```GET /answers/:key```
* Create answer: ```POST /answers```
* Update answer: ```PUT /answers```
* Delete answer: ```DELETE /answers/:key```
* Get history for a given key (returns an array of events in chronological order):
```GET /answers/:key/history```


### HOW TO RUN
#### Pre-requisites
1. Have docker and docker-compose installed and a docker engine running.
2. Have ports 5432 and 8080 available.

The first build and run the application, run:
```bash
make install
```
For later runs, use:
```bash
make start
```
To stop the application run:
```bash
make stop
```
