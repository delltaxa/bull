package main

import (
	"regexp"
)

func verify_session(sess session) bool {
	if verify_sessionid(sess.SessionId) && verify_sessionos(sess.OperatingSystem) {
		return true
	}
	
	return false
}

func verify_sessionos(sessionOS string) bool {
	var supported []string = []string {"windows", "linux"}
	
	for i:=0;i<len(supported);i++ {
		if supported[i] == sessionOS {
			return true
		}
	}

	return false
}

func verify_sessionid(sessionID string) bool {
    pattern := "^[a-f0-9]{6}-[a-f0-9]{6}-[a-f0-9]{6}$"

    regex := regexp.MustCompile(pattern)

    return regex.MatchString(sessionID)
}  