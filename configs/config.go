package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	Log    Log
	Server Server
}

type Log struct {
	Environment string
	Application string
}

type Server struct {
	Port   string
	PortGw string
}

func GetConfig() config {
	return *cfg
}

func init() {
	viper.SetDefault("PORT", "50051")
	viper.SetDefault("PORTGW", "8080")

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	viper.ReadInConfig()

	cfg = &config{
		Log: Log{
			Environment: viper.GetString("ENVIRONMENT"),
			Application: viper.GetString("APPLICATION"),
		},
		Server: Server{
			Port:   viper.GetString("PORT"),
			PortGw: viper.GetString("PORTGW"),
		},
	}
}
