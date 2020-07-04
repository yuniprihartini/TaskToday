package master

import (
	"database/sql"
	"ujian/master/controllers"
	"ujian/master/repositories"
	"ujian/master/usecase"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	serviceRepo := repositories.InitServiceRepoImpl(db)
	serviceUseCase := usecase.InitServiceUseCaseImpl(serviceRepo)
	controllers.ServicesControllers(r, serviceUseCase)

	customersRepo := repositories.InitCustomersRepoImpl(db)
	customersUseCase := usecase.InitCustomersUseCaseImpl(customersRepo)
	controllers.CustomersControllers(r, customersUseCase)

	progressRepo := repositories.InitProgressRepoImpl(db)
	progressUseCase := usecase.InitProgressUseCaseImpl(progressRepo)
	controllers.ProgressControllers(r, progressUseCase)
}
