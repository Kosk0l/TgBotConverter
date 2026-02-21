package app

import (
	converterservice "github.com/Kosk0l/TgBotConverter/internal/Services/ConverterService"
	Dialogservice "github.com/Kosk0l/TgBotConverter/internal/Services/DialogService"
	jobservice "github.com/Kosk0l/TgBotConverter/internal/Services/jobService"
	userservice "github.com/Kosk0l/TgBotConverter/internal/Services/userService"
)

type Services struct {
	User 		*userservice.UserService
	Job			*jobservice.JobService
	Dialog 		*Dialogservice.DialogService
	Converter 	*converterservice.Converter
}

func initServices(infra *Infrastructure) (*Services) {
    return &Services{
        User:      userservice.NewUserService(infra.DB),
        Job:       jobservice.NewJobService(infra.Cache, infra.Minio),
        Dialog:    Dialogservice.NewDialogService(infra.Cache),
        Converter: converterservice.NewConverterService(),
    }
}