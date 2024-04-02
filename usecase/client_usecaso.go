package usecase

import (
	"context"
	"time"

	"github.com/waltherx/honda-backend/domain"
)

type clientUsecase struct {
	clientRepository domain.ClientRepository
	contextTimeout   time.Duration
}

func NewClientUsecase(clientRepository domain.ClientRepository, timeout time.Duration) domain.ClientUsecase {
	return &clientUsecase{
		clientRepository: clientRepository,
		contextTimeout:   timeout,
	}
}

func (tu *clientUsecase) Create(c context.Context, client *domain.Client) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.clientRepository.Create(ctx, client)
}

func (tu *clientUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Client, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.clientRepository.FetchByUserID(ctx, userID)
}
