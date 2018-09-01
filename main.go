package main

import (
	"net"
	"time"
	"log"
	"flag"

	"github.com/mdlayher/arp"
)

var (
	interfaceFlag = flag.String("i", "", "Please specify interface")
	networkFlag = flag.String("n", "", "Please specify network address including cider")
	timeOutFlag = flag.Duration("t", 1000*time.Millisecond,"Please specify the request timeout time in milliseconds (default 1000ms)")
)

func ipCount(cider string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cider)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips[1:len(ips)-1], nil
}

func inc(ip net.IP) {
	for i := len(ip) -1; i >=0; i-- {
		ip[i]++
		if ip[i] > 0 {
			break
		}
	}
}

func main() {
	flag.Parse()

	ifIndex, err := net.InterfaceByName(*interfaceFlag)
	if err != nil {
		log.Fatal(err)
	}

	ipStrings, err := ipCount(*networkFlag)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := arp.Dial(ifIndex)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	log.Println(ipStrings)

	for _, ipString := range ipStrings {

		if err := conn.SetDeadline(time.Now().Add(*timeOutFlag)); err != nil {
			continue
		}

		targetIp := net.ParseIP(ipString).To4()
		hwAddr, _ := conn.Resolve(targetIp)
		log.Printf("%v -> %v", targetIp, hwAddr)
	}
}