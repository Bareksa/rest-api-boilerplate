package models


type Config struct {
	Sentry struct{
		Dsn string `json:"dsn"`
		Timeout int `json:"timeout"`
	} `json:"sentry"` 
}