package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT                    string `mapstructure:"PORT"`
	DbString                string `mapstructure:"DB_STRING"`
	JwtSecret               string `mapstructure:"JWT_SECRET"`
	Environment             string `mapstructure:"ENVIRONMENT"`
	RedisAddr               string `mapstructure:"REDIS_ADDR"`
	RedisPasswd             string `mapstructure:"REDIS_PASSWD"`
	AxiomToken              string `mapstructure:"AXIOM_TOKEN"`
	AxiomOrg                string `mapstructure:"AXIOM_ORG"`
	AxiomDataset            string `mapstructure:"AXIOM_DATASET"`
	GoogleOAuthClientID     string `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleOAuthClientSecret string `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
	GoogleOAuthRedirectURL  string `mapstructure:"GOOGLE_OAUTH_REDIRECT"`
	EncryptionKey           string `mapstructure:"ENC_KEY"`
	GeminiAPIKey            string `mapstructure:"GEMINI_API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	//viper.SetConfigName("app")
	//viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
