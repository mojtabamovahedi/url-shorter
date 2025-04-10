package config

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
	Redis  RedisConfig  `yaml:"redis"`
}

type ServerConfig struct {
	HttpPort uint `yaml:"http_port"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	Schema   string `yaml:"schema"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port uint   `yaml:"port"`
	DB   int    `yaml:"db"`
	TTL  int    `yaml:"ttl"`
}
