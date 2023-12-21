package startup

type Paginated struct {
	Offset int
	Limit  int
}

type Search struct {
	Paginated
	Id        int
	Vanity    string
	Substring string
}

type ApplicantSearch struct {
	Id      int `db:"id"`
	Role    int `db:"role"`
	Startup int
	User    int `db:"user"`
	Status  string
}

type RoleSearch struct {
	Id      int
	Startup int
	Paginated
}

type GoalSearch struct {
	Id      int
	Startup int
	Role    int
	User    int
}
