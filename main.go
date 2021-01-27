// Copyright 2021 Google Inc.
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
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

var (
	computeZone string
	templates   *template.Template
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
	if v := os.Getenv("FAKE_ZONE"); v != "" {
		// TODO(ahmetb) remove before submitting
		computeZone = v
	}

	var err error
	templates, err = template.ParseGlob(filepath.Join("templates", "*.html.tpl"))
	if err != nil {
		log.Fatalf("template parsing error: %v", err)
	}

	port := "8080"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
	addr := ""
	if v := os.Getenv("ADDR"); v != "" {
		addr = v
	}
	log.Println("starting to listen on port " + port)
	http.HandleFunc("/", handle)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err = http.ListenAndServe(net.JoinHostPort(addr, port), nil)
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
		if err := templates.ExecuteTemplate(w, "errorPage", map[string]interface{}{
			"title":         "Error!!1",
			"error_title":   "Cannot determine the compute zone",
			"error_message": "Is it running on a Google Compute Engine or Cloud Run?",
		}); err != nil {
			log.Fatalf("failed to render template: %v", err)
		}
		return
	}

	region := computeZone[:strings.LastIndex(computeZone, "-")]
	var cityName, flagURL string
	dc, ok := datacenters[region]
	if ok {
		cityName = dc.location
		flagURL = dc.flagURL
	}

	if err := templates.ExecuteTemplate(w, "successPage", map[string]interface{}{
		"region_code": computeZone,
		"region_geo":  cityName,
		"flag_url":    flagURL,
	}); err != nil {
		log.Fatalf("failed to render template: %v", err)
	}
}
