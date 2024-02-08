package model

//type Article struct {
//	global.GVA_MODEL
//	Title         string `json:"title" gorm:"type:varchar(255);index;not null;comment:标题"`                                                               //
//	SubTitle      string `json:"subtitle" gorm:"type:varchar(255);not null;default'';comment:副标题"`                                                       //
//	Keywords      string `json:"keywords" gorm:"type:varchar(255);not null;default'';comment:关键字"`                                                       //
//	Description   string `json:"description" gorm:"type:varchar(1000);default:'';comment:描述"`                                                            //
//	Author        string `json:"author" gorm:"type:varchar(255);not null;default:'';comment:文章作者"`                                                       //
//	Company       string `json:"company" gorm:"type:varchar(255);not null;default:'';comment:参赛公司"`                                                      //
//	Phone         string `json:"phone" gorm:"type:varchar(11);not null;default:'';comment:手机号"`                                                          //
//	Source        string `json:"source" gorm:"type:varchar(255);not null;default:'';comment:来源"`                                                         //
//	Content       string `json:"content" gorm:"type:text;not null;comment:文章内容"`                                                                         //
//	Extend        string `json:"extend"  gorm:"type:varchar(1000);not null;default:'';comment:扩展字段;"`                                                    //扩展字段
//	Thumb         string `json:"thumb" form:"thumb" gorm:"type:varchar(255);not null;default:https://qmplusimg.henrongyi.top/gvalogo.png;comment:封面地址;"` //图片地址
//	IsLink        *bool  `json:"isLink"  gorm:"type:int(1);column:is_link;not null; default:2;comment:是否外链 1:否 1:是;size:1;"`                             //是否显示
//	Link          string `json:"link" gorm:"type:varchar(255);not null;default:'';comment:来源"`                                                           //
//	Type          string `json:"type" form:"type" gorm:"type:enum('article','page','video');comment:类型：article,page,video;"`                             //类型
//	SortOrder     uint   `json:"sort_order" form:"sort_order" gorm:"type:int(4);default:0;comment:排序;"`                                                  //排序
//	Status        uint   `json:"status" form:"status" gorm:"type:int(1);default:1;comment:状态 1:待审核 2:审核通过;"`                                             //排序
//	Publish       int    `json:"publish" form:"publish" gorm:"type:int(1);default:2;comment:前台显示 1:是 2:否"`                                               //排序
//	ViewCount     int64  `json:"view_count" gorm:"type:int(11);not null;default:0;comment:展示次数"`                                                         //
//	TickerCount   int64  `json:"tickerCount" gorm:"type:int(11);not null;default:0;comment:投票次数"`                                                        //
//	Awards        int    `json:"awards" form:"awards" gorm:"type:int(1);default:5;comment:获取奖项 5:无奖项 1:一等奖 2:二等奖 3:三等奖 4:特等奖"`                           //排序
//	Top           *bool  `json:"top"  gorm:"type:int(1);not null; default:2;comment:是否外链 1:是; 2:否 "`                                                     //是否显示
//	AuditorId     int    `json:"auditor_id" gorm:"type:int(11);not null;default:0;comment:审核人"`                                                          //
//	AuditorFailed string `json:"auditor_failed" gorm:"type:varchar(1000);not null;comment:审核失败原因"`                                                       //
//}
//
//func (Article) TableName() string {
//	return "article"
//}

//CREATE TABLE `article` (
//`id` int(11) NOT NULL AUTO_INCREMENT,
//`created_at` datetime NOT NULL COMMENT '创建时间',
//`updated_at` datetime NOT NULL COMMENT '更新时间',
//`title` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '标题',
//`subtitle` varchar(150) COLLATE utf8_bin DEFAULT '' COMMENT '副标题',
//`keywords` varchar(150) COLLATE utf8_bin DEFAULT '' COMMENT '关键字',
//`description` varchar(255) COLLATE utf8_bin DEFAULT '' COMMENT '文章描述',
//`author` varchar(60) COLLATE utf8_bin DEFAULT NULL COMMENT '文章作者',
//`company` varchar(60) COLLATE utf8_bin DEFAULT NULL COMMENT '参赛单位',
//`phone` varchar(11) COLLATE utf8_bin NOT NULL COMMENT '手机号',
//`source` varchar(50) COLLATE utf8_bin DEFAULT NULL COMMENT '文章来源',
//`content` text COLLATE utf8_bin NOT NULL COMMENT '文章内容',
//`attribute` json DEFAULT NULL COMMENT '附加属性',
//`thumb` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '封面',
//`is_link` enum('1', '2') COLLATE utf8_bin NOT NULL DEFAULT '1' COMMENT 'isLink',
//`link` varchar(255) COLLATE utf8_bin DEFAULT '' COMMENT 'Link',
//`type` enum('article', 'page', 'video') COLLATE utf8_bin DEFAULT 'article' COMMENT '类型：article/page/video',
//`order` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
//`status` int(1) DEFAULT '1' COMMENT '状态 1:待审核 2:审核通过',
//`publish` int(1) DEFAULT '2' COMMENT '前台显示 1:是 2:否',
//`view_count` int(11) DEFAULT '0' COMMENT '展示次数',
//`ticket_count` int(11) DEFAULT '0' COMMENT '投票次数',
//`awards` int(1) DEFAULT '0' COMMENT '获取奖项 5:无奖项 1:一等奖 2:二等奖 3:三等奖 4:特等奖',
//`top` enum('1', '2') COLLATE utf8_bin NOT NULL DEFAULT '2' COMMENT '置顶 1:是 2:否',
//`auditor_id` int(11) NOT NULL DEFAULT '0' COMMENT '审核人',
//`auditor_result` varchar(255) COLLATE utf8_bin DEFAULT '' COMMENT '审核结果',
//PRIMARY KEY (`id`) USING BTREE
//) ENGINE = InnoDB AUTO_INCREMENT =131 DEFAULT CHARSET = utf8 COLLATE = utf8_bin ROW_FORMAT =DYNAMIC
