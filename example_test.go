package main

import (
	"net/http"
	"testing"
	"time"
	"os"
	"log"
	"net"
)

func TestMain(m *testing.M) {
	time.Sleep(time.Second * 45) // let riak startup
	m.Run()
}

var client = &http.Client{
	Timeout: time.Second*5,
}

func TestCantHitRiakOnLocalhost(t *testing.T) {
	// get ip from hostname
	hitRiakWithAddress("localhost",t)
}

func TestRiakWorks(t *testing.T) {
	// get ip from hostname
	ip := getIpFromHostNameHostname()
	hitRiakWithAddress(ip,t)
}

func hitRiakWithAddress(address string, t *testing.T){
	resp, err := client.Get("http://" +address + ":8098/buckets?buckets")

	if resp != nil {
		t.Log("riak works! yay!")
	}

	if err != nil {
		t.Fatalf("Couldn't talk to riak because of error %v\n", err)
	}
}


func getIpFromHostNameHostname() string{
	hostName, err := os.Hostname()
	log.Printf("Hostname: %s\n", hostName)
	addresses, err := net.LookupHost(hostName)

	if err != nil {
		log.Panicf("Cant get hostname and so we cant talk to riak %v", err)
	}
	log.Printf("Addresses: %s\n", addresses)

	return addresses[0]
}
