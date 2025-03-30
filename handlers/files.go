package handlers

import (
	"file-sharing-backend/config"
	"file-sharing-backend/models"
	"log"
	"net/http"
	"os"
	"fmt"
	"path/filepath"
	"io"
	"time"
	"github.com/gin-gonic/gin"
	"context"
	"encoding/json"
	"file-sharing-backend/workers"
)

// File Upload
func FileUploadHandler(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userIDUint, _ := userID.(uint)

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found"})
		return
	}
	defer file.Close()

	savePath := filepath.Join("uploads", header.Filename)

	tx := config.DB.Begin()

	newFile := models.File{
		UserID:     userIDUint,
		Filename:   header.Filename,
		Filepath:   savePath,
		Visibility: "private",
		UploadedAt: time.Now(),
		ExpiresAt:  time.Now().Add(30 * 24 * time.Hour),
	}

	if err := tx.Create(&newFile).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file info to DB"})
		return
	}

	outFile, err := os.Create(savePath)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	defer outFile.Close()

	if _, err = io.Copy(outFile, file); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing file"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file": newFile})
}

// get all files (redis cached)
func ListFiles(c *gin.Context) {
	ctx := context.Background()

	cachedFiles, err := config.RedisClient.Get(ctx, "files_metadata").Result()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"source": "cache", "files": cachedFiles})
		return
	}

	var files []models.File
	if err := config.DB.Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}

	serializedFiles, _ := json.Marshal(files)
	config.RedisClient.Set(ctx, "files_metadata", serializedFiles, 5*time.Minute)

	c.JSON(http.StatusOK, gin.H{"source": "database", "files": files})
}

// download File
func DownloadFile(c *gin.Context) {
	filename := c.Param("filename")
	var file models.File

	if err := config.DB.Where("filename = ?", filename).First(&file).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.File(file.Filepath)
}

// delete File
func DeleteFile(c *gin.Context) {
	filename := c.Param("filename")
	var file models.File

	if err := config.DB.Where("filename = ?", filename).First(&file).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	if err := os.Remove(file.Filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}

	if err := config.DB.Delete(&file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove file record"})
		return
	}

	ctx := context.Background()
	config.RedisClient.Del(ctx, "files_metadata")

	log.Printf("File deleted: %s", file.Filepath)
	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

// Search Files by filename, user_id
func SearchFiles(c *gin.Context) {
	ctx := context.Background()
	queryKey := fmt.Sprintf("search:%s:%s:%s", c.Query("filename"), c.Query("uploaded_at"), c.Query("file_type"))

	cachedData, err := config.RedisClient.Get(ctx, queryKey).Result()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"source": "cache", "files": cachedData})
		return
	}

	var files []models.File
	query := config.DB.Model(&models.File{})

	if filename := c.Query("filename"); filename != "" {
		query = query.Where("filename ILIKE ?", "%"+filename+"%")
	}
	if uploadedAt := c.Query("uploaded_at"); uploadedAt != "" {
		query = query.Where("DATE(uploaded_at) = ?", uploadedAt)
	}
	if fileType := c.Query("file_type"); fileType != "" {
		query = query.Where("filename ILIKE ?", "%."+fileType)
	}

	if err := query.Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files", "details": err.Error()})
		return
	}

	serializedFiles, _ := json.Marshal(files)
	config.RedisClient.Set(ctx, queryKey, serializedFiles, 5*time.Minute)

	c.JSON(http.StatusOK, gin.H{"source": "database", "files": files})
}

func TriggerDeletion(c *gin.Context) {
	go workers.DeleteExpiredFiles()
	c.JSON(http.StatusOK, gin.H{"message": "File deletion triggered"})
}