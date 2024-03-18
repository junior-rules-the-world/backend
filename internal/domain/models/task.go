package models

type Task struct {
	ID          int    `json:"task_id" db:"task_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	TeamID      int    `json:"team_id" db:"team_id"`
	EventID     int    `json:"event_id" db:"event_id"`
}

type TaskRepository interface {
	Create(task *Task) error
	Update(id int, updated *Task) error
	FindByTeamID(id int) ([]Task, error)
	FindByEventID(id int) ([]Task, error)
}

type TaskUsecase interface {
	Create(task *Task) error
	Update(id int, updated *Task) error
	Close(id int) error
	RequestToClose(id int) error
	FindByTeamID(id int) ([]Task, error)
	FindByEventID(id int) ([]Task, error)
}
