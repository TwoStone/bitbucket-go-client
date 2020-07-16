generate:
	openapi-generator generate \
	--git-user-id TwoStone \
	--git-repo-id bitbucket-go-client \
  	--generator-name go-experimental \
  	--input-spec https://raw.githubusercontent.com/TwoStone/bitbucket-server-api/master/bitbucket-server-api.yaml \
  	--config hack/generator-config.yaml \
  	--api-package bitbucket

build: 
	go mod download
	go build -v ./

install:
	go install -i ./
