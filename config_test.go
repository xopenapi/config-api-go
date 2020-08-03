package config_test

import(
	"testing"
	"github.com/xopenapi/config-api-go"
)

func getClient() *config.APIClient {
	cfg := config.NewConfiguration()
	client := config.NewAPIClient(cfg)

	return client
}

func TestCreateConfig(t *testing.T) {
	client := getClient()
	t.Log(client)
}