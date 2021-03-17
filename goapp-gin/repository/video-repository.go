// package repository

// import (
// 	"goapp-gin/entity"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type VideoRepository interface {
// 	Save(video entity.Video)
// 	Update(video entity.Video)
// 	Delete(video entity.Video)
// 	FindAll() []entity.Video
// }

// type database struct {
// 	connection *gorm.DB
// }

// func NewVideoRepository() VideoRepository {
// 	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kolkata"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to connect to database")
// 	}
// 	db.AutoMigrate(&entity.Video{}, &entity.Person{})
// 	return &database{
// 		connection: db,
// 	}
// }

// func (db *database) Save(video entity.Video) {
// 	db.connection.Create(&video)
// }
// func (db *database) Update(video entity.Video) {
// 	db.connection.Save(&video)
// }
// func (db *database) Delete(video entity.Video) {
// 	db.connection.Delete(&video)
// }
// func (db *database) FindAll() []entity.Video {
// 	var videos []entity.Video
// 	db.connection.Set("gorm:auto_preload", true).Find(&videos)
// 	return videos
// }

package repository

import (
	"goapp-gin/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
