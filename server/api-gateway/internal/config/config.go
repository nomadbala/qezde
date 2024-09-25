package config

import "time"

// App
const (
	MaxHeaderBytes = 1 << 20
	ReadTimeOut    = 10 * time.Second
	WriteTimeOut   = 10 * time.Second
)

func New() {

}
