#marlin
Marlin provides an easy way to generate a `Packages.gz` file for your Cydia/apt repository. 

This package runs on Mac, Linux, and Windows (dpkg and Cygwin not needed).

##Installing
Golang must be installed correctly on your computer in order to compile Marlin correctly.

    go get github.com/blakesmith/ar
    git clone https://github.com/cmelone/marlin.git
    cd marlin
    go build

This will create a binary named marlin in the current folder. Add your deb files to a folder named `debs` (you can configure this in `marlin.go`). Finally, run `marlin` and a `Packages.gz` file will be generated.

##TODO:
- Better error handling
- Ability to create a Release file
- Compile marlin for Windows, Mac, and Linux (and every architecture)

marlin takes most of its code from Evan Swick's [openrepo](//github.com/eswick/openrepo)
