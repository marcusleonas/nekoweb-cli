package nekoweb

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func GetAPIKey() string {
	apiKey := viper.GetString("apikey")

	if apiKey == "" {
		fmt.Println("apikey is undefined. have you run `nekoweb auth <api_key>`?")
		os.Exit(1)
	}

	return apiKey
}
