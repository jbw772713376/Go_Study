package main
// 通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v!\n", ip[0], ip[1], ip[2], ip[3])
}

func stringerTest_study() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
