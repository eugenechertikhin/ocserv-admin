package model

import "time"

type Base struct {
	User     string
	Group    string
	Password string
}

type Online struct {
	Id     int
	User   string
	ExtIp  string
	IntIp  string
	Device string
	Since  time.Duration
}
