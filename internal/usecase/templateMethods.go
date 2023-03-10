package usecase

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/messaging"
)

type TemplateMethodsUseCase struct {
	repoPoints      PointRepo
	repoUsers       UserRepo
	firestoreClient *firestore.Client
	fcmClient       *messaging.Client
}

func New(rp PointRepo, ru UserRepo, fc *firestore.Client, mc *messaging.Client) *TemplateMethodsUseCase {
	return &TemplateMethodsUseCase{
		repoPoints:      rp,
		repoUsers:       ru,
		firestoreClient: fc,
		fcmClient:       mc,
	}
}

func (uc *TemplateMethodsUseCase) someMethodOne() (string, error) {
	return "", nil
}

func (uc *TemplateMethodsUseCase) someMethodTwo() (string, error) {
	return "", nil
}
