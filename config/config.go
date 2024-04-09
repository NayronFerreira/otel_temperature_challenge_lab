package config

import "github.com/spf13/viper"

type Config struct {
	WebValidationCEPServerPort string `mapstructure:"WEB_VALIDATION_CEP_SERVER_PORT"`
	WebTemperatureServerPort   string `mapstructure:"WEB_TEMPERATURE_SERVER_PORT"`
	ViaCepHost                 string `mapstructure:"VIACEP_HOST_API"`
	WeatherHost                string `mapstructure:"WEATHER_HOST_API"`
	WeatherKey                 string `mapstructure:"WEATHER_API_KEY"`
	TimeoutSeconds             string `mapstructure:"TIMEOUT_SECONDS"`
	ServiceName                string `mapstructure:"SERVICE_NAME"`
	CollectorURL               string `mapstructure:"COLLECTOR_URL"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile("secrets.env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config, err
}
