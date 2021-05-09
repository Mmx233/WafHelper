package modules

import (
	"github.com/shirou/gopsutil/net"
)

type checker struct{}

var Checker checker

func (*checker) LoadOk() (bool, error) {
	n1, err := net.Connections("tcp")
	if err != nil {
		return false, err
	}
	var counter uint
	for _, v := range n1 {
		if v.Laddr.Port == 80 || v.Laddr.Port == 443 {
			counter++
		}
	}
	return counter < 40, nil
}
