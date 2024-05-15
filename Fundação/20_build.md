## Building the application

#### Go build

`go build`

#### Building for windows

`GOOS=windows go build`

#### Building for linux

`GOOS=linux go build`

#### Building for linux with ARM64 architecture

`GOOS=linux GOARCH=arm64 go build`

#### List supported architectures/os

`go tool dist list`

#### Show the current os/architecture

`go env GOOS GOARCH`

#### [Supported architectures](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures#possible-platforms-for-goos-and-goarch)

#### Changing the output file name

`go build -o bin-name`
