package modules

import (
	"log"
	"time"
)

func Daemon() {
	for {
		time.Sleep(time.Second * 5)
		ok, err := Checker.TcpLoadOk()
		if err != nil {
			log.Println(err)
			continue
		}
		if !ok {
			if ok2, err := Checker.CpuLoadOk(); !ok2 {
				if err != nil {
					log.Println(err)
					continue
				}
				if err = VeryNginx.OpenBrowserVerity(); err != nil {
					log.Println(err)
				}
			}
			continue
		} else {
			if err = VeryNginx.CloseBrowserVerity(); err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
