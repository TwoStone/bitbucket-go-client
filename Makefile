API_VERSION=v1.2.0

OPENAPI_GENERATOR_VERSION=4.3.1
JAVA=$(shell which java)
BINDIR=hack/bin

fmt:
	go fmt ./...

generate: openapi-generator
	GO_POST_PROCESS_FILE="go fmt" \
	$(JAVA) -jar $(OPENAPI_GENERATOR_JAR) generate \
	--git-user-id TwoStone \
	--git-repo-id bitbucket-go-client \
  	--generator-name go-experimental \
  	--input-spec https://raw.githubusercontent.com/TwoStone/bitbucket-server-api/${API_VERSION}/bitbucket-server-api.yaml \
  	--config hack/generator-config.yaml \
  	--api-package bitbucket \
  	--enable-post-process-file

validate: openapi-generator
	$(JAVA) -jar $(OPENAPI_GENERATOR_JAR) validate \
	--input-spec https://raw.githubusercontent.com/TwoStone/bitbucket-server-api/${API_VERSION}/bitbucket-server-api.yaml \
	--recommend

build: generate fmt
	go mod download
	go build -v ./

openapi-generator:
ifeq (,$(wildcard $(BINDIR)/openapitools/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli.jar))
	mkdir -p $(BINDIR)/openapitools/$(OPENAPI_GENERATOR_VERSION)
	curl https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli-$(OPENAPI_GENERATOR_VERSION).jar \
		-o $(BINDIR)/openapitools/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli.jar
endif
OPENAPI_GENERATOR_JAR=$(BINDIR)/openapitools/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli.jar

