package routers

import (
	"net/http"

	"github.com/HasnathJami/todo-crud/controllers"
)

func Router() {
	//http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {w.Write([]byte("Everything is Ok"))})
	http.HandleFunc("/api/task", controllers.CreateTask)
	http.HandleFunc("/api/tasks", controllers.GetTasks)
	http.HandleFunc("/api/taskById", controllers.GetTaskById)
	http.HandleFunc("/api/updateTask", controllers.UpdateTask)
	http.HandleFunc("/api/deleteTask", controllers.DeleteTask)
}