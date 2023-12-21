package protocol

import (
	"backend/shared"
)

type Task struct {
	Id          int                `json:"id" db:"id" example:"123"`
	StartupRole shared.StartupRole `json:"startupRole" db:"role" example:"developer"`
	Title       string             `json:"title" db:"title"  example:"Develop company website"`
	Content     string             `json:"content" db:"content"  example:"html content..."`
	DependsOn   DepList            `json:"dependsOn" db:"depends_on"  example:"123,444,653,4"`
	CompleteFn  string             `json:"complete_fn" db:"complete_fn" example:"websiteDeveloped"`
}
