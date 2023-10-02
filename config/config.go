package config

type Config struct {
	ServerConfig `yaml:"server"`
	MySQLConfig  `yaml:"mysql"`
	RedisConfig  `yaml:"redis"`
}

type ServerConfig struct {
	Mode string `yaml:"mode"`
	Port int    `yaml:"port"`
}

type MySQLConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"dbname"`
	Port     int    `yaml:"port"`
}

type RedisConfig struct {
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Port         int    `yaml:"port"`
	DB           int    `yaml:"db"`
	PoolSize     int    `yaml:"pool_size"`
	MinIdleConns int    `yaml:"min_idle_conns"`
}
