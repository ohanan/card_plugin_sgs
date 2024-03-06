package service

import (
	"github.com/ohanan/card_proto/pkg/protoservice"
	"github.com/ohanan/card_proto/pkg/protoservice/proto"
)

func NewSgs() protoservice.PluginXServer {
	return &SgsServer{}
}

type SgsServer struct {
}

func (s *SgsServer) StartGame(helper *protoservice.Helper, req *proto.StartGame_Req, resp *proto.StartGame_Resp) {
	go NewGame1V1(helper, req.Seed).Start()
}

func (s *SgsServer) GetPluginInfo(helper *protoservice.Helper, req *proto.GetPluginInfo_Req, resp *proto.GetPluginInfo_Resp) {
	resp.Info = &proto.PluginInfo{
		Version: &proto.Version{
			Major: 0,
			Minor: 0,
			Patch: 1,
		},
		Author:      "ohanan",
		Name:        "sgs",
		Description: "sgs",
	}
}
