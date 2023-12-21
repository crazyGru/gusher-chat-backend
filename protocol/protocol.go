package protocol

type completeFn func() bool

type ProtocolSearch struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
