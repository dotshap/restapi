package main

import (
	"fmt"
	"net"

	"github.com/go-redis/redis"
)

func configureRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "mypassword",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

func getRecord(client *redis.Client, key string) string {
	value, err := client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
		return redis.Nil.Error()
	}

	// var servers = Servers{}
	// err = json.Unmarshal([]byte(value), &servers)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return Servers{}
	// }

	return value
}

// func saveRecord(client *redis.Client, key string) Servers {
// 	json, err := json.Marshal(Servers{Ip: "192.168.1.1"})
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	err = client.Set(key, json, 0).Err()
// 	if err != nil {
// 		fmt.Println(err)
// 		return Servers{}
// 	}

// 	return getRecord(client, key)
// }

func saveIP(client *redis.Client, key string) string {

	maxIP := getMaxIP()
	if maxIP != "" {

		value := nextIP(net.ParseIP(maxIP), 1)
		fmt.Println("new IP is: ", value)

		nextIP := ip2int(value)
		fmt.Println("new IP in Int is : ", nextIP)
		err := client.Set(key, nextIP, 0).Err()
		if err != nil {
			fmt.Println(err)
			return redis.Nil.Error()
		} else {
			fmt.Println("New IP saved: ", nextIP)
		}

		saveMaxIP(client, nextIP)
		return getRecord(client, key)
	}
	return redis.Nil.Error()
}

// func nextIPInInt(iP net.IP, i int) {
// 	panic("unimplemented")
// }

func saveMaxIP(client *redis.Client, value uint32) {

	err := client.Set(maxIpAddress, value, 0).Err()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Max IP saved: ", value)
	}
}

func getMaxIP() string {

	value, err := client.Get(maxIpAddress).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// fmt.Println("getMaxIP value from DB: ", value)
	uinitValue := stringToUinit(value)
	// fmt.Println("getMaxIP stringToUinit: ", uinitValue)
	ip := int2ip(uinitValue)
	// fmt.Println("getMaxIP int2ip: ", ip)
	return ip.String()
}
