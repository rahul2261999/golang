package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lc0.dev:3000/learn?coursename=reactjs&payemntid=1234567890"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myUrl)

	parseUrl, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println("scheme:", parseUrl.Scheme)
	fmt.Println("host:", parseUrl.Host)
	fmt.Println("path:", parseUrl.Path)
	fmt.Println("raw query:", parseUrl.RawQuery)
	fmt.Println("query:", parseUrl.Query())

	qparams := parseUrl.Query()

	for key, value := range qparams {
		fmt.Printf("%s: %s\n", key, value)
	}

	// prepare a new url with new parameters

	newUrl := &url.URL{
		Scheme:   "https",
		Host:     "lc0.dev:3000",
		Path:     "/learn",
		RawQuery: "username=john&password=doe",
	}

	anotherUrl := newUrl.String()

	fmt.Println("New url:", anotherUrl)

}
