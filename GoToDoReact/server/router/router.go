package router

import(
	"github.com/Shadabdullah/GoToDoReact/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{

	router := mux.NewRouter()
	router.HandleFunc("/api/task",middleware.GetAllTasks).Methods("GET","OPTIONS")
	router.HandleFunc("/api/tasks",middleware.CreatTask).Methods("POST","OPTIONS")
	router.HandleFunc("/api/tasks/{id}",middleware.TaskComplete).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/undoTask/{id}",middleware.UndoTask).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}",middleware.DeleteTask).Methods("DELETE","OPTIONS")
	router.HandleFunc("/api/deleteAllTask",middleware.DeleteAllTasks).Methods("PUT","OPTIONS")
}