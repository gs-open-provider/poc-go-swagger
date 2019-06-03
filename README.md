# POC-GO-SWAGGER

A POC Project to test how to use Go-Swagger

# Swagger CLI Setup
* copy swagger executable to "usr/local/bin" and "go/bin"

# Intial Folders
* internal
* spec
* output

# Swagger Command for Code Generation
* [CMD] swagger generate server -f spec/root.yml

# Swagger Generated Folders
* cmd
* restapi
* models

# Start Server
* go run cmd/sample-news-feed-server/main.go --port=9090

# Web Page URL
* http://localhost:9090/news-api/v1/test
* Web URL should print ""operation GetTest has not yet been implemented"

# To define implementation and remove initial welcome message
* Go to "restapi/operations" folder and change the file "sample_new_feedP_api.go" with below 2 lines
GetTestHandler:      nil,
PostTestHandler:     nil,