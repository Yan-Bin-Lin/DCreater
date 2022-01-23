package database

import (
	"time"
)

type User struct {
	Uid        int       `json:"uid" xorm:"not null pk autoincr INT(11) 'uid'"`
	Username   string    `json:"uusername" xorm:"not null comment('account number') VARCHAR(40) 'username'"`
	Password   string    `json:"upassword" xorm:"not null CHAR(40) 'password'"`
	Email      string    `json:"uemail" xorm:"not null VARCHAR(40) 'email'"`
	Createtime time.Time `json:"ucreatetime" xorm:"default 'current_timestamp()' comment('REGIST TIME') DATETIME 'createtime'"`
}

type Owner struct {
	Oid         int       `json:"oid" xorm:"not null pk autoincr INT(11) 'oid'"`
	Uid         int       `json:"ouid" xorm:"not null INT(11) 'uid'"`
	Nickname    string    `json:"onickname" xorm:"not null VARCHAR(50) 'nickname'"`
	Uniquename  string    `json:"ouniquename" xorm:"not null comment('id of User-defined') VARCHAR(50) 'uniquename'"`
	Description string    `json:"odescription" xorm:"comment('abstract') VARCHAR(255) 'description'"`
	Createtime  time.Time `json:"ocreatetime" xorm:"default 'current_timestamp()' DATETIME 'createtime'"`
	Updatetime  time.Time `json:"oupdatetime" xorm:"default 'current_timestamp()' comment('last login time') DATETIME 'updatetime'"`
}

type Blogtype struct {
	Typeid int    `json:"typeid" xorm:"not null pk autoincr TINYINT(4) 'typeid'"`
	Name   string `json:"name" xorm:"not null VARCHAR(20) 'name'"`
}

type Blog struct {
	Bid         int       `json:"bid" xorm:"not null pk autoincr INT(11) 'bid'"`
	Oid         int       `json:"boid" xorm:"not null INT(11) 'oid'"`
	Name        string    `json:"bname" xorm:"not null VARCHAR(100) 'name'"`
	Super       int       `json:"bsuper" xorm:"INT(11) 'super'"`
	Like        int       `json:"blike" xorm:"not null default 0 INT(11) 'like'"`
	Hate        int       `json:"bhate" xorm:"not null default 0 INT(11) 'hate'"`
	Viewtime    int       `json:"bviewtime" xorm:"not null default 0 INT(11) 'viewtime'"`
	Description string    `json:"bdescription" xorm:"comment('abstract') VARCHAR(255) 'description'"`
	Type        int       `json:"btype" xorm:"not null TINYINT(4) 'type'"`
	Urlpath     string    `json:"burlpath" xorm:"not null VARCHAR(750) 'urlpath'"`
	Createtime  time.Time `json:"bcreatetime" xorm:"default 'current_timestamp()' DATETIME 'createtime'"`
	Updatetime  time.Time `json:"bupdatetime" xorm:"default 'current_timestamp()' DATETIME 'updatetime'"`
}

type Article struct {
	Name  string `json:"name" xorm:"not null pk CHAR(44) 'name'"`
	Super int    `json:"super" xorm:"not null INT(11) 'super'"`
}

/* parse to html response */
type UserOut struct {
	Uid      int    `json:"uid" xorm:"not null pk autoincr INT(11) 'uid'"`
	Username string `json:"username" xorm:"not null comment('account number') VARCHAR(40) 'username'"`
	OwnerSub `xorm:"extends"`
}

type OwnerSub struct {
	SubOid         string `json:"subOid" xorm:"VARCHAR(512) 'oid'"`
	SubUniquename  string `json:"subOuniquename" xorm:"VARCHAR(512) 'uniquename'"`
	SubNickname    string `json:"subOnickname" xorm:"VARCHAR(512) 'nickname'"`
	SubDescription string `json:"subOdescription" xorm:"VARCHAR(512) 'description'"`
}

type OwnerOut struct {
	Owner   `xorm:"extends"`
	BlogSub `xorm:"extends"`
}

type BlogList struct {
	OUniquename string `json:"ouniquename" xorm:"not null comment('id of User-defined') VARCHAR(50) 'uniquename'"`
	ONickname   string `json:"onickname" xorm:"not null VARCHAR(50) 'nickname'"`
	Blog        `xorm:"extends"`
}

type BlogProj struct {
	BlogList `xorm:"extends"`
	BlogSub  `xorm:"extends"`
}

type BlogSub struct {
	SubBid         string `json:"subBid" xorm:"VARCHAR(512) 'subBid'"`
	SubName        string `json:"subBname" xorm:"VARCHAR(512) 'subBname'"`
	SubDescription string `json:"subBdescription" xorm:"VARCHAR(512) 'subBdescription'"`
	SubCreatetime  string `json:"subBcreatetime" xorm:"VARCHAR(512) 'subBcreatetime'"`
	SubUpdatetime  string `json:"subBupdatetime" xorm:"VARCHAR(512) 'subBupdatetime'"`
}
