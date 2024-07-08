package config

type MysqlConfig struct {
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	DbName   string `mapstructure:"dbname" json:"dbName"`
}
