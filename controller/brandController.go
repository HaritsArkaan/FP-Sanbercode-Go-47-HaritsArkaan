package controller

import (
	"net/http"
	"time"

	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BrandInput struct {
	Name string `json:"name"`
}

// GetAllBrand godoc
// @Summary Get all Brand.
// @Description Get a list of Brand.
// @Tags Brand
// @Produce json
// @Success 200 {object} []models.Brand
// @Router /brands [get]
func GetAllBrand(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var names []models.Brand
	db.Find(&names)

	c.JSON(http.StatusOK, gin.H{"data": names})
}

// CreateBrand godoc
// @Summary Create New Brand.
// @Description Creating a new Brand.
// @Tags Brand
// @Param Body body BrandInput true "the body to create a new Brand"
// @Produce json
// @Success 200 {object} models.Brand
// @Router /brands [post]
func CreateBrand(c *gin.Context) {
	// Validate input
	var input BrandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Brand
	Brand := models.Brand{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&Brand)

	c.JSON(http.StatusOK, gin.H{"data": Brand})
}

// GetBrandById godoc
// @Summary Get Brand.
// @Description Get an Brand by id.
// @Tags Brand
// @Produce json
// @Param id path string true "Brand id"
// @Success 200 {object} models.Brand
// @Router /brands/{id} [get]
func GetBrandById(c *gin.Context) { // Get model if exist
	var name models.Brand

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&name).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": name})
}

// GetPhoneByBrandId godoc
// @Summary Get Phones.
// @Description Get all Phones by BrandId.
// @Tags Brand
// @Produce json
// @Param id path string true "Brand id"
// @Success 200 {object} []models.Phone
// @Router /brands/{id}/phones [get]
func GetPhonesByBrandId(c *gin.Context) { // Get model if exist
	var phones []models.Phone

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("brand_id = ?", c.Param("id")).Find(&phones).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": phones})
}

// GetModelByBrandId godoc
// @Summary Get Models.
// @Description Get all Models by BrandId.
// @Tags Brand
// @Produce json
// @Param id path string true "Brand id"
// @Success 200 {object} []models.Model
// @Router /brands/{id}/models [get]
func GetModelsByBrandId(c *gin.Context) { // Get model if exist
	var models []models.Model

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("brand_id = ?", c.Param("id")).Find(&models).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": models})
}

// UpdateBrand godoc
// @Summary Update Brand.
// @Description Update Brand by id.
// @Tags Brand
// @Produce json
// @Param id path string true "Brand id"
// @Param Body body BrandInput true "the body to update age rating category"
// @Success 200 {object} models.Brand
// @Router /brands/{id} [patch]
func UpdateBrand(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var brand models.Brand
	if err := db.Where("id = ?", c.Param("id")).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input BrandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Brand
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&brand).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": brand})
}

// DeleteBrand godoc
// @Summary Delete one Brand.
// @Description Delete a Brand by id.
// @Tags Brand
// @Produce json
// @Param id path string true "Brand id"
// @Success 200 {object} map[string]boolean
// @Router /brands/{id} [delete]
func DeleteBrand(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var brand models.Brand
	if err := db.Where("id = ?", c.Param("id")).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&brand)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
