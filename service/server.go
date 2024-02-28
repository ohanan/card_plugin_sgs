package service

import (
	"github.com/ohanan/card_proto/pkg/proto"
)

var _ proto.PluginXServer = (*SGSServer)(nil)

type SGSServer struct {
}

func (s *SGSServer) ListMode(remote proto.CardXClient, req *proto.ListMode_Req, resp *proto.ListMode_Resp) {
	resp.Infos = []*proto.ModeInfo{
		{
			Name: "单挑",
		},
	}
}

func (s *SGSServer) GetPluginInfo(remote proto.CardXClient, req *proto.GetPluginInfo_Req, resp *proto.GetPluginInfo_Resp) {
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
