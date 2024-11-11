package config

import "time"

const (
	MySQLHost     = "mysql"
	MySQLPort     = "3307"
	MySQLDatabase = "users-api"
	MySQLUsername = "root"
	MySQLPassword = "RaTa8855"
	CacheDuration = 30 * time.Second
	MemcachedHost = "memcached"
	MemcachedPort = "11211"
	JWTKey        = "ThisIsAnExampleJWTKey!"
	JWTDuration   = 24 * time.Hour
)
