package main

// type TaskController struct {
// 	TaskService TaskService
// }

// func (tc *TaskController) GetTasks(c *gin.Context) {
// 	tasks := tc.TaskService.GetTasks()
// 	c.JSON(http.StatusOK, tasks)
// }

// type TaskService struct {
// 	TaskRepo TaskRepository
// }

// func (ts *TaskService) GetTasks() Tasks {
// 	return ts.TaskRepo.GetTasks()
// }

// type TaskRepository interface {
// 	GetTasks() Tasks
// }

// type InMemoryTaskRepository struct {
// 	Tasks Tasks
// }

// func (repo *InMemoryTaskRepository) GetTasks() Tasks {
// 	return repo.Tasks
// }

func main() {}

// 	taskRepo := &InMemoryTaskRepository{
// 		Tasks: Tasks{
// 			{ID: 1, Title: "Task 1", Status: TaskStatusTodo},
// 			{ID: 2, Title: "Task 2", Status: TaskStatusTodo},
// 		},
// 	}

// 	taskService := TaskService{
// 		TaskRepo: taskRepo,
// 	}

// 	taskController := &TaskController{
// 		TaskService: taskService,
// 	}

// 	router := gin.Default()

// 	router.GET("/tasks", taskController.GetTasks)

// 	cfg, err := config.New()
// 	if err != nil {
// 		os.Exit(1)
// 	}

// 	router.Run(fmt.Sprintf(":%d", cfg.Port))
// }
