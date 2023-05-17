package usecase

import (
	"bebasinfo/domain"
	"context"
	"github.com/google/uuid"
	"math"
	"time"
)

type newsUsecase struct {
	pgNewsRepository  domain.PosgresqlNewsRepository
	rssNewsRepository domain.RSSNewsRepository
	apiNewsRepository domain.APINewsRepository
	contextTimeout    time.Duration
}

func NewNewsUsecase(pnr domain.PosgresqlNewsRepository, rnr domain.RSSNewsRepository, inr domain.APINewsRepository, timeout time.Duration) domain.NewsUsecase {
	return &newsUsecase{
		pgNewsRepository:  pnr,
		rssNewsRepository: rnr,
		apiNewsRepository: inr,
		contextTimeout:    timeout,
	}
}

func (n newsUsecase) Find(ctx context.Context, newsId uuid.UUID) ([]domain.News, error) {
	news, _, err := n.pgNewsRepository.Find(ctx, newsId, "", "", 1, 10)
	return news, err
}

func (n newsUsecase) Search(ctx context.Context, date string, source string, page int, limit int) ([]domain.News, domain.PaginatedResponse, error) {
	news, total, err := n.pgNewsRepository.Find(ctx, uuid.Nil, date, source, page, limit)
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

// -------
// newsResoure = rss, newsapi, all
// source(for rss) = detik, kompas, cnn, all
// category(for newsapi) = technology, health, sports, business, science, entertainment, general
// -------
func (n newsUsecase) Store(ctx context.Context, newsResource string, category string, source string) ([]domain.News, error) {
	var news []domain.News

	switch newsResource {
	case "rss":
		news, _ = n.rssNewsRepository.GetFromRSS(ctx, source)
	case "newsapi":
		news, _ = n.apiNewsRepository.GetFromAPI(ctx, category)
	default:
		newsApi, _ := n.rssNewsRepository.GetFromRSS(ctx, source)
		news = append(news, newsApi...)
		newsRss, _ := n.apiNewsRepository.GetFromAPI(ctx, category)
		news = append(news, newsRss...)
	}

	for _, newsItem := range news {
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
