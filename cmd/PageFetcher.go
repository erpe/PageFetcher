package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL PATH \n", os.Args[0])
	}

	destinationPath := os.Args[2]

	urlString := os.Args[1]

	url, err := url.Parse(urlString)

	if err != nil {
		log.Fatal(err)
	}

	destinationFileName := path.Base(url.String())

	resp, err := http.Get(url.String())

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		blob, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		destination := filepath.Join(destinationPath, destinationFileName)

		if err := ioutil.WriteFile(destination, blob, 0644); err != nil {
			log.Fatal(err)
		}

		log.Println("saved file: ", destination)
	}
}
