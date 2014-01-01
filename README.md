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
# Fetch the source
$ go get github.com/Sonelli/gojuice

# Build it!
$ go install gojuice
```

## Usage

```bash
 Usage:
 $GOPATH/bin/gojuice <decryption passphrase>
```
