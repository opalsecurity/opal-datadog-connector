SHELL := /bin/bash

MAIN_GO_SRC=main.go

GOAUTO=gin -i -p 8091 --appPort 8090 run $(MAIN_GO_SRC)
SECRETS_DOTFILE=.secrets

run:
	source $(SECRETS_DOTFILE) && $(GOAUTO)
	
gen-openapi:
	openapi-generator generate -g go-gin-server -i <(curl -L https://raw.githubusercontent.com/opalsecurity/opal-connector-spec/main/openapi-connector-spec.yaml) -c openapi_config.json