package helper

type Log struct {
}

func (l *Log) out(args ...string) {
	println(args)
}
