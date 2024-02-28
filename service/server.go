package service

import (
	"github.com/ohanan/card_proto/pkg/proto"
)

func NewSgs() proto.PluginXServer {
	return &SgsServer{}
}

type SgsServer struct {
}

func (s *SgsServer) GetPluginInfo(remote proto.CardXClient, req *proto.GetPluginInfo_Req, resp *proto.GetPluginInfo_Resp) {
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
