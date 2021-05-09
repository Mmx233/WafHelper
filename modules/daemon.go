package modules

import (
	"log"
	"time"
)

func Daemon() {
	for {
		time.Sleep(time.Second * 5)
		ok, err := Checker.LoadOk()
		if err != nil {
			log.Println(err)
			continue
		}
		if !ok {
			if err = VeryNginx.OpenBrowserVerity(); err != nil {
				log.Println(err)
				continue
			}
		} else {
			if err = VeryNginx.CloseBrowserVerity(); err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
