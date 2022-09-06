package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
)

func debugCounters(port uint64) {
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		logrus.Fatal(err)
	}
}
