package hmArticle

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	modelExample "github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hmArticle"
	hmArticleReq "github.com/flipped-aurora/gin-vue-admin/server/model/hmArticle/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
)

type ArticleService struct {
}

// CreateArticle 创建article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) CreateArticle(requestURI string, article *hmArticle.Article) (err error) {
	// 前端创建 && 开启时间验证
	if requestURI == "/viewArticle/createArticle" && global.GVA_CONFIG.Article.OpenVerify {
		// 验证可投稿时间
		err = articleService.verifyTime(global.GVA_CONFIG.Article.SubmissionTime)
		if err != nil {
			return err
		}
	}
	err = global.GVA_DB.Create(article).Error
	return err
}

// 验证有效时间 verifyTimeStr = 2024-04-15 00:00:00,2024-04-21 23:59:59
func (articleService *ArticleService) verifyTime(verifyTimeStr string) (err error) {
	if global.GVA_CONFIG.Article.OpenVerify {
		// 验证稿时间
		verifyTimeArr := strings.Split(verifyTimeStr, ",")
		if len(verifyTimeArr) != 2 {
			return errors.New("时间段设置错误")
		}
		timeStar := utils.TimeStrToTimestamp(verifyTimeArr[0])
		if time.Now().Before(timeStar) {
			return errors.New(fmt.Sprintf("时间未开始 有效时间：%s", verifyTimeStr))
		}
		timeEnd := utils.TimeStrToTimestamp(verifyTimeArr[1])
		if time.Now().After(timeEnd) {
			return errors.New(fmt.Sprintf("时间已结束 有效时间：%s", verifyTimeStr))
		}

	}
	return nil
}

// DeleteArticle 删除article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) DeleteArticle(ID string) (err error) {
	err = global.GVA_DB.Delete(&hmArticle.Article{}, "id = ?", ID).Error
	return err
}

// DeleteArticleByIds 批量删除article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) DeleteArticleByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]hmArticle.Article{}, "id in ?", IDs).Error
	return err
}

// UpdateArticle 更新article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) UpdateArticle(requestURI string, article hmArticle.Article) (err error) {

	var articleOld hmArticle.Article
	global.GVA_DB.Model(&articleOld).Where("id = ?", article.ID).First(&articleOld)
	if articleOld.ID == 0 {
		return errors.New("记录不存在")
	}
	//var f *example.FileUploadAndDownloadService.DeleteFile
	var f *example.FileUploadAndDownloadService

	if articleOld.Thumb != article.Thumb {
		//  删除文件
		var fileParams modelExample.ExaFileUploadAndDownload
		fileParams.Url = articleOld.Thumb
		_ = f.DeleteFile(fileParams)
	}
	// 前端修改 && 开启时间验证
	if articleOld.Content != article.Content && requestURI == "/viewArticle/updateArticle" {
		//  删除文件
		var fileParams modelExample.ExaFileUploadAndDownload
		fileParams.Url = articleOld.Content
		err2 := f.DeleteFile(fileParams)
		if err2 != nil {
			fmt.Println("err2:", err2)
		}
	}
	if requestURI == "/viewArticle/updateArticle" && global.GVA_CONFIG.Article.OpenVerify {

		// 验证可投稿时间
		err = articleService.verifyTime(global.GVA_CONFIG.Article.SubmissionTime)
		if err != nil {
			return err
		}
	}
	err = global.GVA_DB.Updates(&article).Error
	return err
}

// ArticleVote 文件投票
// Author [haima](http://haimait.top)
func (articleService *ArticleService) ArticleVote(c *gin.Context, r hmArticleReq.ArticleVoteParams) (err error) {
	userId := utils.GetUserID(c)

	expiration := time.Duration(3600*24) * time.Second // 失效时长
	// 判断最大投票数 每人每类每天最大投票次数
	if len(r.IDs) > global.GVA_CONFIG.Article.MaxVoteCount {
		return errors.New("投票数超过最大票数")
	}

	// 判断是否开启redis
	if global.GVA_REDIS == nil {
		return err
	}
	key := fmt.Sprintf("vote%s:%d:%d", time.Now().Format("20060102"), r.CategoryId, userId) // `vote20240214:categoryId:userId`

	// 判断可投票的时间段  4.15宣传周：4月15日至4月21日
	if global.GVA_CONFIG.Article.OpenVerify {
		// 验证时间
		voteTime := global.GVA_CONFIG.Article.VoteTime
		voteTimeArr := strings.Split(voteTime, ",")
		if len(voteTimeArr) != 2 {
			return errors.New("投票时间段设置错误")
		}
		timeStar := utils.TimeStrToTimestamp(voteTimeArr[0])
		if time.Now().Before(timeStar) {
			return errors.New(fmt.Sprintf("投票时间未开始 有效时间：%s", voteTime))
		}
		timeEnd := utils.TimeStrToTimestamp(voteTimeArr[1])
		if time.Now().After(timeEnd) {
			return errors.New(fmt.Sprintf("投票时间已结束 有效时间：%s", voteTime))
		}
		//验证是否已投票
		count, err := global.GVA_REDIS.Exists(c, key).Result()
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("此类作品今日已投票数，请明日后再尝试")
		}
		// 增加redis中文章票数
		marshalJ, _ := json.Marshal(r)
		if err := global.GVA_REDIS.Set(c, key, string(marshalJ), expiration).Err(); err != nil {
			return err
		}
		//pipe := global.GVA_REDIS.TxPipeline()
		//pipe.Set(c, key, int64(len(r.IDs)))
		//pipe.Expire(c, key, expiration)
		//_, err = pipe.Exec(c)
		//if err != nil {
		//	return err
		//}
	}

	// 增加mysql中文章票数
	var article hmArticle.Article
	err = global.GVA_DB.Model(&article).Where("id in ?", r.IDs).UpdateColumn("ticket_count", gorm.Expr("ticket_count + ?", 1)).Error
	if err != nil {
		return err
	}

	return nil
}

// GetArticle 根据ID获取article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) GetArticle(ID string) (article hmArticle.Article, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&article).Error
	return
}

// GetArticleInfoList 分页获取article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) GetArticleInfoList(info hmArticleReq.ArticleSearch) (list []hmArticle.Article, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hmArticle.Article{}).Preload("Category")
	var articles []hmArticle.Article
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Author != "" {
		db = db.Where("author = ?", info.Author)
	}
	if info.Company != "" {
		db = db.Where("company = ?", info.Company)
	}
	if info.Phone != "" {
		db = db.Where("phone = ?", info.Phone)
	}
	if info.IsLink != 0 {
		db = db.Where("is_link = ?", info.IsLink)
	}
	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	if info.Publish != 0 {
		db = db.Where("publish = ?", info.Publish)
	}

	if info.Awards != 0 {
		if info.Awards == 6 { //获取奖项 5:无奖项 6:已中奖文章 1:一等奖 2:二等奖 3:三等奖 4:特等奖
			db = db.Where("awards != ?", 5) //中奖
		} else {
			db = db.Where("awards = ?", info.Awards)
		}
	}
	if info.Top != 0 {
		db = db.Where("top = ?", info.Top)
	}
	if info.AuditorId != 0 {
		db = db.Where("auditor_id = ?", info.AuditorId)
	}
	if info.CategoryId != 0 {
		db = db.Where("category_id = ?", info.CategoryId)
	}
	if info.CreatedBy != 0 {
		db = db.Where("created_by = ?", info.CreatedBy)
	}
	if info.UpdatedBy != 0 {
		db = db.Where("updated_by = ?", info.UpdatedBy)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["sort_order"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("id desc")
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&articles).Error
	return articles, total, err
}
