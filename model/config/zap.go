package config

type ZapConfig struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"` // 级别
	FileName   string `mapstructure:"fileName" json:"fileName" yaml:"fileName"`
	MaxSize    int    `mapstructure:"maxSize" json:"maxSize" yaml:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge" json:"maxAge" yaml:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups" json:"maxBackups" yaml:"maxBackups"`
}
