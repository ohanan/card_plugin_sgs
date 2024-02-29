package service

import (
	"github.com/ohanan/card_proto/pkg/protoservice"
	"github.com/ohanan/card_proto/pkg/protoservice/proto"
)

type Game1V1 struct {
	*protoservice.Helper
}

func (g *Game1V1) Start() {
}
func (g *Game1V1) selectRole() {
	v := proto.Action{}
	v.Options = []*proto.ActionOption{
		{
			Option: &proto.ActionOption_Action{},
		},
	}
}
