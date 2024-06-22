package helpers

import "github.com/spf13/viper"

type EnvConfig struct {
	SqlServer   string
	SqlUser     string
	SqlPassword string
	SqlPort     int
	SqlDatabase string
}

func LoadConfig(path string) (cfg EnvConfig, err error) {
	if GetEnv("AppMode", "debug") == "release" {
		cfg = EnvConfig{
			SqlServer:   GetEnv("SQL_SERVER", ""),
			SqlUser:     GetEnv("SQL_USER", ""),
			SqlPassword: GetEnv("SQL_PASSWORD", ""),
			SqlPort:     GetEnvAsInt("SQL_PORT", 1306),
			SqlDatabase: GetEnv("SQL_DATABASE", ""),
		}
		return cfg, nil
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return EnvConfig{}, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return EnvConfig{}, err
	}
	return cfg, nil
}
