package config

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
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
