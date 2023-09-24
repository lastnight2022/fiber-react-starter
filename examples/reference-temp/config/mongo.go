package config

type Mongo struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	DbName   string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`
	UserName string `mapstructure:"user_name" json:"user_name" yaml:"user_name"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
