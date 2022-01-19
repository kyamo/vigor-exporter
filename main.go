package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	vigor "github.com/T4cC0re/vigor-node-exporter/vigor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var username = flag.String("username", "", "username to authenticate to the Vigor")
var password = flag.String("password", "", "password to authenticate to the Vigor")
var ip = flag.String("ip", "", "ip the Vigor is reachable on")

var v *vigor.Vigor

func loginIfError(err error) {
	if err != nil {
		print(err)
		v.Login(*username, *password)
	}
}

func main() {
	flag.Parse()

	var err error
	v, err = vigor.New(*ip)
	if err != nil {
		panic(err)
	}

	v.Login(*username, *password)

	v.UpdateStatus()
	v.FetchStatus()

	go func() {
		for {
			time.Sleep(5 * time.Second)

			loginIfError(v.UpdateStatus())
			loginIfError(v.FetchStatus())
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9103", nil))
}
