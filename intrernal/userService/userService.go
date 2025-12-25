package userservice

type UserRepository interface {
	GetById()()
	Creare()()
	Update()()
	UpdateLastSeen()()
}

type UserService struct {
	repo UserRepository // Объект интерфейса - компановка
}

// Конструктор
func NewService(repo UserRepository) (*UserService) {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetByIdService() () {
	
}

func (u *UserService) CreateService() () {

}

func (u *UserService) UpdateService() () {

}

func (u *UserService) UpdateLastSeenService() () {

}

