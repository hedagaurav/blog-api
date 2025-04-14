package seeder

import (
	"github.com/go-faker/faker/v4"
	"github.com/hedagaurav/blog-api/config"
	"github.com/hedagaurav/blog-api/models"
	"log"
	"time"
)

func seedData() {
	db := config.DB

	// Seed users
	for i := 0; i < 10000; i++ {
		user := models.User{}
		if err := faker.FakeData(&user); err != nil {
			log.Println("Error seeding user:", err)
			continue
		}
		if err := db.Create(&user).Error; err != nil {
			log.Println("Error creating user:", err)
		}
	}

	// Retrieve all user IDs
	var userIDs []uint
	if err := db.Model(&models.User{}).Pluck("id", &userIDs).Error; err != nil {
		log.Fatalf("Failed to get user IDs: %v", err)
	}

	// Seed 100,000 posts
	var posts []models.Post
	for i := 0; i < 100000; i++ {
		post := models.Post{
			Author:    userIDs[i%len(userIDs)],
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := faker.FakeData(&post); err != nil {
			log.Fatalf("Failed to generate fake post: %v", err)
		}
		posts = append(posts, post)

		// Insert in batches to avoid memory issues
		if len(posts) >= 1000 {
			if err := db.CreateInBatches(posts, 500).Error; err != nil {
				log.Fatalf("Error seeding posts: %v", err)
			}
			posts = nil
		}
	}

	// Insert remaining posts
	if len(posts) > 0 {
		if err := db.CreateInBatches(posts, 500).Error; err != nil {
			log.Fatalf("Error seeding remaining posts: %v", err)
		}
	}

	log.Println("Posts seeded!")

}
