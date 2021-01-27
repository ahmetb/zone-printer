package main

var (
	// Datacenter list is adopted from https://cloud.google.com/compute/docs/regions-zones/
	// and Cloud Run regions are at https://cloud.google.com/run/docs/locations
	datacenters = map[string]struct {
		location string
		flagURL  string // flag images must be public domain, as SVG, ideally from Wikipedia
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
			flagURL:  "https://upload.wikimedia.org/wikipedia/commons/5/5b/Flag_of_Hong_Kong.svg",
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
			flagURL:  "https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_South_Korea.svg",
		},
		"australia-southeast1": {
			location: "Sydney, Australia",
			flagURL:  "https://upload.wikimedia.org/wikipedia/en/b/b9/Flag_of_Australia.svg",
		},
	}
)
