package Initializers

import (
	"fmt"
	"log"
	"os"
	"time"
	"wan-api-kol-event/Models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	//Get database url from environment variables (defined in .env file)
	var dsn string = os.Getenv("DB_URL")

	//Connect with postgres
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             50 * time.Millisecond, // Slow SQL threshold
				LogLevel:                  logger.Warn,           // Log level
				IgnoreRecordNotFoundError: false,                 // Dont ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,                 // Include params in the SQL log
				Colorful:                  true,                  // Disable color
			},
		),
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// // Auto-migrate the database (Creates the table if it doesn't exist)
	// err = DB.AutoMigrate(&Models.Kol{})
	// if err != nil {
	// 	log.Fatal(" Failed to migrate database:", err)
	// }

	// fmt.Println("Database migration completed!")

	// // Insert dummy data if the table is empty
	// SeedKols()
}

func SeedKols() {
	var count int64
	DB.Model(&Models.Kol{}).Count(&count)

	if count == 0 {

		dummyKols := []Models.Kol{}
		for i := 1; i <= 30; i++ {
			dummyKols = append(dummyKols, Models.Kol{
				KolID:                int64(i),
				UserProfileID:        int64(i + 100),
				Language:             "English",
				Education:            "Bachelor's Degree",
				ExpectedSalary:       int64(50000 + (i * 1000)),
				ExpectedSalaryEnable: true,
				ChannelSettingTypeID: int64(i % 5),
				IDFrontURL:           fmt.Sprintf("https://example.com/idfront-%d.jpg", i),
				IDBackURL:            fmt.Sprintf("https://example.com/idback-%d.jpg", i),
				PortraitURL:          fmt.Sprintf("https://example.com/portrait-%d.jpg", i),
				RewardID:             int64(i),
				PaymentMethodID:      int64(i % 3),
				TestimonialsID:       int64(i % 7),
				VerificationStatus:   i%2 == 0,
				Enabled:              true,
				ActiveDate:           time.Now(),
				Active:               true,
				CreatedBy:            "admin",
				CreatedDate:          time.Now(),
				ModifiedBy:           "admin",
				ModifiedDate:         time.Now(),
				IsRemove:             false,
				IsOnBoarding:         i%3 == 0,
				Code:                 fmt.Sprintf("KOL-%d", i),
				PortraitRightURL:     fmt.Sprintf("https://example.com/portrait-right-%d.jpg", i),
				PortraitLeftURL:      fmt.Sprintf("https://example.com/portrait-left-%d.jpg", i),
				LivenessStatus:       i%2 == 0,
			})
		}

		// Bulk insert the dummy data
		if err := DB.Create(&dummyKols).Error; err != nil {
			log.Fatal("Failed to seed dummy data:", err)
		}

	} else {
		fmt.Println(" KOL table already has data, skipping seeding.")
	}
}
