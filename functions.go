package main

import (
	"fmt"
	"time"
)

func display_sessions(s []session) {
	var timenow int64 = time.Now().UTC().Unix()

        var longest_username int
        for i:=0;i<len(sessions);i++ {
            if len(sessions[i].Username+"@"+sessions[i].Hostname) > longest_username {
                    longest_username = len(sessions[i].Username+"@"+sessions[i].Hostname)
            }
        }

        if longest_username < 4 {
            longest_username = 4
        }

        fmt.Println("SessionID            IPAddress       Os      User",gchar(" ", longest_username - 4)+"Status")
        fmt.Println("-------------------- --------------- -------",gchar("-", longest_username)+" --------")
        for i:=0;i<len(sessions);i++ {
                var status string = Fore["RED"]+"Inactive"+Fore["RESET"]
                if sessions[i].Seen > timenow - 5 {
                        status = Fore["GREEN"]+"Active"+Fore["RESET"]
                }
                fmt.Println(Fore["YELLOW"]+sessions[i].SessionId+Fore["RESET"],Fore["BLUE"]+sessions[i].IPAddress+Fore["RESET"]+gchar(" ", 15 - len(sessions[i].IPAddress)),sessions[i].OperatingSystem+gchar(" ", 7 - len(sessions[i].OperatingSystem)),Fore["RED"]+sessions[i].Username+"@"+sessions[i].Hostname+Fore["RESET"]+gchar(" ", longest_username - len(sessions[i].Username+"@"+sessions[i].Hostname)),status)
        }

        if len(sessions) == 0 {
                fmt.Println(Fore["YELLOW"]+"None"+Fore["BLUE"]+"                 None            "+Fore["RESET"]+"None"+Fore["RED"]+"    None",Fore["RESET"]+gchar(" ", longest_username - 4)+"None")
        }
}

func gchar(c string, l int) string {
	var result string

	for i:=0;i<l;i++ {
		result+=c
	}

        return result
}   