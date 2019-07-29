# Go News
A mini news website using Golang, REST API and MongoDB

## Build
### Import libraries
Enable and using go module in current terminal (Go >= 1.11)
```bash
$ export GO111MODULE=on
$ go mod vendor
```
### Run in docker
Build the docker image
```bash
$ make docker
```
Run project in docker
```bash
$ make compose
```
Next, open your browser and access to link [localhost:8080](http://localhost:8080) \
Press `Ctrl+C` to stop

Clean temporary data from docker
```bash
$ make clean
```
## Structure
### Library
- [github.com/gorilla/mux](https://github.com/gorilla/mux)
- [github.com/kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig)