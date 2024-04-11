// 自动生成模板TblKolResourceClean
package pkgCrawlReports

// tblKolResourceClean清洗表 结构体  TblKolResourceClean
type TblKolResourceClean struct {
	Id                int     `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`                                                    //id字段
	Avatar            string  `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;size:255;"`                                                //头像
	AverageViews      *int    `json:"averageViews" form:"averageViews" gorm:"column:average_views;comment:均观看量;size:10;"`                            //均观看量
	CategoryId        *uint8  `json:"categoryId" form:"categoryId" gorm:"type:tinyint;column:category_id;comment:;"`                                 //categoryId字段
	ChannelAccount    string  `json:"channelAccount" form:"channelAccount" gorm:"column:channel_account;comment:频道账号;size:100;"`                     //频道账号
	ChannelId         string  `json:"channelId" form:"channelId" gorm:"column:channel_id;comment:频道ID;size:100;"`                                    //频道ID
	ChannelUrl        string  `json:"channelUrl" form:"channelUrl" gorm:"column:channel_url;comment:频道链接;size:600;"`                                 //频道链接
	CreateTime        *int    `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;size:10;"`                                  //创建时间
	Email             string  `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:300;"`                                                   //邮箱
	Fans              *int    `json:"fans" form:"fans" gorm:"column:fans;comment:粉丝数量;size:10;"`                                                     //粉丝数量
	InteractionRate   *int    `json:"interactionRate" form:"interactionRate" gorm:"column:interaction_rate;comment:互动率;size:10;"`                    //互动率
	Language          string  `json:"language" form:"language" gorm:"column:language;comment:语言;size:100;"`                                          //语言
	LastCrawlTime     *int    `json:"lastCrawlTime" form:"lastCrawlTime" gorm:"column:last_crawl_time;comment:最近抓取时间;size:10;"`                      //最近抓取时间
	LastPublishedTime *int    `json:"lastPublishedTime" form:"lastPublishedTime" gorm:"column:last_published_time;comment:最近发布时间;size:10;"`          //最近发布时间
	Nickname          *string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:昵称;size:100;"`                                          //昵称
	OtherContact      *string `json:"otherContact" form:"otherContact" gorm:"column:other_contact;comment:其它联系方式;size:300;"`                         //其它联系方式
	Platform          uint8   `json:"platform" form:"platform" gorm:"type:tinyint;column:platform;comment:平台：1 - TikTok 2 - YouTube 3 - Instagram;"` //平台：1 - TikTok 2 - YouTube 3 - Instagram
	PolicyId          *uint8  `json:"policyId" form:"policyId" gorm:"type:tinyint;column:policy_id;comment:关联tbl_kol_resource_polices;"`             //关联tbl_kol_resource_polices
	RegionCode        *int    `json:"regionCode" form:"regionCode" gorm:"column:region_code;comment:地区编码;size:10;"`                                  //地区编码
	Signature         *string `json:"signature" form:"signature" gorm:"column:signature;comment:红人简介;"`                                              //红人简介
	Source            *uint8  `json:"source" form:"source" gorm:"type:tinyint;column:source;comment:;"`                                              //source字段
	TotalLikes        *int    `json:"totalLikes" form:"totalLikes" gorm:"column:total_likes;comment:总点赞数;size:10;"`                                  //总点赞数
	TotalVideos       *int    `json:"totalVideos" form:"totalVideos" gorm:"column:total_videos;comment:总视频数;size:10;"`                               //总视频数
	TotalViews        *int    `json:"totalViews" form:"totalViews" gorm:"column:total_views;comment:总观看量;size:19;"`                                  //总观看量
	UpdateTime        *int    `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;size:10;"`                                  //更新时间
	WhichServer       *uint8  `json:"whichServer" form:"whichServer" gorm:"type:tinyint;column:which_server;comment:关联tbl_scrapy_sites 字段;"`         //关联tbl_scrapy_sites 字段
}

// TableName tblKolResourceClean清洗表 TblKolResourceClean自定义表名 tbl_kol_resource_clean
func (TblKolResourceClean) TableName() string {
	return "tbl_kol_resource_clean"
}
