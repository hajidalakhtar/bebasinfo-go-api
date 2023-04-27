package usecase

import (
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/domain"
	"context"
	"github.com/google/uuid"
	"time"
)

type businessUsecase struct {
	businessRepo   domain.BusinessRepository
	contextTimeout time.Duration
}

func NewBusinessUsecase(u domain.BusinessRepository, timeout time.Duration) domain.BusinessUsecase {
	return &businessUsecase{
		businessRepo:   u,
		contextTimeout: timeout,
	}
}

func (b businessUsecase) Find(ctx context.Context, term string, sortBy string, limit int, offset int, openAt string) ([]domain.Business, error) {
	return b.businessRepo.Find(ctx, term, sortBy, limit, offset, openAt)
}

func (b businessUsecase) Store(ctx context.Context, bs *domain.Business) error {
	return b.businessRepo.Store(ctx, bs)

}

func (b businessUsecase) Update(ctx context.Context, bs *domain.Business, id uuid.UUID) error {
	return b.businessRepo.Update(ctx, bs, id)

}

func (b businessUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return b.businessRepo.Delete(ctx, id)

}
