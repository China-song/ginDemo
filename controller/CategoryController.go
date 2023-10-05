package controller

import (
	"github.com/China-song/ginEssential/model"
	"github.com/China-song/ginEssential/repository"
	"github.com/China-song/ginEssential/response"
	"github.com/China-song/ginEssential/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.ICategoryRepository
}

func NewCategoryController() ICategoryController {
	categoryController := CategoryController{Repository: repository.NewCategoryRepository()}
	categoryController.Repository.(repository.CategoryRepository).DB.AutoMigrate(model.Category{})
	return categoryController
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"category": category}, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 获取body参数
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	// 获取path参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryID)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	// 更新分类名称
	updateCategory, err := c.Repository.Update(*category, requestCategory.Name)
	if err != nil {
		response.Fail(ctx, nil, "更新数据出错")
		return
	}
	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	showCategory, err := c.Repository.SelectById(categoryID)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"category": showCategory}, "查看成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {

	// 获取path参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	_, err := c.Repository.SelectById(categoryID)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	if err := c.Repository.DeleteById(categoryID); err != nil {
		response.Fail(ctx, gin.H{"error": err.Error()}, "删除数据出错")
		return
	}
	response.Success(ctx, nil, "删除成功")
}
