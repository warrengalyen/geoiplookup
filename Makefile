VERSION=`git describe --tags`
LDFLAGS=-ldflags "-X main.version=${VERSION}"

geoiplookup: geoiplookup.go
	go get github.com/oschwald/geoip2-golang
	go build ${LDFLAGS} geoiplookup.go
	strip geoiplookup

clean:
	rm -rf pkg src geoiplookup
