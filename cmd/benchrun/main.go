package main

import (
	"log"
	"math/rand"

	iradix "github.com/hashicorp/go-immutable-radix"
	"github.com/miku/benchtrie"
)

func main() {
	// hosts := []string{
	// 	"ads.example.com",
	// 	"ads.example.net",
	// 	"ads.agency.org",
	// 	"klingt.net",
	// 	"git.klingt.net",
	// 	"foo.bar.cloudfront.com",
	// }
    // OR
    // hosts := []string{
    //  "192.168.123.255"
    hosts := benchtrie.GenerateIPv4Addr(100)
	//hosts := benchtrie.GenerateHosts(1000)
	r := iradix.New()
	for _, host := range hosts {
		log.Println(host)
		r, _, _ = r.Insert([]byte(host), nil)
	}

	key := hosts[rand.Intn(len(hosts))-1]
	v, ok := r.Get([]byte(key))
	log.Println("LOOKUP", key, v, ok)
}
