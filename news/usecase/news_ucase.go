package usecase

import (
	"bebasinfo/domain"
	"context"
	"math"
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

func (n newsUsecase) Find(ctx context.Context, date string, source string, page int, limit int) ([]domain.News, domain.PaginatedResponse, error) {
	news, total, err := n.pgNewsRepository.Find(ctx, date, source, page, limit)
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	nextPage := page + 1
	if nextPage > totalPages {
		nextPage = 0
	}
	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 0
	}

	paginate := domain.PaginatedResponse{
		TotalItems:  total,
		TotalPages:  totalPages,
		CurrentPage: page,
		NextPage:    nextPage,
		PrevPage:    prevPage,
	}
	return news, paginate, err
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
