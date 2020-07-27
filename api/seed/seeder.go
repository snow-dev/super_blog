package seed

import (
	"github.com/jinzhu/gorm"
	models "github.com/snow-dev/super_blog/models"
	"log"
)

var users = []models.User{
	{
		Nickname: "Jon Snow",
		Email:    "snow@mail.dev",
		Password: "sesamo",
	},
	{
		Nickname: "John Titor",
		Email:    "titor@mail.dev",
		Password: "sesamo",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "The masters house",
		Content: "A history full of adventures",
	},

	models.Post{
		Title:   "The masters slaves",
		Content: "A history full of pleasures",
	},
}

func Load(db *gorm.DB) {
	err := db.DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed user table: %v", err)
		}

		posts[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
