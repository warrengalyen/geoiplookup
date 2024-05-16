package main

import (
	"flag"
	"fmt"
	"os"
)

// Flags
var (
	data_dir       = flag.String("d", "/usr/share/GeoIP", "database directory or file")
	country        = flag.Bool("c", false, "return country name")
	iso            = flag.Bool("i", false, "return country iso code")
	showhelp       = flag.Bool("h", false, "show help")
	verbose_output = flag.Bool("v", false, "verbose/debug output")
	showversion    = flag.Bool("V", false, "show version number")
	version        = "dev"
	licenseKey     string // GeoLite2 license key for updating
)

// URLs
const (
	repo_url    = "https://github.com/warrengalyen/geoiplookup/releases"
	version_url = "https://api.github.com/repos/warrengalyen/geoiplookup/releases/latest"
)

func main() {

	flag.Parse()

	if *showversion {
		fmt.Println(fmt.Sprintf("Version: %s", version))
		os.Exit(1)
	}

	if len(flag.Args()) != 1 || *showhelp {
		Usage()
		os.Exit(1)
	}

	lookup := flag.Args()[0]

	if lookup == "db-update" {
		UpdateGeoLite2Country()
	} else {
		Lookup(lookup)
	}
}

// Print the help function
var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-i] [-c] [-d <database directory>] <ipaddress|hostname|db-update>\n", os.Args[0])
	fmt.Println("\nGeoiplookup uses the GeoLite2-Country database to find the Country that an IP address or hostname originates from.")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Fprintf(os.Stderr, "%s 8.8.8.8\t\t\tReturn the country ISO code and name\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "%s -d ~/GeoIP 8.8.8.8\t\tUse a different database directory\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "%s -i 8.8.8.8\t\t\tReturn just the country ISO code\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "%s -c 8.8.8.8\t\t\tReturn just the country name\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "%s db-update\t\t\tUpdate the GeoLite2-Country database (do not run more than once a month)\n", os.Args[0])
}
