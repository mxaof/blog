package config

type Server struct {
	Mysql MysqlConfig `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis RedisConfig `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap   ZapConfig   `mapstructure:"zap" json:"zap" yaml:"zap"`
}
