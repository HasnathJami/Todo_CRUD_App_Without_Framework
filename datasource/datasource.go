package datasource

import (
	"github.com/HasnathJami/todo-crud/models"
)

func GetData() []models.Task {
    var tasks []models.Task
    tasks = append(tasks, models.Task{Id:"1", Name:"Coding", Priority:"High", Done:true}) 
	  tasks = append(tasks, models.Task{Id:"2", Name:"Gym", Priority:"Medium", Done:false}) 
    
	  return tasks
  }
 


