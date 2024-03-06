package service

import (
	"github.com/ohanan/card_proto/pkg/protoservice"
	"github.com/ohanan/card_proto/pkg/protoservice/proto"
)

func NewGame1V1(helper *protoservice.Helper, seed uint64) *Game1V1 {
	return &Game1V1{
		Helper: helper,
		Rand:   helper.NewRand(seed),
	}
}

type Game1V1 struct {
	*protoservice.Helper
	*protoservice.Rand
}

func (g *Game1V1) Start() {
	g.start()
}
func (g *Game1V1) start() {
}

func (g *Game1V1) isFirstHand() bool {
	const (
		first = iota + 1
		second
		random
	)
selectFirstHand:
	result := g.AskAction("", []*proto.Action{
		(&proto.Action{}).WithCard(&proto.Action_SelectCard{
			Options: []*proto.Action_SelectCard_Option{
				{
					Option: &proto.Action_ActionOption{
						Id:   first,
						Name: "先手",
					},
					Card: nil,
				},
				{
					Option: &proto.Action_ActionOption{
						Id:   second,
						Name: "后手",
					},
					Card: nil,
				},
				{
					Option: &proto.Action_ActionOption{
						Id:   random,
						Name: "随机",
					},
					Card: nil,
				},
			},
		}),
	}).Result
	if result.Type != proto.UserActionType_PRIMARY {
		goto selectFirstHand
	}
	switch result.Id {
	case first:
		return true
	case second:
		return false
	case random:
		return g.UintN(2) > 0
	}
	goto selectFirstHand
}
