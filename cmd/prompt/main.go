package main

import (
	"github.com/ohanan/card_plugin_sgs/service"
	"github.com/ohanan/card_proto/cmd/prompt"
)

func main() {
	prompt.Run(service.NewSgs())
}
