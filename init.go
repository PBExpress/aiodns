package main

import (
	_ "embed"

	"github.com/Workiva/go-datastructures/set"
)

//go:generate ./update-list.sh

//go:embed embed/tldnList
var tldnList string

//go:embed embed/bypassList
var bypassList string

//go:embed embed/specList
var specList string

func init() {
	defaultUpstream.Set("tls://dot.pub")
	defaultUpstream.Set("tls://dns.alidns.com")
	defaultUpstream.Set("https://doh.pub/dns-query")
	defaultUpstream.Set("https://dns.alidns.com/dns-query")

	// specUpstream.Set("quic://dns.adguard.com")
	// specUpstream.Set("https://odvr.nic.cz/doh")
	// specUpstream.Set("https://doh.opendns.com/dns-query")
	specUpstream.Set("https://8.8.8.8/dns-query")
	specUpstream.Set("https://9.9.9.11/dns-query")
	specUpstream.Set("https://162.159.36.1/dns-query")
	specUpstream.Set("https://149.112.112.11:5053/dns-query")
	specUpstream.Set("sdns://AQEAAAAAAAAADjIwOC42Ny4yMjAuMjIwILc1EUAgbyJdPivYItf9aR6hwzzI1maNDL4Ev6vKQ_t5GzIuZG5zY3J5cHQtY2VydC5vcGVuZG5zLmNvbQ")

	// fallUpstream.Set("tcp://9.9.9.11:9953")
	fallUpstream.Set("tcp://149.112.112.11:9953")
	fallUpstream.Set("tls://rubyfish.cn")
	fallUpstream.Set("https://cn-east.iqiqzz.com/cdn")
	fallUpstream.Set("https://cn-south.iqiqzz.com/cdn")
	fallUpstream.Set("https://1.15.50.48/verse")
	fallUpstream.Set("https://106.52.218.142/verse")
	// fallUpstream.Set("https://i.233py.com/dns-query")
	fallUpstream.Set("tls://dns.tuna.tsinghua.edu.cn:8853")
	fallUpstream.Set("https://dns.tuna.tsinghua.edu.cn:8443/dns-query")

	bootUpstream.Set("tls://223.5.5.5")
	bootUpstream.Set("tls://1.12.12.12")
	bootUpstream.Set("tcp://114.114.114.114")
	bootUpstream.Set("udp://114.114.115.115")
}

var initSpecDomains = set.New(
	"dl.google.com",
	"googleapis.cn",
	"googleapis.com",
	"gstatic.com",
)

var initSpecUpstreams = []string{
	"https://odvr.nic.cz/doh",
	"https://doh.dns.sb/dns-query",
	// "https://dns.twnic.tw/dns-query",
	"https://dns.adguard.com/dns-query",
	"https://doh.opendns.com/dns-query",
	"sdns://AQUAAAAAAAAACjguMjAuMjQ3LjIg0sJUqpYcHsoXmZb1X7yAHwg2xyN5q1J-zaiGG-Dgs7AoMi5kbnNjcnlwdC1jZXJ0LnNoaWVsZC0yLmRuc2J5Y29tb2RvLmNvbQ",
	"sdns://AQMAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMDo1NDQzINErR_JS3PLCu_iZEIbq95zkSV2LFsigxDIuUso_OQhzIjIuZG5zY3J5cHQuZGVmYXVsdC5uczEuYWRndWFyZC5jb20",
}
