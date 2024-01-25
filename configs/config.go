package configs

import "github.com/spf13/viper"

type conf struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPServerPort     string `mapstructure:"GRPC_SERVER_PORT"`
	GRAPHQLServerPort string `mapstructure:"GRAPHQL_SERVER_Port"`
	RabbitMqUser      string `mapstructure:"RABBIT_MQ_USER"`
	RabbitMqPassword  string `mapstructure:"RABBIT_MQ_PASSWORD"`
	RabbitMqHost      string `mapstructure:"RABBIT_MQ_HOST"`
	RabbitMqPort      string `mapstructure:"RABBIT_MQ_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	var conf *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&conf)

	if err != nil {
		panic(err)
	}

	return conf, nil

}
