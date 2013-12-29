gojuice
=======

## Description

An open source tool for recieving and decrypting your latest JuiceSSH CloudSync backup.

* Obtains an OAUTH2 authentication token from Google API
* Authenticates with the JuiceSSH API
* Retrieves your latest encrypted CloudSync backup in JSON format
* Decrypts the backup using a user provided passphrase

## Build & Run

First download and install Go. 
On OSX this is as easy as:

```bash
$ brew install go
```

For other linux/windows/freebsd check http://golang.org

Then build it!

```bash
# Clone this repo
$ git clone git@bitbucket.org:sonelli/gojuice.git

# Fetch all of the Go module dependencies
$ cd gojuice
$ go get ./...

# Build it!
$ go build main.go
```

## Usage

```bash
 Usage:
 ./gojuice <decryption passphrase>
```


