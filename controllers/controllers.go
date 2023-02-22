package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HasnathJami/todo-crud/datasource"
	"github.com/HasnathJami/todo-crud/models"
	"github.com/HasnathJami/todo-crud/utils"
)

var tasks []models.Task = datasource.GetData()

func CreateTask(w http.ResponseWriter, r *http.Request){
  if r.Method == "POST"{
      w.Header().Set("Content-Type", "application-json")
      w.Header().Set("Access-Control-Allow-Method", "POST")

      var task models.Task

      //body, err := ioutil.ReadAll(r.Body)
      //utils.checkError(err)
      //json.Unmarshal(body, &task)
      utils.ParseBody(r, &task)
      tasks = append(tasks, task)

      res, _ := json.Marshal(task)
      w.Write(res)
     //w.Write([]byte(`{"message" : "Task Created Successfully"}`))
  }  else {
         w.WriteHeader(http.StatusBadRequest) 
         w.Write([]byte("Bad Request"))
  }
}

func GetTasks(w http.ResponseWriter, r *http.Request){
     if r.Method == "GET"{
       w.Header().Set("Content-Type","application-json")
       //w.Header().Set("Access-Control-Allow-Methods","GET")
       res, _ := json.Marshal(tasks)
       //w.Write([]byte(`{"message" : tasks}`))
       //w.Write(res)
       w.Write([]byte("data:" + string(res)))
       //fmt.Fprintf(w, "Fetch Success")
       //w.Write([]byte("message:" + "Messge fetched successfully"))
       //json.NewEncoder(w).Encode(&tasks)
     } else {
         w.WriteHeader(http.StatusBadRequest) 
         w.Write([]byte("Bad Request")) 
     }
    
}

func GetTaskById(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application-json")
    w.Header().Set("Access-Control-Allow-Method", "GET")
    if r.Method == "GET"{
        id := r.URL.Query().Get("id")

        for _, value := range tasks{
            if id == value.Id{
                json.NewEncoder(w).Encode(value)
                return
            }
        }
    json.NewEncoder(w).Encode("Task Not Found")

    } else{
         w.WriteHeader(http.StatusBadRequest) 
         w.Write([]byte("Bad Request")) 
    }
}

func UpdateTask(w http.ResponseWriter, r *http.Request){
    if r.Method == "PUT"{
      w.Header().Set("Content-Type", "application-json")
      w.Header().Set("Access-Control-Allow-Methods", "PUT")
      //id := r.FormValue("id")
      //json.NewEncoder(w).Encode(tasks)
      var task models.Task
      json.NewDecoder(r.Body).Decode(&task)
      //id := r.URL.Query()["id"][0] 
      id := r.URL.Query().Get("id")
      //id, err := strconv.ParseInt(param, 0, 0)
      //checkError(err)
    for index,value := range tasks {
        if id == value.Id{
            tasks = append(tasks[:index],tasks[index+1:]...)

            //json.NewEncoder(w).Encode(task)

            task.Id = r.URL.Query().Get("id")
            tasks = append(tasks, task)

            w.Write([]byte(`{"message" : "Task Updated Successfully"}`))
            return
        }
    }

    json.NewEncoder(w).Encode("Task Not Found")
    //fmt.Fprintf(w, "Update Success")

    } else{
        w.WriteHeader(http.StatusBadRequest) 
        w.Write([]byte("Bad Request")) 
    }
}

func DeleteTask(w http.ResponseWriter, r *http.Request){
    if r.Method == "DELETE" {
       w.Header().Set("Content-Type", "application-json")
       w.Header().Set("Access-Control-Allow-Methods", "DELETE")

      //var tasks []models.Task
      id := r.FormValue("id")
      //id := r.URL.Query()["id"][0] 
      //id, err := strconv.ParseInt(param, 0, 0)
      //checkError(err)
      json.NewDecoder(r.Body).Decode(&tasks)

      for index,value := range tasks{
         if id == value.Id {
            tasks = append(tasks[:index], tasks[index+1:]...)
             json.NewEncoder(w).Encode("Task deleted successfully")
             return
         } 
        }
         json.NewEncoder(w).Encode("Task Not Found")

     } else {
          w.WriteHeader(http.StatusBadRequest) 
          w.Write([]byte("Bad Request")) 
       } 
}