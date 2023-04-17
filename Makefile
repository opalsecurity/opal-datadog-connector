SHELL := /bin/bash

MAIN_GO_SRC=main.go

GOAUTO=gin -i -p 8091 --appPort 8090 run main.go
ENV_SETTER_SCRIPT=.secrets

run:
	source $(ENV_SETTER_SCRIPT) dev-local && $(GOAUTO)
	
# TODO: move .yaml to a separate repo
gen-openapi:
	openapi-generator generate -g go-gin-server -i ../opal/web/backend/lib/sync/connections/customconnector/custom_connector.yaml -c openapi_config.json