package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	DBDriver      string           `mapstructure:"DB_DRIVER"`
	DBHost        string           `mapstructure:"DB_HOST"`
	DBPort        string           `mapstructure:"DB_PORT"`
	DBUser        string           `mapstructure:"DB_USER"`
	DBPassword    string           `mapstructure:"DB_PASSWORD"`
	DBName        string           `mapstructure:"DB_NAME"`
	WebServerPort string           `mapstructure:"WEB_SERVER_PORT"`
	JwtSecret     string           `mapstructure:"JWT_SECRET"`
	JwtExpiration int              `mapstructure:"JWT_EXPIRATION"`
	TokenAuth     *jwtauth.JWTAuth //TODO: Implementar inversão de dependência
}

func LoadConfig(path string) (*config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	//? Sobrescreve as variaveis do .env pelas variáveis de ambiente da máquina
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JwtSecret), nil)

	return cfg, nil
}
