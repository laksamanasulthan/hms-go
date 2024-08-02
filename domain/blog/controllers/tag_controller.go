package controller

import (
	"fmt"
	blog_model "hms-go/domain/blog/models"
	blog_service "hms-go/domain/blog/services"
	blog_validate "hms-go/domain/blog/validation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	serv blog_service.TagService
}

func NewTagController(serv blog_service.TagService) *TagController {
	return &TagController{serv: serv}
}

func (c *TagController) GetAll(ctx *gin.Context) {
	val, err := c.serv.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, val)
}

func (c *TagController) GetOne(ctx *gin.Context) {

	idStr := ctx.Param("id")

	// Convert the ID from string to uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	val, err := c.serv.GetOne(uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, val)
}

func (c *TagController) Create(ctx *gin.Context) {
	var body blog_validate.StoreTagValidation

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate the input
	if err := blog_validate.ValidateStoreProduct(&body); err != nil {
		// Handle validation errors
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTag := &blog_model.Tag{
		Tag: body.Tag,
	}

	fmt.Printf("Created NewTag: %+v\n", newTag)

	val, err := c.serv.Create(newTag)

	fmt.Printf("Created Tag Val: %+v\n", val)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, val)
}
