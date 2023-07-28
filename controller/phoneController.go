package controller

import (
	"net/http"
	"strings"
	"time"

	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhoneInput struct {
	BrandID  uint   `json:"brand_id"`
	ColorID  uint   `json:"color_id"`
	ModelID  uint   `json:"model_id"`
	ReviewID uint   `json:"review_id"`
	Price    string `json:"price"`
	Storage  uint   `json:"storage"`
}

// GetAllPhone godoc
// @Summary Get all Phone.
// @Description Get a list of Phone.
// @Tags Phone
// @Produce json
// @Param price query string false "price filter (case insensitive search)"
// @Param sortByprice query string false "Sort by price (asc or desc)"
// @Param sortByBrandID query string false "Sort by BrandID (asc or desc)"
// @Param sortByColorID query string false "Sort by ColorID (asc or desc)"
// @Param sortByModelD query string false "Sort by ModelD (asc or desc)"
// @Param sortByReviewID query string false "Sort by ReviewID (asc or desc)"
// @Param sortByCreatedAt query string false "Sort by created_at (asc or desc)"
// @Success 200 {object} []models.Phone
// @Router /phones [get]
func GetAllPhone(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	price := c.Query("price")
	sortByPrice := c.Query("sortByPrice")
	sortByBrandID := c.Query("sortByBrandID")
	sortByColorID := c.Query("sortByColorID")
	sortByModelD := c.Query("sortByModelD")
	sortByReviewID := c.Query("sortByReviewID")
	sortByCreatedAt := c.Query("sortByCreatedAt")

	var names []models.Phone

	query := db

	if price != "" {
		query = query.Where("LOWER(price) LIKE ?", "%"+strings.ToLower(price)+"%")
	}

	if sortByPrice != "" {
		if strings.ToLower(sortByPrice) == "asc" {
			query = query.Order("price asc")
		} else if strings.ToLower(sortByPrice) == "desc" {
			query = query.Order("price desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortByTitle parameter"})
			return
		}
	}

	if sortByBrandID != "" {
		if strings.ToLower(sortByBrandID) == "asc" {
			query = query.Order("brand_id asc")
		} else if strings.ToLower(sortByBrandID) == "desc" {
			query = query.Order("brand_id desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortByRatingID parameter"})
			return
		}
	}

	if sortByColorID != "" {
		if strings.ToLower(sortByColorID) == "asc" {
			query = query.Order("color_id asc")
		} else if strings.ToLower(sortByColorID) == "desc" {
			query = query.Order("color_id desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortByRatingID parameter"})
			return
		}
	}

	if sortByModelD != "" {
		if strings.ToLower(sortByModelD) == "asc" {
			query = query.Order("model_id asc")
		} else if strings.ToLower(sortByModelD) == "desc" {
			query = query.Order("model_id desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortByRatingID parameter"})
			return
		}
	}

	if sortByReviewID != "" {
		if strings.ToLower(sortByReviewID) == "asc" {
			query = query.Order("review_id asc")
		} else if strings.ToLower(sortByReviewID) == "desc" {
			query = query.Order("review_id desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortByRatingID parameter"})
			return
		}
	}

	if sortByCreatedAt != "" {
		if strings.ToLower(sortByCreatedAt) == "asc" {
			query = query.Order("created_at asc")
		} else if strings.ToLower(sortByCreatedAt) == "desc" {
			query = query.Order("created_at desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortByCreatedAt parameter"})
			return
		}
	}

	if err := query.Find(&names).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": names})
}

// CreatePhone godoc
// @Summary Create New Phone.
// @Description Creating a new Phone.
// @Tags Phone
// @Param Body body PhoneInput true "the body to create a new Phone"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Phone
// @Router /phones [post]
func CreatePhone(c *gin.Context) {
	// Validate input
	var input PhoneInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Phone
	phone := models.Phone{BrandID: input.BrandID, ColorID: input.ColorID, ModelID: input.ModelID, ReviewID: input.ReviewID, Price: input.Price, Storage: input.Storage}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&phone)

	c.JSON(http.StatusOK, gin.H{"data": phone})
}

// GetPhoneById godoc
// @Summary Get Phone.
// @Description Get an Phone by id.
// @Tags Phone
// @Produce json
// @Param id path string true "Phone id"
// @Success 200 {object} models.Phone
// @Router /phones/{id} [get]
func GetPhoneById(c *gin.Context) { // Get model if exist
	var name models.Phone

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&name).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": name})
}

// UpdatePhone godoc
// @Summary Update Phone.
// @Description Update Phone by id.
// @Tags Phone
// @Produce json
// @Param id path string true "Phone id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param Body body PhoneInput true "the body to update age rating category"
// @Success 200 {object} models.Phone
// @Router /phones/{id} [patch]
func UpdatePhone(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var phone models.Phone
	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input PhoneInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Phone
	updatedInput.BrandID = input.BrandID
	updatedInput.ColorID = input.ColorID
	updatedInput.ModelID = input.ModelID
	updatedInput.ReviewID = input.ReviewID
	updatedInput.Price = input.Price
	updatedInput.Storage = input.Storage
	updatedInput.UpdatedAt = time.Now()

	db.Model(&phone).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": phone})
}

// DeletePhone godoc
// @Summary Delete one Phone.
// @Description Delete a Phone by id.
// @Tags Phone
// @Produce json
// @Param id path string true "Phone id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /phones/{id} [delete]
func DeletePhone(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var phone models.Phone
	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&phone)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
