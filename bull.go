package main

var prompt string = "bull > "

type session struct {
	SessionId string
	IPAddress string
	Username string
	Hostname string
	OperatingSystem string
	Job string
	Alias string
	Seen int64
}

var selected_session string
var sessions []session = []session{}

var Server_Info map[string]string = map[string]string {
	"version": "1.07-B",
	"engine_server": "0.0.0.0:8080",
}