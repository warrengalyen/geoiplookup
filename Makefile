VERSION=$(git describe --tags)

geoiplookup: geoiplookup.go
	go get github.com/oschwald/geoip2-golang
	go build -ldflags "-X main.version=${VERSION}" geoiplookup.go
	strip geoiplookup

clean:
	rm -rf pkg src geoiplookup
