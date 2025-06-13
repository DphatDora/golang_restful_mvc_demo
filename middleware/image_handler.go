package middleware

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func isValidImage(ext string) bool {
	validExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}

	return validExtensions[ext]
}

func HandleAvatarUpload(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		c.Next()
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isValidImage(ext) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
		c.Abort()
		return
	}
	// Get user ID from URL parameter
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		c.Abort()
		return
	}

	os.MkdirAll("user-images", os.ModePerm)

	fileName := fmt.Sprintf("avatar_%s%s", id, ext)
	filePath := filepath.Join("user-images", fileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		c.Abort()
		return
	}

	imageURL := fmt.Sprintf("/user-images/%s", fileName)
	c.Set("image_url", imageURL)
	c.Next()
}
