package Config

import (
	"Api/Controller"
	"Api/Repositories"
	"Api/Service"
	"sync"
)

type IServiceContainer interface {
	InjectClinicalController() Controller.SesorController
}

type Kernel struct{}

func (k *Kernel) InjectClinicalController() Controller.SesorController {
	clinicalRepository := Repositories.SensorRepository{DB: DB}
	clinicalService := &Service.SensorService{ISensorRepository: clinicalRepository}
	clinicalController := Controller.SesorController{ISesorService: clinicalService}

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
