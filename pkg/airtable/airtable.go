package airtable

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
)

const (
	// airTableKeyVariable  = "AIRTABLE_KEY"
	airTableKeyVariable = ""
	// airTableAppVariable  = "AIRTABLE_APP"
	airTableAppVariable = ""
	// airTableHostVariable = "AIRTABLE_HOST"
	airTableHostVariable = "https://api.airtable.com/v0/"
	defaultAirTableUrl   = "https://api.airtable.com/v0/%s/names?api_key=%s"
	airTableUrlString    = "https://api.airtable.com/v0/%s/names?api_key=%s"
)

type AirTableClientInterface interface{}

type AirTableClient struct {
	AirTableKey    string
	AirTableUrl    *string
	AirTableClient http.Client
}

func InitializeClient() (*AirTableClient, error) {

	// for glog and anything else
	flag.Parse()

	glog.Info("Starting airtable service")

	airTableUrl, err := generateAirTableURL()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to connect to generate AirTable URL %v", err))
	}

	airTableKey, isSet := os.LookupEnv(airTableKeyVariable)

	if !isSet || airTableKey == "" {
		return nil, errors.New("The AIRTABLE_KEY environment variable is not set")
	}

	return &AirTableClient{
		AirTableKey:    airTableKey,
		AirTableUrl:    airTableUrl,
		AirTableClient: initAirTableClient(),
	}, nil

}

func generateAirTableURL() (*string, error) {
	airTableAppID, isSet := os.LookupEnv(airTableAppVariable)

	if !isSet {
		return nil, errors.New("AirTable App ID is not set")
	}

	airTableHost, isSet := os.LookupEnv(airTableHostVariable)

	if !isSet {
		airTableHost = defaultAirTableUrl
	}

	url := fmt.Sprintf(airTableUrlString, airTableHost, airTableAppID)

	glog.Infof("Initialized AirTable URL: %v", url)

	return &url, nil
}

func initAirTableClient() http.Client {
	return http.Client{
		Timeout: time.Second * 15,
	}
}
