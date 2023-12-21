package role

import "time"

type Goal struct {
	id        int
	parent    int
	startup   int
	role      int
	user      int
	title     string
	status    string
	created   time.Time
	modified  time.Time
	activated time.Time
	reached   time.Time
	approved  time.Time
	blocked   time.Time
	hidden    time.Time
	notified  bool
}

func (g *Goal) Id() int {
	return g.id
}

func (g *Goal) Title() string {
	return g.title
}
