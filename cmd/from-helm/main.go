package main

import (
	"github.com/tom721/from-helm-server/pkg/server"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func main() {
	logf.SetLogger(zap.New())
	server.Start()
}
