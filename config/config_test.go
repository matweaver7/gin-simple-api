package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"matweaver.com/simple-rest-api/config"
)

// helper function to set up environment variables for testing
func setEnvVars(envMap map[string]string) {
	for key, value := range envMap {
		os.Setenv(key, value)
	}
}

// helper function to clean up environment variables after testing
func unsetEnvVars(keys []string) {
	for _, key := range keys {
		os.Unsetenv(key)
	}
}

func TestNewConfigLocalEnv(t *testing.T) {
	// Assuming there's a .env file in the project root for testing
	// You might need to create this file or mock `godotenv.Load` function behavior
	e := map[string]string{
		"ENV":                    "LOCAL",
		"PORT":                   "777",
		"MONGO_DB_NAME":          "root",
		"MONGO_COLLECTION":       "cow",
		"MONGO_TIMEOUT_DURATION": "10",
		"MONGO_URI":              "fail",
	}
	var keys []string
	for key := range e {
		keys = append(keys, key)
	}

	setEnvVars(e)
	defer unsetEnvVars(keys)

	testCases := []struct {
		path       string
		shouldPass bool
	}{
		{
			path:       "../.env",
			shouldPass: true,
		},
		{
			path:       "",
			shouldPass: false,
		},
	}

	for _, v := range testCases {
		config, err := config.NewConfig(v.path)

		if v.shouldPass == true {
			assert.NoError(t, err)
			assert.NotNil(t, config)

			assert.Equal(t, e["PORT"], config.Port, "Should Pass: Setting Config from Environment Variables")
			assert.Equal(t, e["ENV"], config.Env, "Should Pass: Setting Config from Environment Variables")
			assert.Equal(t, e["MONGO_DB_NAME"], config.Mongo.DatabaseName, "Should Pass: Setting Config from Environment Variables")
			assert.Equal(t, e["MONGO_URI"], config.Mongo.URI, "Should Pass: Setting Config from Environment Variables")
		} else {
			assert.Error(t, err)
		}
	}
}
