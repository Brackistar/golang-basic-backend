package aws

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lookUpEnvironmentVariable(t *testing.T) {
	var testCases = []struct {
		envName       string
		envVal        string
		searchFor     string
		expectedError error
	}{
		{"NOT_TEST_ENV_VAR", "1234", "TEST_ENV_VAR", errors.New("no environment variable \"TEST_ENV_VAR\" found")}, // No environment variable with the same name as serarched
		{"TEST_ENV_VAR", "1234", "TEST_ENV_VAR", nil},                                                              // Environment variable with same name as searched for exists
	}

	for _, tc := range testCases {
		cleanEnv := makeRestEnv(os.Environ())

		t.Setenv(tc.envName, tc.envVal)
		assert := assert.New(t)

		err := lookUpEnvironmentVariable(tc.searchFor)

		if tc.expectedError != nil {
			assert.EqualError(err, tc.expectedError.Error())
		} else {
			assert.Nil(err)
		}

		cleanEnv()
	}
}

func Test_validateEnvironment(t *testing.T) {
	var testCases = []struct {
		envVars       map[string]string
		expectedError error
	}{
		{ // Required environment variables exist
			map[string]string{
				"URL_PREFIX":  "1234",
				"BUCKET_NAME": "1234",
				"SECRET_NAME": "1234",
			},
			nil,
		},
		{ // Required environment variables exist, and also other environment variables
			map[string]string{
				"URL_PREFIX":   "1234",
				"TEST_ENVVAR2": "1234",
				"BUCKET_NAME":  "1234",
				"TEST_ENVVAR3": "1234",
				"SECRET_NAME":  "1234",
				"TEST_ENVVAR1": "1234",
			},
			nil,
		},
		{ // No environmental variable with name URL_PREFIX exists
			map[string]string{
				"BUCKET_NAME": "1234",
				"SECRET_NAME": "1234",
			},
			errors.New("no environment variable \"URL_PREFIX\" found"),
		},
		{ // No environmental variable with name BUCKET_NAME exists
			map[string]string{
				"URL_PREFIX":  "1234",
				"SECRET_NAME": "1234",
			},
			errors.New("no environment variable \"BUCKET_NAME\" found"),
		},
		{ // No environmental variable with name SECRET_NAME exists
			map[string]string{
				"URL_PREFIX":  "1234",
				"BUCKET_NAME": "1234",
			},
			errors.New("no environment variable \"SECRET_NAME\" found"),
		},
		{ // Multiple environmental variables doesn't exist
			map[string]string{
				"NOT_URL_PREFIX": "1234",
				"BUCKET_NAME":    "1234",
			},
			errors.New("no environment variable \"URL_PREFIX\" found"),
		},
	}

	for _, tc := range testCases {
		assert := assert.New(t)
		cleanEnv := makeRestEnv(os.Environ())
		registerEnvVar(tc.envVars, t)

		err := validateEnvironment()

		if tc.expectedError == nil {
			assert.Nil(err)
		} else {
			assert.EqualError(err, tc.expectedError.Error())
		}

		cleanEnv()
	}
}

func registerEnvVar(variables map[string]string, t *testing.T) {
	for key, val := range variables {
		t.Setenv(key, val)
	}
}

func makeRestEnv(vals []string) func() {
	return func() {
		os.Clearenv()
		for _, v := range vals {
			parts := strings.SplitN(v, "=", 2)
			os.Setenv(parts[0], parts[1])
		}
	}
}
