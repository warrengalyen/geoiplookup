package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

// Update GeoLite2-Country.mmdb
func UpdateGeoLite2Country() {

	key := os.Getenv("LICENSEKEY")
	if key == "" && licenseKey != "" {
		key = licenseKey
	}

	if key == "" {
		fmt.Println("Error: GeoIP License Key not set.")
		os.Exit(1)
	}

	update_url := fmt.Sprintf("https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&license_key=%s&suffix=tar.gz", key)

	Debug("Updating GeoLite2-Country.mmdb")

	// check the output directory is writeable
	if _, err := os.Stat(*data_dir); os.IsNotExist(err) {
		os.MkdirAll(*data_dir, os.ModePerm)
	}

	if err := DownloadToFile("/tmp/GeoLite2-Country.tar.gz", update_url); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := ExtractDatabase(*data_dir, "/tmp/GeoLite2-Country.tar.gz"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := os.Remove("/tmp/GeoLite2-Country.tar.gz"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Extract the database from the tar.gz
func ExtractDatabase(dst string, targz string) error {

	Debug(fmt.Sprintf("Extracting %s", targz))

	re, _ := regexp.Compile(`GeoLite2\-Country\.mmdb$`)

	r, err := os.Open(targz)
	if err != nil {
		return err
	}

	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {
		// if no more files are found return
		case err == io.EOF:
			return nil
		// return any other error
		case err != nil:
			return err
		// if the header is nil, just skip it
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// check the file type
		switch header.Typeflag {

		case tar.TypeReg:

			if re.Match([]byte(target)) {
				outfile := filepath.Join(dst, "GeoLite2-Country.mmdb")

				f, err := os.OpenFile(outfile, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
				if err != nil {
					return err
				}

				Debug(fmt.Sprintf("Copy GeoLite2-Country.mmdb to %s", outfile))

				if _, err := io.Copy(f, tr); err != nil {
					return err
				}

				f.Close()
			}
		}
	}
}
