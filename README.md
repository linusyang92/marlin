#Caca

this package, written in go, provides an easy way to generate a `Packages.gz` file to host a Cydia/APT repository

caca can run on windows, mac, and linux.

##Installing

you must have golang installed on your system, with the GOPATH and GOROOT set in your environment

    go get github.com/blakesmith/ar
    git clone https://github.com/cmelone/caca.git

once you are in the folder you can run `go build`. This will generate a binary named `caca`.

Create a folder named `debs` in the directory, and place as many deb files as you need.

after that, run the binary and a packages.gz file will be generated very quickly

##to do:

better error handling

###based on [openrepo](//github.com/eswick/openrepo) by @eswick
