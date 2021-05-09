package modules

import (
	"WafHelper/util"
	"fmt"
	"log"
	"strings"
)

type veryNginx struct {
	browserVerityOpened bool
	path                string
}

var VeryNginx = veryNginx{
	path: "/opt/verynginx/verynginx/configs/config.json",
}

func (a *veryNginx) reloadNginx() error {
	_, e := util.Shell.Exec("nginx -s reload")
	return e
}

func (a *veryNginx) changeBrowserVerity(b bool) error {
	config, err := util.File.Read(a.path)
	if err != nil {
		return err
	}
	config = strings.Replace(config, `"browser_verify_enable":`+fmt.Sprint(!b), `"browser_verify_enable":`+fmt.Sprint(b), 1)
	err = util.File.Write(a.path, config)
	if err == nil {
		err = a.reloadNginx()
	}
	return err
}

func (a *veryNginx) OpenBrowserVerity() error {
	err := a.changeBrowserVerity(true)
	if err == nil {
		if !a.browserVerityOpened {
			log.Println("Verity ON")
		}
		a.browserVerityOpened = true
	}
	return err
}

func (a *veryNginx) CloseBrowserVerity() error {
	if !a.browserVerityOpened {
		return nil
	}
	err := a.changeBrowserVerity(false)
	if err == nil {
		if a.browserVerityOpened {
			log.Println("Verity OFF")
		}
		a.browserVerityOpened = false
	}
	return err
}
