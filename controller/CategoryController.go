package controller

import (
	"errors"
	"github.com/China-song/ginEssential/common"
	"github.com/China-song/ginEssential/model"
	"github.com/China-song/ginEssential/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})
	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		log.Printf("requestCategory.Name: %s", requestCategory.Name)
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	c.DB.Create(&requestCategory)

	response.Success(ctx, gin.H{"category": requestCategory}, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 获取body参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		log.Printf("requestCategory.Name: %s", requestCategory.Name)
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	// 获取path参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory model.Category
	if errors.Is(c.DB.First(&updateCategory, categoryID).Error, gorm.ErrRecordNotFound) {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	// 更新分类名称
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var showCategory model.Category
	if errors.Is(c.DB.First(&showCategory, categoryID).Error, gorm.ErrRecordNotFound) {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"category": showCategory}, "查看成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {

	// 获取path参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var deleteCategory model.Category
	if errors.Is(c.DB.First(&deleteCategory, categoryID).Error, gorm.ErrRecordNotFound) {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	c.DB.Delete(&deleteCategory)
	response.Success(ctx, gin.H{"category": deleteCategory}, "删除成功")
}
