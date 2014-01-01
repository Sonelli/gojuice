gojuice
=======

## Description

An open source tool for recieving and decrypting your latest JuiceSSH CloudSync backup.

* Obtains an OAUTH2 authentication token from Google API
* Authenticates with the JuiceSSH API
* Retrieves your latest encrypted CloudSync backup in JSON format
* Decrypts the backup using a user provided passphrase

## Binary Downloads

### Darwin (Apple Mac)

 * [bin/gojuice\_1.0.0\_darwin\_386.zip](gojuice_1.0.0_darwin_386.zip)
 * [bin/gojuice\_1.0.0\_darwin\_amd64.zip](gojuice_1.0.0_darwin_amd64.zip)

### FreeBSD

 * [bin/gojuice\_1.0.0\_freebsd\_386.zip](gojuice_1.0.0_freebsd_386.zip)
 * [bin/gojuice\_1.0.0\_freebsd\_amd64.zip](gojuice_1.0.0_freebsd_amd64.zip)
 * [bin/gojuice\_1.0.0\_freebsd\_arm.zip](gojuice_1.0.0_freebsd_arm.zip)

### Linux

 * [bin/gojuice\_1.0.0\_amd64.deb](gojuice_1.0.0_amd64.deb)
 * [bin/gojuice\_1.0.0\_armhf.deb](gojuice_1.0.0_armhf.deb)
 * [bin/gojuice\_1.0.0\_i386.deb](gojuice_1.0.0_i386.deb)
 * [bin/gojuice\_1.0.0\_linux\_386.tar.gz](gojuice_1.0.0_linux_386.tar.gz)
 * [bin/gojuice\_1.0.0\_linux\_amd64.tar.gz](gojuice_1.0.0_linux_amd64.tar.gz)
 * [bin/gojuice\_1.0.0\_linux\_arm.tar.gz](gojuice_1.0.0_linux_arm.tar.gz)

### MS Windows

 * [bin/gojuice\_1.0.0\_windows\_386.zip](gojuice_1.0.0_windows_386.zip)
 * [bin/gojuice\_1.0.0\_windows\_amd64.zip](gojuice_1.0.0_windows_amd64.zip)

### NetBSD

 * [bin/gojuice\_1.0.0\_netbsd\_386.zip](gojuice_1.0.0_netbsd_386.zip)
 * [bin/gojuice\_1.0.0\_netbsd\_amd64.zip](gojuice_1.0.0_netbsd_amd64.zip)
 * [bin/gojuice\_1.0.0\_netbsd\_arm.zip](gojuice_1.0.0_netbsd_arm.zip)

### OpenBSD

 * [bin/gojuice\_1.0.0\_openbsd\_386.zip](gojuice_1.0.0_openbsd_386.zip)
 * [bin/gojuice\_1.0.0\_openbsd\_amd64.zip](gojuice_1.0.0_openbsd_amd64.zip)

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
