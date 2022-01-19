package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

func nextIP(ip net.IP, inc uint) net.IP {
	i := ip.To4()
	v := uint(i[0])<<24 + uint(i[1])<<16 + uint(i[2])<<8 + uint(i[3])
	v += inc
	v3 := byte(v & 0xFF)
	v2 := byte((v >> 8) & 0xFF)
	v1 := byte((v >> 16) & 0xFF)
	v0 := byte((v >> 24) & 0xFF)
	return net.IPv4(v0, v1, v2, v3)
}

// func nextIPInInt(ip net.IP, inc uint) uint {
// 	i := ip.To4()
// 	v := uint(i[0])<<24 + uint(i[1])<<16 + uint(i[2])<<8 + uint(i[3])
// 	v += inc
// 	return v
// }

// func uIntToIP(v uint) net.IP {
// 	v3 := byte(v & 0xFF)
// 	v2 := byte((v >> 8) & 0xFF)
// 	v1 := byte((v >> 16) & 0xFF)
// 	v0 := byte((v >> 24) & 0xFF)
// 	return net.IPv4(v0, v1, v2, v3)
// }

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func stringToUinit(value string) uint32 {
	u32, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	wd := uint32(u32)
	return wd
}

func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func stringToIP(ip string) net.IP {
	stringToInt := stringToUinit(ip)
	return int2ip(stringToInt)
}
