package main

import (
	"fmt"
)

func main() {
	r := DocumentRoot()
	r.NewValue("host").Value("localhost")
	ip := r.NewChild("ips")
	ip.NewValue("lo").Value("127.0.0.1")
	ip.NewValue("eth0").Value("192.23.178.4")
	fmt.Println(r)

	c1 := r.GetChild("host")
	fmt.Printf("child \"host\" = %s\n", c1)
	c2 := r.GetChild("ips")
	fmt.Printf("child \"ips\"  = %s\n", c2)

	v1 := r.GetValue("host")
	fmt.Printf("value \"host\" = %s\n", v1)
	v2 := r.GetValue("ips")
	fmt.Printf("value \"ips\"  = %s\n", v2)

	v1.Value("home")
	fmt.Println(r)

	os := r.NewChild("specs").NewChild("os")
	os.NewValue("type").Value("linux")
	os.NewValue("distribution").Value("centos")
	os.NewValue("version").Value("7")
	fmt.Println(r)
}
