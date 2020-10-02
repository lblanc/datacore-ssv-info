package main

import (
    "fmt"
	"net/http"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}
func restRequest(dcsserver string, restserver string, user string, passwd string, target string)  http.Request {
	req, err := http.NewRequest("GET", "http://"+ restserver +"/RestService/rest.svc/1.0/"+ target, nil)
	if err != nil {
		fmt.Println("\nError when creating http request\nExit !")
		os.Exit(0)
	}
	req.Header.Add("ServerHost", dcsserver)
	req.Header.Add("Authorization", "Basic " + user +" "+ passwd)
	return *req
}

func bToT(b float64) float64 {
	return float64(b /1024 /1024 /1024 /1024)
}

func bToG(b float64) float64 {
	return float64(b /1024 /1024 /1024)
}