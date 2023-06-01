package usecase

import (
	"bebasinfo/domain"
	"bebasinfo/news/usecase/helper"
	"context"
	"errors"
	"github.com/google/uuid"
	"math"
	"time"
)

type newsUsecase struct {
	pgNewsRepository      domain.PosgresqlNewsRepository
	rssNewsRepository     domain.RSSNewsRepository
	apiNewsRepository     domain.APINewsRepository
	apiNewsDataRepository domain.APINewsDataRepository
	contextTimeout        time.Duration
}

func NewNewsUsecase(pnr domain.PosgresqlNewsRepository, rnr domain.RSSNewsRepository, inr domain.APINewsRepository, ndr domain.APINewsDataRepository, timeout time.Duration) domain.NewsUsecase {
	return &newsUsecase{
		pgNewsRepository:      pnr,
		rssNewsRepository:     rnr,
		apiNewsRepository:     inr,
		apiNewsDataRepository: ndr,
		contextTimeout:        timeout,
	}
}

func (n newsUsecase) Find(ctx context.Context, newsId uuid.UUID) ([]domain.News, error) {
	news, _, err := n.pgNewsRepository.Find(ctx, newsId, "", nil, 1, 10)
	return news, err
}

func (n newsUsecase) Search(ctx context.Context, date string, source []string, page int, limit int) ([]domain.News, domain.PaginatedResponse, error) {
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

func (n newsUsecase) Store(ctx context.Context, newsResource string, category string, source []string) ([]domain.News, error) {
	var news []domain.News
	switch newsResource {
	case "rss":
		news, _ = n.rssNewsRepository.GetFromRSS(ctx, source)
	case "newsapi":
		news, _ = n.apiNewsRepository.GetFromAPI(ctx, category)
	case "newsdata":
		news, _, _ = n.apiNewsDataRepository.GetFromAPI(ctx, category, "")
	default:
		return nil, errors.New("news resource not found")

	}
	//randNews := helper.ShuffleArray(news)
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

func (n newsUsecase) StoreMultiplePagesFromNewsDataApi(ctx context.Context, category string) ([]domain.News, error) {
	var merged []domain.News
	var nextPage string

	for i := 0; i < 5; i++ {
		news, page, err := n.apiNewsDataRepository.GetFromAPI(ctx, category, nextPage)
		if err != nil {
			return nil, err
		}
		nextPage = page
		merged = append(merged, news...)
	}

	randNews := helper.ShuffleArray(merged)
	var storedNews []domain.News

	for _, newsItem := range randNews {
		_, err := n.pgNewsRepository.FindByTitle(ctx, newsItem.Title)
		if err != nil {
			err = n.pgNewsRepository.Store(ctx, newsItem)
			if err != nil {
				return nil, err
			}
			storedNews = append(storedNews, newsItem)
		}
	}

	return storedNews, nil
}
