package user

import "errors"

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Admin    bool   `db:"admin"`
}

type Search struct {
	Id    int  `db:"id"`
	Admin bool `db:"admin"`
}

type UserRepository interface {
	Find(params Search) ([]*User, error)
	Count(params Search) (int, error)
	Save(u *User)
	Delete(u *User)
	AddBlock(id int, blocked_id int, note string) error
	RemoveBlock(id int, blocked_id int) error
	CanInvite(id uint, friendID uint) bool
}
type Module struct {
	repo UserRepository
}

func NewModule(repo UserRepository) *Module {
	return &Module{repo: repo}
}

func (m *Module) FindUser(params Search) (*User, error) {
	list, err := m.repo.Find(params)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	return nil, errors.New("not found")
}

func (m *Module) CanInvite(pID uint, userID uint) bool {
	return m.repo.CanInvite(pID, userID)
}
