package model

// Task is the request model for our api.
type Task struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

// Verify is used to verify if two tasks have same data.
func (t Task) CheckValue(t2 Task) bool {
	if t.Name == t2.Name && t.Status == t2.Status {
		return true
	}
	return false
}
