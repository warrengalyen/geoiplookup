# GeoIPLookup - geoiplookup for GeoLite2-Country written in Go

geoiplookup is a geoiplookup replacement for the [free GeoLite2-Country](https://dev.maxmind.com/geoip/geoip2/geolite2/),

It currently only supports the free GeoLite2-Country database, and there is no planned support for the other types.


## Features

- Drop-in replacement for the now defunct geoiplookup utility
- Works with the current Maxmind database format (mmdd)
- IPv4, IPv6 and fully qualified domain name (FQDN) support
- Built-in database update support
- Options to return just the country iso (`US`) or country name (`United States`), rather than the full `GeoIP Country Edition: US, United States`


## Installing

Linux amd64 binaries are supplied with releases.

```
bunzip geoiplookup_linux_amd64.bz2
chmod 755 geoiplookup_linux_amd64
sudo mv geoiplookup_linux_amd64 /usr/local/bin/geoiplookup
```

If you with to replace an existing defunct implementation of geoiplookup, then simply name the file `geoiplookup`.


## Compiling from source

You must have golang installed. There is one external library required ([oschwald/geoip2-golang](https://github.com/oschwald/geoip2-golang)) which is downloaded automatically when you run `make`:

```
git clone git@github.com:warrengalyen/geoiplookup.git
cd geoiplookup
make
```

## Basic usage
```
Usage: geoiplookup [-i] [-c] [-d <database directory>] <ipaddress|hostname|db-update>
Options:
  -V	show version number
  -c	return country name
  -d string
    	database directory or file (default "/usr/share/GeoIP")
  -h	show help
  -i	return country iso code
  -v	verbose/debug output
Examples:
./geoiplookup 8.8.8.8			    Return the country ISO code and name
./geoiplookup -d ~/GeoIP 8.8.8.8	Use a different database directory
./geoiplookup -i 8.8.8.8			Return just the country ISO code
./geoiplookup -c 8.8.8.8			Return just the country name
./geoiplookup db-update			    Update the GeoLite2-Country database (do not run more than once a month)
```
