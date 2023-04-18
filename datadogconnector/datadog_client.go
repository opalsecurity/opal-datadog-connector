package datadogconnector

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

var client *datadog.APIClient

func GetDatadogClient() *datadog.APIClient {
	if client == nil {
		client = initDatadogClient()
	}
	return client
}

func initDatadogClient() *datadog.APIClient {
	configuration := datadog.NewConfiguration()
	return datadog.NewAPIClient(configuration)
}
