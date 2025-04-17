package Config

import (
	"Api/Controller"
	"Api/Repositories"
	"Api/Service"
	"sync"
)

type IServiceContainer interface {
	InjectAuthController() Controller.AuthController
	InjectClinicalController() Controller.ClinicalController
}

type Kernel struct{}

func (k *Kernel) InjectAuthController() Controller.AuthController {
	userRepository := &Repositories.UserRepository{DB: DB}
	userService := &Service.UserService{IUserRepository: userRepository}
	authController := Controller.AuthController{IUserService: userService}

	return authController
}

func (k *Kernel) InjectClinicalController() Controller.ClinicalController {
	clinicalRepository := Repositories.ClinicalRepository{DB: DB}
	clinicalService := &Service.ClinicalService{IClinicalRepository: clinicalRepository}
	clinicalController := Controller.ClinicalController{IClinicalService: clinicalService}

	return clinicalController
}

var (
	k             *Kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &Kernel{}
		})
	}
	return k
}
