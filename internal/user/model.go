package user

//"pet_project_final/internal/taskService"

type User struct {
	//gorm.Model
	ID    uint32 `gorm:"primarykey"`
	Email string //`json:"email"`
	//Password string `json:"password"`
	//Task     []taskService.Task
}
