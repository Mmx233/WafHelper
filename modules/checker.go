package modules

import (
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/net"
)

type checker struct{}

var Checker checker

func (*checker) TcpLoadOk() (bool, error) {
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

func (*checker) CpuLoadOk() (bool, error) {
	n2, err := load.Avg()
	if err != nil {
		return false, err
	}
	return n2.Load1 < 5, nil
}
