package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUserName string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUserName: "Shivaganesh",
		authorizedPassword: "password",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return username == service.authorizedUserName &&
		password == service.authorizedPassword
}
