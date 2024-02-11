package hmArticle

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/hmArticle"
    hmArticleReq "github.com/flipped-aurora/gin-vue-admin/server/model/hmArticle/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ArticleApi struct {
}

var articleService = service.ServiceGroupApp.HmArticleServiceGroup.ArticleService


// CreateArticle 创建article表
// @Tags Article
// @Summary 创建article表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hmArticle.Article true "创建article表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /article/createArticle [post]
func (articleApi *ArticleApi) CreateArticle(c *gin.Context) {
	var article hmArticle.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := articleService.CreateArticle(&article); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteArticle 删除article表
// @Tags Article
// @Summary 删除article表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hmArticle.Article true "删除article表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /article/deleteArticle [delete]
func (articleApi *ArticleApi) DeleteArticle(c *gin.Context) {
	ID := c.Query("ID")
	if err := articleService.DeleteArticle(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteArticleByIds 批量删除article表
// @Tags Article
// @Summary 批量删除article表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /article/deleteArticleByIds [delete]
func (articleApi *ArticleApi) DeleteArticleByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := articleService.DeleteArticleByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateArticle 更新article表
// @Tags Article
// @Summary 更新article表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hmArticle.Article true "更新article表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /article/updateArticle [put]
func (articleApi *ArticleApi) UpdateArticle(c *gin.Context) {
	var article hmArticle.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := articleService.UpdateArticle(article); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindArticle 用id查询article表
// @Tags Article
// @Summary 用id查询article表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hmArticle.Article true "用id查询article表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /article/findArticle [get]
func (articleApi *ArticleApi) FindArticle(c *gin.Context) {
	ID := c.Query("ID")
	if rearticle, err := articleService.GetArticle(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rearticle": rearticle}, c)
	}
}

// GetArticleList 分页获取article表列表
// @Tags Article
// @Summary 分页获取article表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hmArticleReq.ArticleSearch true "分页获取article表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /article/getArticleList [get]
func (articleApi *ArticleApi) GetArticleList(c *gin.Context) {
	var pageInfo hmArticleReq.ArticleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := articleService.GetArticleInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
