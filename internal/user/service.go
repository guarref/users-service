package user

type UserService struct {
	urepo UserRepository
	//srepo taskService.TaskRepository
}

func NewUserService(urepo UserRepository) *UserService {
	return &UserService{urepo: urepo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.urepo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.urepo.GetAllUsers()
}

func (s *UserService) GetUserByID(id uint32) (User, error) {
	return s.urepo.GetUserByID(id)
}

func (s *UserService) UpdateUserByID(id uint32, user User) (User, error) {
	return s.urepo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint32) error {
	return s.urepo.DeleteUserByID(id)
}
