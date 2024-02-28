package service

import (
	"github.com/ohanan/card_proto/pkg/proto"
)

var _ proto.PluginXServer = (*SGSServer)(nil)

type SGSServer struct {
}

func (S SGSServer) GetPluginInfo(remote proto.CardXClient, req *proto.GetPluginInfoReq, resp *proto.GetPluginInfoResp) {
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
