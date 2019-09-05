package restapi

import (
	"fmt"
	"os"
	"strconv"

	"api.tantakatu.com/logging"
	"api.tantakatu.com/restapi/operations"
	"api.tantakatu.com/routes"
	"gopkg.in/myles-mcdonnell/loglight.v3"
)

type config struct {
	logOutputFormat     string
	logDebug            bool
	useMockEmailService bool
	apiKey              string
}

// InitApi initialises the server
func InitApi(api *operations.TantakatuAPI) {

	conf := config{
		logOutputFormat:     getOptionalEnvStr("LOG_OUTPUT_FORMAT", "DEBUG"),
		logDebug:            getEnvBool("LOG_OUTPUT_DEBUG"),
		useMockEmailService: getEnvBool("MOCK_SERVICE"),
		apiKey:              getMandatoryEnvStr("API_KEY"),
	}

	logging.Initialise(loglight.NewLogger(
		conf.logDebug,
		logging.BuildFormatter(conf.logOutputFormat),
	))

	api.Logger = func(msg string, args ...interface{}) {
		logging.LogInfo(logging.Api, nil, args)
	}

	if conf.useMockEmailService {

		logging.LogInfo(logging.Api, nil, "Using Mock Service")
		routes.BindRoutes(api)

	} else {

		//Implement production code

	}
}

func getEnvBool(key string) bool {
	str := os.Getenv(key)
	bool, _ := strconv.ParseBool(str)
	return bool
}

func getMandatoryEnvStr(key string) string {
	str := os.Getenv(key)
	if str == "" {
		fmt.Println("need envvar: " + key)
	}
	return str
}

func getOptionalEnvStr(key string, def string) string {
	str := os.Getenv(key)
	if str == "" {
		return def
	}
	return str
}