package seeder

import (
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/domain"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	_ = CreateBusiness(db, domain.Business{
		Alias:       "fork-boise",
		Name:        "Fork",
		ImageURL:    "https://s3-media4.fl.yelpcdn.com/bphoto/P9mNoEUBfeSbgmJEla4jmQ/o.jpg",
		IsClosed:    false,
		URL:         "https://www.yelp.com/biz/fork-boise?adjust_creative=DSj6I8qbyHf-Zm2fGExuug&utm_campaign=yelp_api_v3&utm_medium=api_v3_business_search&utm_source=DSj6I8qbyHf-Zm2fGExuug",
		ReviewCount: 2069,
		//CategoriesID: []uint{
		//	1,
		//},
		Categories: domain.Categories{
			{
				Alias: "newamerican",
				Title: "American (New)",
			},
			{
				Alias: "breakfast_brunch",
				Title: "Breakfast & Brunch",
			},
			{
				Alias: "burgers",
				Title: "Burgers",
			},
		},
		Rating: 4.0,
		Coordinates: domain.Coordinates{
			Latitude:  43.616389,
			Longitude: -116.203056,
		},
		Transactions: []string{"delivery"},
		Price:        "$$",
		Location: domain.Location{
			Address1:       "199 N 8th St",
			Address2:       "",
			Address3:       "",
			City:           "Boise",
			ZipCode:        "83702",
			Country:        "US",
			State:          "ID",
			DisplayAddress: []string{"199 N 8th St", "Boise, ID 83702"},
		},
		Phone:        "+12082871700",
		DisplayPhone: "(208) 287-1700",
		Distance:     314.400925836215,
	})
	_ = CreateBusiness(db, domain.Business{
		Alias:       "bistro-lavoile",
		Name:        "Bistro La Voile",
		ImageURL:    "https://s3-media1.fl.yelpcdn.com/bphoto/HoDdwJpChwCK3qTtTgQ2fg/o.jpg",
		IsClosed:    false,
		URL:         "https://www.yelp.com/biz/bistro-la-voile-boston-2?adjust_creative=DSj6I8qbyHf-Zm2fGExuug&utm_campaign=yelp_api_v3&utm_medium=api_v3_business_search&utm_source=DSj6I8qbyHf-Zm2fGExuug",
		ReviewCount: 487,
		Categories: domain.Categories{
			{
				Alias: "french",
				Title: "French",
			},
			{
				Alias: "wine_bars",
				Title: "Wine Bars",
			},
			{
				Alias: "breakfast_brunch",
				Title: "Breakfast & Brunch",
			},
		},
		Rating: 4.5,
		Coordinates: domain.Coordinates{
			Latitude:  42.3536556540011,
			Longitude: -71.0443869543746,
		},
		Transactions: []string{"delivery"},
		Price:        "$$",
		Location: domain.Location{
			Address1:       "261 Newbury St",
			Address2:       "",
			Address3:       "",
			City:           "Boston",
			ZipCode:        "02116",
			Country:        "US",
			State:          "MA",
			DisplayAddress: []string{"261 Newbury St", "Boston, MA 02116"},
		},
		Phone:        "+16172223333",
		DisplayPhone: "(617) 222-3333",
		Distance:     172.940757994041,
	})
	_ = CreateBusiness(db, domain.Business{
		Alias:       "giordanos-chicago",
		Name:        "Giordano's",
		ImageURL:    "https://s3-media4.fl.yelpcdn.com/bphoto/nPZutUsxRbWlK04Y1wA2JQ/o.jpg",
		IsClosed:    false,
		URL:         "https://www.yelp.com/biz/giordanos-chicago-4?adjust_creative=DSj6I8qbyHf-Zm2fGExuug&utm_campaign=yelp_api_v3&utm_medium=api_v3_business_search&utm_source=DSj6I8qbyHf-Zm2fGExuug",
		ReviewCount: 8046,
		Categories: domain.Categories{
			{
				Alias: "pizza",
				Title: "Pizza",
			},
			{
				Alias: "italian",
				Title: "Italian",
			},
		},
		Rating: 4.0,
		Coordinates: domain.Coordinates{
			Latitude:  41.8961837,
			Longitude: -87.6284699,
		},
		Transactions: []string{"delivery"},
		Price:        "$$",
		Location: domain.Location{
			Address1:       "223 W Jackson Blvd",
			Address2:       "",
			Address3:       "",
			City:           "Chicago",
			ZipCode:        "60606",
			Country:        "US",
			State:          "IL",
			DisplayAddress: []string{"223 W Jackson Blvd", "Chicago, IL 60606"},
		},
		Phone:        "+13124630000",
		DisplayPhone: "(312) 463-0000",
		Distance:     1078.4950495646095,
	})

}
