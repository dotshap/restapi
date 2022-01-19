package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	// for http request
	"github.com/gorilla/mux"

	// redis
	"github.com/go-redis/redis"
)

var client *redis.Client

// init Method

func init() {
	client = configureRedis()
	handleMaxIp()
}

func handleMaxIp() {
	ip := getMaxIP()
	if ip == "" {
		intIp := ip2int(net.ParseIP("10.0.0.1"))
		saveMaxIP(client, intIp)
	}
}

// MAIN Fucntion

func main() {

	// Route Hanlders/ Endpoints
	routeHandler()
}

func getServers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parms := mux.Vars(r)
	key := parms["publicKey"]
	fmt.Println("PublicKey is: ", key)
	response := getRecord(client, key)

	fmt.Println("existing record in ip is :", stringToIP(response))

	if response == redis.Nil.Error() {
		response = saveIP(client, key)
	}
	stringtoIP := stringToIP(response)
	json.NewEncoder(w).Encode(stringtoIP)
}

func routeHandler() {
	r := mux.NewRouter()
	r.HandleFunc("/api/client/{publicKey}", getServers).Methods("GET")
	log.Fatal(http.ListenAndServe(portName, r))
}
