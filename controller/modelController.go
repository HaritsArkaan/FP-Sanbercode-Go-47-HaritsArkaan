package controller

import (
	"net/http"
	"time"

	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ModelInput struct {
	Name    string `json:"name"`
	BrandID uint   `json:"brand_id"`
}

// GetAllModel godoc
// @Summary Get all Model.
// @Description Get a list of Model.
// @Tags Model
// @Produce json
// @Success 200 {object} []models.Model
// @Router /models [get]
func GetAllModel(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var names []models.Model
	db.Find(&names)

	c.JSON(http.StatusOK, gin.H{"data": names})
}

// CreateModel godoc
// @Summary Create New Model.
// @Description Creating a new Model.
// @Tags Model
// @Param Body body ModelInput true "the body to create a new Model"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Model
// @Router /models [post]
func CreateModel(c *gin.Context) {
	// Validate input
	var input ModelInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Model
	model := models.Model{Name: input.Name, BrandID: input.BrandID}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&model)

	c.JSON(http.StatusOK, gin.H{"data": model})
}

// GetModelById godoc
// @Summary Get Model.
// @Description Get an Model by id.
// @Tags Model
// @Produce json
// @Param id path string true "Model id"
// @Success 200 {object} models.Model
// @Router /models/{id} [get]
func GetModelById(c *gin.Context) { // Get model if exist
	var name models.Model

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&name).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": name})
}

// GetPhoneByModelId godoc
// @Summary Get Phones.
// @Description Get all Phones by ModelId.
// @Tags Model
// @Produce json
// @Param id path string true "Model id"
// @Success 200 {object} []models.Phone
// @Router /models/{id}/phones [get]
func GetPhonesByModelId(c *gin.Context) { // Get model if exist
	var phones []models.Phone

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("model_id = ?", c.Param("id")).Find(&phones).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": phones})
}

// UpdateModel godoc
// @Summary Update Model.
// @Description Update Model by id.
// @Tags Model
// @Produce json
// @Param id path string true "Model id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param Body body ModelInput true "the body to update age rating category"
// @Success 200 {object} models.Model
// @Router /models/{id} [patch]
func UpdateModel(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var model models.Model
	if err := db.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input ModelInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Model
	updatedInput.Name = input.Name
	updatedInput.BrandID = input.BrandID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&model).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": model})
}

// DeleteModel godoc
// @Summary Delete one Model.
// @Description Delete a Model by id.
// @Tags Model
// @Produce json
// @Param id path string true "Model id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /models/{id} [delete]
func DeleteModel(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var model models.Model
	if err := db.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&model)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
