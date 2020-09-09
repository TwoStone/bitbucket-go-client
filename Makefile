generate:
	GO_POST_PROCESS_FILE="go fmt" \
	openapi-generator generate \
	--git-user-id TwoStone \
	--git-repo-id bitbucket-go-client \
  	--generator-name go-experimental \
  	--input-spec https://raw.githubusercontent.com/TwoStone/bitbucket-server-api/v1.0.3/bitbucket-server-api.yaml \
  	--config hack/generator-config.yaml \
  	--api-package bitbucket \
  	--enable-post-process-file

build: 
	go mod download
	go build -v ./

install:
	go install -i ./
