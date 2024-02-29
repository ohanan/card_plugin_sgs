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

func (s *SgsServer) StartMode(helper *protoservice.Helper, req *proto.StartMode_Req, resp *proto.StartMode_Resp) {
	// TODO implement me
	panic("implement me")
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
		Modes: []*proto.ModeInfo{
			{
				Name:        "",
				Description: "",
			},
		},
	}
}
