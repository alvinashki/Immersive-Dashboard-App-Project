package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	loginData "gp3/features/login/data"
	loginDelivery "gp3/features/login/delivery"
	loginUsecase "gp3/features/login/usecase"

	menteeData "gp3/features/mentee/data"
	menteeDelivery "gp3/features/mentee/delivery"
	menteeUsecase "gp3/features/mentee/usecase"

	userData "gp3/features/user/data"
	userDelivery "gp3/features/user/delivery"
	userUsecase "gp3/features/user/usecase"

	classData "gp3/features/class/data"
	classDelivery "gp3/features/class/delivery"
	classUsecase "gp3/features/class/usecase"

	logsData "gp3/features/logs/data"
	logsDelivery "gp3/features/logs/delivery"
	logsUsecase "gp3/features/logs/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	loginDataFactory := loginData.New(db)
	loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	loginDelivery.New(e, loginUsecaseFactory)

	menteeDataFactory := menteeData.New(db)
	menteeUsecaseFactory := menteeUsecase.New(menteeDataFactory)
	menteeDelivery.New(e, menteeUsecaseFactory)

	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	classDataFactory := classData.New(db)
	classUsecaseFactory := classUsecase.New(classDataFactory)
	classDelivery.New(e, classUsecaseFactory)

	logsDataFactory := logsData.New(db)
	logsUsecaseFactory := logsUsecase.New(logsDataFactory)
	logsDelivery.New(e, logsUsecaseFactory)
}
