package main

import (
	"fmt"
	//"bufio"
    "io/ioutil"
    "net/http"
	"encoding/json"
	"flag"
	"os"
	"time"
	"strconv"
	"strings"
	//"reflect"
)



func main() {
	// Define flags
	dcsserver := flag.String("dcsserver","","DataCore SANsymphony server")
	restserver := flag.String("restserver","","DataCore SANsymphony REST api server")
	user := flag.String("user","","DataCore SANsymphony user")
	passwd := flag.String("passwd","","DataCore SANsymphony password")

	//Parse
	flag.Parse()

	//Check flags
	if *dcsserver == "" {
		fmt.Println("\nError! No DataCore SANsymphony server set")
		fmt.Println("Add -dcsserver=x where x is IP or DNS of DataCore SANsymphony server")
		fmt.Println("Exit!")
		os.Exit(0)
	} else if *restserver == "" {
		fmt.Println("\nWarning ! No DataCore SANsymphony REST api server set")
		fmt.Println("Add -restserver=x where x is IP or DNS of DataCore SANsymphony REST api server")
		fmt.Println("We will use DataCore SANsymphony server as REST api server\n")
		restserver = dcsserver
	}
	if *user == "" {
		fmt.Println("\nError ! No DataCore SANsymphony server user set")
		fmt.Println("Add -user=x where x is user for DataCore SANsymphony server")
		fmt.Println("\nExit!")
		os.Exit(0)
	}
	if *passwd == "" {
		fmt.Println("\nError ! No DataCore SANsymphony server password set")
		fmt.Println("Add -passwd=x where x is password for DataCore SANsymphony server")
		fmt.Println("\nExit!")
		os.Exit(0)
	}
	
	//req_hosts := restRequest(*dcsserver , *restserver , *user , *passwd , "hosts")
	//req_servers := restRequest(*dcsserver , *restserver , *user , *passwd , "servers")
	req_servergroups := restRequest(*dcsserver , *restserver , *user , *passwd , "servergroups")
	//fmt.Println(reflect.TypeOf(*req))

	client := &http.Client{
	}

	// get hosts
	/*
	resp_hosts, err := client.Do(&req_hosts)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		os.Exit(0)
	} else if resp_hosts.Status != "200 OK"{
		fmt.Printf("The HTTP request failed with error %s\n", resp_hosts.Status)
		os.Exit(0)
	} 
	  
	data, _ := ioutil.ReadAll(resp_hosts.Body)
	var hosts []host
	json.Unmarshal([]byte(data), &hosts)
	  
	//fmt.Println(hosts)
	


	//get servers
	resp_servers, err := client.Do(&req_servers)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		os.Exit(0)
	} else if resp_servers.Status != "200 OK"{
		fmt.Printf("The HTTP request failed with error %s\n", resp_servers.Status)
		os.Exit(0)
	}
	
	data, _ = ioutil.ReadAll(resp_servers.Body)
	var servers []server
	json.Unmarshal([]byte(data), &servers)
	
	*/


	//get servergroups
	resp_servergroups, err := client.Do(&req_servergroups)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		os.Exit(0)
	} else if resp_servergroups.Status != "200 OK"{
		fmt.Printf("The HTTP request failed with error %s\n", resp_servergroups.Status)
		os.Exit(0)
	}
	
	data, _ := ioutil.ReadAll(resp_servergroups.Body)
	var servergroups []servergroup
	json.Unmarshal([]byte(data), &servergroups) 


	//print out

	for _, item := range servergroups{
		fmt.Printf("\n%s (Local group: %v)\n",item.Caption,item.OurGroup)
		for _, item2 := range item.ExistingProductKeys{
			s := strings.TrimPrefix(item2.ExpirationDate, "/Date(")
			s = s[:len(s)-7]
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
			}
			fmt.Println(item2.ProductName)
			fmt.Println(item2.LastFive)
			fmt.Println(time.Unix(i/1000, 0))
			
			
		}

	}


	/*
	for _, item := range hosts{
		client := &http.Client{
		}
		target_preffered := "hosts/" + item.Id + "/preferredservers"

		req_preferredservers := restRequest(*dcsserver , *restserver , *user , *passwd , target_preffered)
		resp_preferredservers, err := client.Do(&req_preferredservers)
		data, err = ioutil.ReadAll(resp_preferredservers.Body)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
		
		var preferredservers []preferredserver
		json.Unmarshal([]byte(data), &preferredservers)
		if len(preferredservers) > 0 {
			fmt.Println(preferredservers[0].Caption)
		} else {
			fmt.Println(item.Caption, "|", "None")
		}
	   
	}
	*/
	
	
	
	

}