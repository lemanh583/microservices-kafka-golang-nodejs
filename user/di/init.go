package di

import (
	"user-services/handle"
	"user-services/repository"
	"user-services/service"

	"gorm.io/gorm"
)

type ServiceDI struct {
	UserRepository repository.UserRepository
	UserService    service.UserService
	UserHandle     handle.UserHandle
	MailService    service.MailService
}

func Initialized(db *gorm.DB) *ServiceDI {
	// repository
	ur := repository.NewUserRepository(db)
	// service
	ms := service.NewMailService()
	us := service.NewUserService(ur)
	ks := service.NewKafkaService("localhost:9092", "send-mail")

	// handle
	uh := handle.NewUserHandle(us, ms, ks)
	return &ServiceDI{
		UserRepository: ur,
		UserService:    us,
		UserHandle:     uh,
		MailService:    ms,
	}
}
