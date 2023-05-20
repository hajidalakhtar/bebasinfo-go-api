package helper

import (
	"bebasinfo/domain"
	"math/rand"
	"time"
)

func ShuffleArray(arr []domain.News) []domain.News {
	rand.Seed(time.Now().UnixNano())

	// Create a map to store news by source
	newsMap := make(map[string][]domain.News)

	// Group news by source
	for _, news := range arr {
		newsMap[news.Source] = append(newsMap[news.Source], news)
	}

	// Shuffle each group of news
	for _, group := range newsMap {
		rand.Shuffle(len(group), func(i, j int) {
			group[i], group[j] = group[j], group[i]
		})
	}

	// Concatenate shuffled groups into a single array
	shuffledArray := []domain.News{}
	for _, group := range newsMap {
		shuffledArray = append(shuffledArray, group...)
	}

	return shuffledArray
}
