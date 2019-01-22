package helper

import "log"

type Log struct {
}

func (l *Log) out(args ...string) {
	log.Println(args)
}
