package controller

import (
	"net/http"
	"time"

	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhoneInput struct {
	BrandID uint   `json:"brand_id"`
	ColorID uint   `json:"color_id"`
	ModelID uint   `json:"model_id"`
	Price   string `json:"price"`
	Storage uint   `json:"storage"`
}

// GetAllPhone godoc
// @Summary Get all Phone.
// @Description Get a list of Phone.
// @Tags Phone
// @Produce json
// @Success 200 {object} []models.Phone
// @Router /phones [get]
func GetAllPhone(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var names []models.Phone
	db.Find(&names)

	c.JSON(http.StatusOK, gin.H{"data": names})
}

// CreatePhone godoc
// @Summary Create New Phone.
// @Description Creating a new Phone.
// @Tags Phone
// @Param Body body PhoneInput true "the body to create a new Phone"
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
	phone := models.Phone{BrandID: input.BrandID, ColorID: input.ColorID, ModelID: input.ModelID, Price: input.Price, Storage: input.Storage}
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
