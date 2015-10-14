package room

import (
	"github.com/yangsf5/moba/server/proto"
)

func (r *Room) NotifyHCPlayerInfos() {
	players := []proto.PlayerInfo{}
	for _, user := range r.sessions {
		player := proto.PlayerInfo{
			Name:      user.Name(),
			CurrentHP: user.GetHP(),
		}
		players = append(players, player)
	}
	msg := &proto.RCPlayerInfos{players}
	r.Broadcast(proto.Encode(r.serviceName, msg))
}

func (r *Room) NotifyHCBattleStatus() {
	msg := &proto.RCBattleStatus{r.battleStatus}
	r.Broadcast(proto.Encode(r.serviceName, msg))
}