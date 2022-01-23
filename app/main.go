package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Yan-Bin-Lin/DCreater/app/router"
	"github.com/Yan-Bin-Lin/DCreater/app/setting"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	runServe("main", router.HostSwitch{Engine: router.MainRouter()})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func runServe(serve string, hs router.HostSwitch) {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", setting.Servers[serve].Port),
		Handler:      hs,
		ReadTimeout:  setting.Servers[serve].ReadTimeout,
		WriteTimeout: setting.Servers[serve].WriteTimeout,
	}
	g.Go(func() error {
		return s.ListenAndServe()
	})
}
