package config

import "time"

const (
	ListenAddr     = "0.0.0.0:9090"
	PodLogTailLine = 2000
	DBUser         = "root"
	DBPassword     = "1qaz@WSX"
	DBHost         = "yunxue521.top"
	DBPort         = "32306"
	DBName         = "kubemanage"
	MaxIdleTime    = 20
	MaxOpenConns   = 100
	MaxLifetime    = 20 * time.Second
)
