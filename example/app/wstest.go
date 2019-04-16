package app

import (
	"github.com/wailovet/osmanthuswine/src/core"
	"gopkg.in/olahol/melody.v1"
)

type Wstest struct {
	core.WebSocket
}

func (that *Wstest) HandleConnect(*melody.Session) {
	//panic("implement me")
}

func (that *Wstest) HandlePong(*melody.Session) {
	//panic("implement me")
}

func (that *Wstest) HandleMessage(s *melody.Session, d []byte) {
	that.GetMelody().Broadcast(d)
	//panic("implement me")
}

func (that *Wstest) HandleMessageBinary(*melody.Session, []byte) {
	//panic("implement me")
}

func (that *Wstest) HandleSentMessage(*melody.Session, []byte) {
	//panic("implement me")
}

func (that *Wstest) HandleSentMessageBinary(*melody.Session, []byte) {
	//panic("implement me")
}

func (that *Wstest) HandleDisconnect(*melody.Session) {
	//panic("implement me")
}

func (that *Wstest) HandleError(*melody.Session, error) {
	//panic("implement me")
}
