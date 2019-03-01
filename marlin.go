/*
	marlin - marlin.go
	ipgof
	Jun 2016
	MIT License
*/

package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/dsnet/compress/bzip2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func IsEmptyDir(name string) (bool, error) {
	entries, err := ioutil.ReadDir(name)
	if err != nil {
		return false, err
	}
	return len(entries) == 0, nil
}

func getPackageList() string {
	var retVal string

	files, err := ioutil.ReadDir("debs")
	if err != nil {
		fmt.Println("Couldn't find debs folder\n")
		//panic(err)
	}
	//check if debs folder is empty
	empty, err := IsEmptyDir("debs")
	if err != nil {
		log.Fatalln(err)
	}
	if empty {
		fmt.Printf("debs folder is empty\n")
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), "deb") {
			continue
		}
		filePath := path.Join("debs", file.Name())
		fileReader, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer fileReader.Close()

		control, err := readDebControlFile(fileReader)
		if err != nil {
			panic(err)
		}

		// Size
		control += "Size: "
		control += strconv.FormatInt(file.Size(), 10)
		control += "\n"

		// MD5Sum
		fileReader.Seek(0, 0)
		hash := md5.New()
		_, err = io.Copy(hash, fileReader)
		if err != nil {
			panic(err)
		}

		sum := fmt.Sprintf("%x", hash.Sum(nil))

		control += "MD5sum: "
		control += sum
		control += "\n"

		control += "Filename: "
		control += "./debs/" + file.Name()
		control += "\n"

		retVal += control + "\n"
	}

	return retVal
}

func getGzippedPackageList() ([]byte, []byte) {
	var b bytes.Buffer

	gz, err := bzip2.NewWriter(&b, nil)
	if err != nil {
		panic(err)
	}
	defer gz.Close()
	orig := []byte(getPackageList())
	_, err = gz.Write(orig)
	if err != nil {
		panic(err)
	}

	err = gz.Close()
	if err != nil {
		panic(err)
	}

	return b.Bytes(), orig

}

func main() {
	d1, d2 := getGzippedPackageList()
	err := ioutil.WriteFile("Packages.bz2", d1, 0644)
	check(err)
	err = ioutil.WriteFile("Packages", d2, 0644)
	check(err)
}
