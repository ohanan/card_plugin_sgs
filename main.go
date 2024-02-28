package main

import (
	"github.com/ohanan/card_plugin_sgs/service"
	"github.com/ohanan/card_proto/pkg/plugin"
)

func main() {
	plugin.ServePlugin("sgs", &service.SGSServer{})
}
