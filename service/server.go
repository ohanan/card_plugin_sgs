package service

import (
	"github.com/ohanan/card_proto/pkg/protoservice"
)

func NewSgs() protoservice.PluginXServer {
	return &SgsServer{}
}

type SgsServer struct {
}

func (s *SgsServer) GetPluginInfo(remote protoservice.CardXClient, req *protoservice.GetPluginInfo_Req, resp *protoservice.GetPluginInfo_Resp) {
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
