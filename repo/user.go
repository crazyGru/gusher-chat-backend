package repo

import (
	"backend/user"
	"backend/user/block"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func (m *UserRepository) Find(params user.Search) ([]*user.User, error) {
	var result []*user.User
	query := "select id, username, admin from user where 1"
	if params.Id != 0 {
		query = query + " and id=:id"
	}
	if params.Admin {
		query = query + " and admin = 1"
	}
	rows, err := m.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u := &user.User{}
		if err := rows.StructScan(u); err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

func (m *UserRepository) Count(params user.Search) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *UserRepository) Save(u *user.User) {
	//TODO implement me
	panic("implement me")
}

func (m *UserRepository) Delete(u *user.User) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) AddBlock(id int, blocked_id int, note string) error {
	var b = block.UserBlock{
		UserId:    id,
		BlockedId: blocked_id,
		Date:      time.Now(),
		Note:      note,
	}

	query := `INSERT INTO user_blocks (user_id, blocked_id, date, note) VALUES (:user_id, :blocked_id, :date, :note)`
	_, err := u.db.NamedExec(query, b)
	return err
}

func (u *UserRepository) RemoveBlock(id int, blocked_id int) error {
	// Define the DELETE query. This query deletes the entry where both UserId and BlockedId match.
	query := `DELETE FROM user_blocks WHERE user_id = :user_id AND blocked_id = :blocked_id`

	// Create a map for query parameters
	params := map[string]interface{}{
		"user_id":    id,
		"blocked_id": blocked_id,
	}

	// Execute the query using NamedExec
	_, err := u.db.NamedExec(query, params)
	return err
}

func (u *UserRepository) CanInvite(id int, friendID int) bool {
	// Define a query to check if there is a block between id and friendID.
	query := `SELECT COUNT(*) FROM user_blocks WHERE user_id = $1 AND blocked_id = $2`

	// Variable to store the count result
	var count int

	// Execute the query
	err := u.db.Get(&count, query, id, friendID)

	// Check for errors and the count result
	if err != nil || count > 0 {
		// If there's an error or if the count is greater than 0, return false
		return false
	}

	// If there's no block, return true
	return true
}
