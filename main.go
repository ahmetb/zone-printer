// Copyright 2020 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

var (
	computeZone string
)

func main() {
	if !metadata.OnGCE() {
		log.Println("warn: not running with metadata service present")
	} else {
		zone, err := metadata.Zone()
		if err != nil {
			log.Fatalf("failed to get compute zone: %+v", err)
		}
		computeZone = zone
		log.Printf("info: determined zone: %q", zone)
	}

	port := "8080"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
	log.Println("starting to listen on port " + port)
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}

func handle(w http.ResponseWriter, r *http.Request) {
	var srcIP string
	if ipHeader := r.Header.Get("X-Forwarded-For"); ipHeader != "" {
		srcIP = ipHeader
	} else {
		srcIP = r.RemoteAddr
	}
	log.Printf("received request method=%s path=%q src=%q", r.Method, r.URL.Path, srcIP)

	if computeZone == "" {
		fmt.Fprintf(w, `<!DOCTYPE html>
				<h1>Cannot determine the compute zone :(</h1>
				<p>Is it running on a Google Compute Engine or Cloud Run?</p>`)
		return
	}

	region := computeZone[:strings.LastIndex(computeZone, "-")]
	dc, ok := datacenters[region]

	fmt.Fprintf(w, `<!DOCTYPE html>
	<h4>Welcome from Google Cloud datacenters at:<h4>`)
	if !ok {
		// cannot determine datacenter from zone, just use zone name
		fmt.Fprintf(w, `<h1>%s!</h1>`, computeZone)

	} else {
		fmt.Fprintf(w, `<h1>%s</h1>
		<h3>You are now connected to &quot;%s&quot;</h3>
		<img src="%s" style="width: 480px; height: auto; border: 1px solid #444;"/>`, dc.location, computeZone, dc.flagURL)
	}
	fmt.Fprintf(w, `
		<p>
			Based on where you visit from, Google Cloud Load Balancer routes your request
			to the closest compute region the application is deployed in.
		</p>
		<p>
			<small><a href="https://github.com/ahmetb/zone-printer">[source code]</a></small>
		</p>`)
}

var (
	// Datacenter list is adopted from https://cloud.google.com/compute/docs/regions-zones/
	// also Cloud Run regions are at https://cloud.google.com/run/docs/locations
	datacenters = map[string]struct {
		location string
		flagURL  string // flag images must be public domain
	}{
		"northamerica-northeast1": {
			location: "Montréal, Canada",
			flagURL:  "https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Canada_%28Pantone%29.svg",
		},
		"us-central1": {
			location: "Council Bluffs, Iowa, USA",
			flagURL:  "https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg",
		},
		"us-west1": {
			location: "The Dalles, Oregon, USA",
			flagURL:  "https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg",
		},
		"us-east4": {
			location: "Ashburn, Virginia, USA",
			flagURL:  "https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg",
		},
		"us-east1": {
			location: "Moncks Corner, South Carolina, USA",
			flagURL:  "https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg",
		},
		"southamerica-east1": {
			location: "São Paulo, Brazil",
			flagURL:  "https://upload.wikimedia.org/wikipedia/en/0/05/Flag_of_Brazil.svg",
		},
		"europe-north1": {
			location: "Hamina, Finland",
			flagURL:  "https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Finland.svg",
		},
		"europe-west1": {
			location: "St. Ghislain, Belgium",
			flagURL:  "https://upload.wikimedia.org/wikipedia/commons/6/65/Flag_of_Belgium.svg",
		},
		"europe-west2": {
			location: "London, U.K.",
			flagURL:  "https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg",
		},
		"europe-west3": {
			location: "Frankfurt, Germany",
			flagURL:  "https://upload.wikimedia.org/wikipedia/en/b/ba/Flag_of_Germany.svg",
		},
		"europe-west4": {
			location: "Eemshaven, Netherlands",
			flagURL:  "https://upload.wikimedia.org/wikipedia/commons/2/20/Flag_of_the_Netherlands.svg",
		},
		"europe-west6": {
			location: "Zurich, Switzerland",
			flagURL:  "https://upload.wikimedia.org/wikipedia/commons/f/f3/Flag_of_Switzerland.svg",
		},
		"asia-south1": {
			location: "Mumbai, India",
			flagURL:  "https://upload.wikimedia.org/wikipedia/en/4/41/Flag_of_India.svg",
		},
		"asia-southeast1": {
			location: "Jurong West, Singapore",
			flagURL:  "https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Singapore.svg",
		},
		"asia-southeast2": {
			location: "Jakarta, Indonesia",
			flagURL:  "https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_Indonesia.svg",
		},
		"asia-east1": {
			location: "Changhua County, Taiwan",
			flagURL:  "https://upload.wikimedia.org/wikipedia/commons/7/72/Flag_of_the_Republic_of_China.svg",
		},
		"asia-east2": {
			location: "Hong Kong",
			flagURL: "https://upload.wikimedia.org/wikipedia/commons/5/5b/Flag_of_Hong_Kong.svg",
		},
		"asia-northeast1": {
			location: "Tokyo, Japan",
			flagURL:  "https://upload.wikimedia.org/wikipedia/en/9/9e/Flag_of_Japan.svg",
		},
		"asia-northeast2": {
			location: "Osaka, Japan",
			flagURL:  "https://upload.wikimedia.org/wikipedia/en/9/9e/Flag_of_Japan.svg",
		},
		"asia-northeast3": {
			location: "Seoul, South Korea",
			flagURL: "https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_South_Korea.svg",
		},
		"australia-southeast1": {
			location: "Sydney, Australia",
			flagURL:  "https://upload.wikimedia.org/wikipedia/en/b/b9/Flag_of_Australia.svg",
		},
	}
)
