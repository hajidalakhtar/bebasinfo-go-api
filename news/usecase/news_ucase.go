package usecase

import (
	"bebasinfo/domain"
	"context"
	"time"
)

type newsUsecase struct {
	pgNewsRepository  domain.PosgresqlNewsRepository
	rssNewsRepository domain.RSSNewsRepository
	contextTimeout    time.Duration
}

func NewNewsUsecase(pnr domain.PosgresqlNewsRepository, rnr domain.RSSNewsRepository, timeout time.Duration) domain.NewsUsecase {
	return &newsUsecase{
		pgNewsRepository:  pnr,
		rssNewsRepository: rnr,
		contextTimeout:    timeout,
	}
}

func (n newsUsecase) Find(ctx context.Context, date string, source string) ([]domain.News, error) {
	return n.pgNewsRepository.Find(ctx, date, source)
}

func (n newsUsecase) Store(ctx context.Context, source string) ([]domain.News, error) {
	var news []domain.News
	newsFromRSS, _ := n.rssNewsRepository.GetFromRSS(ctx, source)
	for _, newsItem := range newsFromRSS {
		_, err := n.pgNewsRepository.FindByTitle(ctx, newsItem.Title)
		if err != nil {
			err = n.pgNewsRepository.Store(ctx, newsItem)
			if err != nil {
				return nil, err
			}
			news = append(news, newsItem)
		}
	}
	return news, nil
}

//
//func (b businessUsecase) Store(ctx context.Context, bs *domain.Business) error {
//	return b.businessRepo.Store(ctx, bs)
//
//}
//
//func (b businessUsecase) Update(ctx context.Context, bs *domain.Business, id uuid.UUID) error {
//	return b.businessRepo.Update(ctx, bs, id)
//
//}
//
//func (b businessUsecase) Delete(ctx context.Context, id uuid.UUID) error {
//	return b.businessRepo.Delete(ctx, id)
//
//}
