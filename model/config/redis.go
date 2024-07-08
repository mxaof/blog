package config

type RedisConfig struct {
	Address  string `mapstructure:"address" json:"address" yaml:"address"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Db       int    `mapstructure:"db" json:"db" yaml:"db"`
}
