package main

var prompt string = "bull > "

type session struct {
	SessionId string
	IPAddress string
	Username string
	Hostname string
	OperatingSystem string
	Job string
	Alias string ///////
	Seen int64
}

var selected_session string
var sessions []session = []session{}