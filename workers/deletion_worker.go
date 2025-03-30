package workers

import (
	"file-sharing-backend/config"
	"file-sharing-backend/models"
	"log"
	"time"
	"os"
)

func DeleteExpiredFiles() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Checking for expired files...")

		var expiredFiles []models.File
		now := time.Now()

		if err := config.DB.Where("expires_at <= ?", now).Find(&expiredFiles).Error; err != nil {
			log.Println("Error fetching expired files:", err)
			continue
		}

		for _, file := range expiredFiles {
			if err := os.Remove(file.Filepath); err != nil {
				log.Println("Error deleting file from storage:", file.Filepath, err)
				continue
			}

			if err := config.DB.Delete(&file).Error; err != nil {
				log.Println("Error deleting file record from database:", err)
				continue
			}

			log.Println("Deleted expired file:", file.Filename)
		}
	}
}
