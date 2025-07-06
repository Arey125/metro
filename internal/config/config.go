package config

import (
	"fmt"
	"os"
	"strconv"

	"metro/internal/metro"
	"metro/internal/users"

	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	Port     int
	Db       string
	Oauth    users.OauthConfig
	Secure   bool
	CertFile string
	KeyFile  string
	Metro    metro.MetroServiceConfig
}

func getString(name string) string {
	res := os.Getenv(name)
	if res == "" {
		panic(fmt.Sprintf("%s environment variable is not set", name))
	}
	return res
}

func getInt(name string) int {
	stringRes := getString(name)
	res, err := strconv.Atoi(stringRes)
	if err != nil {
		panic(fmt.Sprintf("cannot parse %s environment variable as an integer", name))
	}
	return res
}

func getBool(name string) bool {
	res := os.Getenv(name)
	if res == "true" {
		return true
	}
	if res == "false" {
		return false
	}
	panic(fmt.Sprintf("%s environment variable is not set", name))
}

func Get() config {
	port := getInt("PORT")
	db := getString("DB")

	id := getString("CLIENT_ID")
	secret := getString("CLIENT_SECRET")
	callbackUrl := getString("CLIENT_CALLBACK_URL")
	sessionSecret := getString("SESSION_SECRET")

	secure := getBool("SECURE")
	certFile := ""
	keyFile := ""
	if secure {
		certFile = getString("CERT_FILE")
		keyFile = getString("KEY_FILE")
	}

	dataCollectionIntervalMs := getInt("DATA_COLLECTION_INTERVAL_MS")

	return config{
		Port: port,
		Db:   db,
		Oauth: users.OauthConfig{
			Id:            id,
			Secret:        secret,
			CallbackUrl:   callbackUrl,
			SessionSecret: sessionSecret,
		},
		Secure:   secure,
		CertFile: certFile,
		KeyFile:  keyFile,

		Metro: metro.MetroServiceConfig{
			DataCollectionIntervalMs: dataCollectionIntervalMs,
		},
	}
}
