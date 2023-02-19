package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)


var engine_server string = "0.0.0.0:8080"
var version string = "1.0.4-BETA"

func main() {
	animate("Starting the BUll-Framework")
	init_help()

	rand.Seed(time.Now().UnixNano())

	if len(os.Args) > 1 {
		engine_server = os.Args[1]
	}

	fmt.Printf("%s", logo+"\n")

	Log.Info("BUll is on version "+Fore.YELLOW+version+Fore.RESET+"\n")
	Log.Info("Starting BULL-ENGINE on "+Fore.YELLOW+engine_server+Fore.RESET+"\n\n")

	Log.Warning("Bull is not an fully interactive reverse-shell handler.\n")

	http.HandleFunc("/", Handle)

    go func() {
        if err := http.ListenAndServe(":8080", nil); err != nil {
            Log.Error(err.Error())
        }
    }()

	makeshell()
}