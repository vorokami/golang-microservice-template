package app

import (
	"context"
	"golang-microservice-template/config"
	"golang-microservice-template/database"
	"golang-microservice-template/firebase"
	router "golang-microservice-template/internal/controller/http"
	"golang-microservice-template/internal/usecase"
	"golang-microservice-template/internal/usecase/repo/pointrepo"
	"golang-microservice-template/internal/usecase/repo/userrepo"
	"golang-microservice-template/logger"
	"net/http"
)

func Run(cfg *config.Config) {

	// init logger
	loggerFactory := logger.NewLoggerFactory(*cfg)

	// establish DB connection
	databaseUsersFactory := database.NewFactory(cfg.DbUsers, loggerFactory)
	databasePointsFactory := database.NewFactory(cfg.DbPoints, loggerFactory)

	// FCM
	_, err := firebase.InitFirebase(cfg.FIREBASE_AUTH_KEY)
	// if Firebase is not working - exit
	if err != nil {
		panic(err)
	}

	firestoreClient, err := firebase.GetApp().Firestore(context.Background())
	if err != nil {
		panic(err)
	}

	FcmClient, err := firebase.GetMessaging()
	if err != nil {
		panic(err)
	}

	// Use case
	templateMethodsUseCase := usecase.New(
		pointrepo.New(databasePointsFactory),
		userrepo.New(databaseUsersFactory, loggerFactory),
		firestoreClient,
		FcmClient,
	)

	appLogger, closeLogger, err := loggerFactory.NewLogger()
	if err != nil {
		panic(err)
	}
	defer closeLogger()

	// HTTP Server
	httpRouter := router.NewHTTPHandler(templateMethodsUseCase, appLogger)
	err = http.ListenAndServe(":"+cfg.Server.Port, httpRouter)
	if err != nil {
		panic(err)
	}

}
