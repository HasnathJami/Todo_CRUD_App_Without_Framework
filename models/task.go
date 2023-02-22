package models

type Task struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Priority string `json:"priority"`
	Done     bool   `json:"done"`
}