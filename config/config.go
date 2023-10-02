package config

type Config struct {
	ServerConfig
	MySQLConfig
	RedisConfig
}

type ServerConfig struct {
	Mode string `env:"SERVER_MODE,default=debug"`
	Port int    `env:"SERVER_PORT,default=8080"`
}

type MySQLConfig struct {
	Host     string `env:"MYSQL_HOST,default=127.0.0.1"`
	User     string `env:"MYSQL_USER,default=sample"`
	Password string `env:"MYSQL_PASSWORD,default=123456"`
	DB       string `env:"MYSQL_DB,default=sample"`
	Port     int    `env:"MYSQL_PORT,default=3306"`
}

type RedisConfig struct {
	Host         string `env:"REDIS_HOST,default=127.0.0.1"`
	User         string `env:"REDIS_USER,default=sample"`
	Password     string `env:"REDIS_PASSWORD,default=123456"`
	Port         int    `env:"REDIS_PORT,default=6379"`
	DB           int    `env:"REDIS_DB,default=0"`
	PoolSize     int    `env:"REDIS_POOL_SIZE,default=100"`
	MinIdleConns int    `env:"REDIS_MIN_IDLE_CONNS,default=20"`
}
