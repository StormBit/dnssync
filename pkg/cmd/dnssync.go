package cmd

import (
	"os"
	"flag"
	"fmt"
	"encoding/json"

	"github.com/stormbit/dnssync/pkg/dnssync"
)

// Execute a thing
func Execute() {
	var dnsProvider string
	flag.StringVar(&dnsProvider, "provider", "cloudflare", "DNS Provider")
	
	var stateFilePath string
	flag.StringVar(&stateFilePath, "state", "", "State File")
	flag.Parse()

	if len(stateFilePath) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a State File")
		os.Exit(1)
	}

	var state dnssync.ZoneState
    stateFile, err := os.Open(stateFilePath)
	defer stateFile.Close()
    if err != nil {
		fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(stateFile)
	jsonParser.Decode(&state)
	
	dnsSync := dnssync.NewCloudflareDNSSync(state.Name, os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	dnsSync.ApplyZoneState(state)
}