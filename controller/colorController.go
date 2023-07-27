package controller

import (
	"net/http"
	"time"

	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ColorInput struct {
	Name string `json:"name"`
}

// GetAllColor godoc
// @Summary Get all Color.
// @Description Get a list of Color.
// @Tags Color
// @Produce json
// @Success 200 {object} []models.Color
// @Router /colors [get]
func GetAllColor(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var names []models.Color
	db.Find(&names)

	c.JSON(http.StatusOK, gin.H{"data": names})
}

// CreateColor godoc
// @Summary Create New Color.
// @Description Creating a new Color.
// @Tags Color
// @Param Body body ColorInput true "the body to create a new Color"
// @Produce json
// @Success 200 {object} models.Color
// @Router /colors [post]
func CreateColor(c *gin.Context) {
	// Validate input
	var input ColorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Color
	color := models.Color{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&color)

	c.JSON(http.StatusOK, gin.H{"data": color})
}

// GetColorById godoc
// @Summary Get Color.
// @Description Get an Color by id.
// @Tags Color
// @Produce json
// @Param id path string true "Color id"
// @Success 200 {object} models.Color
// @Router /colors/{id} [get]
func GetColorById(c *gin.Context) { // Get model if exist
	var name models.Color

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&name).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": name})
}

// GetPhoneByColorId godoc
// @Summary Get Phones.
// @Description Get all Phones by ColorId.
// @Tags Color
// @Produce json
// @Param id path string true "Color id"
// @Success 200 {object} []models.Phone
// @Router /colors/{id}/phones [get]
func GetPhonesByColorId(c *gin.Context) { // Get model if exist
	var phones []models.Phone

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("color_id = ?", c.Param("id")).Find(&phones).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": phones})
}

// UpdateColor godoc
// @Summary Update Color.
// @Description Update Color by id.
// @Tags Color
// @Produce json
// @Param id path string true "Color id"
// @Param Body body ColorInput true "the body to update age rating category"
// @Success 200 {object} models.Color
// @Router /colors/{id} [patch]
func UpdateColor(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var color models.Color
	if err := db.Where("id = ?", c.Param("id")).First(&color).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input ColorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Color
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&color).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": color})
}

// DeleteColor godoc
// @Summary Delete one Color.
// @Description Delete a Color by id.
// @Tags Color
// @Produce json
// @Param id path string true "Color id"
// @Success 200 {object} map[string]boolean
// @Router /colors/{id} [delete]
func DeleteColor(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var color models.Color
	if err := db.Where("id = ?", c.Param("id")).First(&color).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&color)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
