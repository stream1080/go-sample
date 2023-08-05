package config

type Config struct {
	*ServerConfig `yaml:"server"`
	*MySQLConfig  `yaml:"mysql"`
	*RedisConfig  `yaml:"redis"`
	*LogConfig    `yaml:"log"`
}

type ServerConfig struct {
	Mode string `yaml:"mode"`
	Port int    `yaml:"port"`
}

type MySQLConfig struct {
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DB           string `yaml:"dbname"`
	Port         int    `yaml:"port"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
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

type LogConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}
