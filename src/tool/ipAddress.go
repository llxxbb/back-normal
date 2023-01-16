package tool

import (
	"fmt"
	"net"
	"strings"
)

func GetOutBoundIP() (ip string, err error) {
	out := "114.114.114.114:53"
	fmt.Printf("use %s test local out bound,", out)
	conn, err := net.Dial("udp", out)
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(" the result is: " + localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
