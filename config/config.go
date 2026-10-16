package config

import "log"

type AppConfig struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}
