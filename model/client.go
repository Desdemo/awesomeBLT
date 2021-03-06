package model

import "time"

type Client struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	Org             int64     `xorm:"comment('组织编码')"`
	OrgName         string    `json:"org_name"`
	CreatedOn       time.Time `json:"created_on" xorm:"comment('创建时间')"`
	Code            string    `json:"code" xorm:"comment('客户编码')"`
	City            string    `json:"city" xorm:"comment('城市')"`
	Province        string    `json:"province" xorm:"comment('省份')"`
	AccountPeriod   string    `json:"account_period" xorm:"varchar(25) comment('账期')"` //账期
	AccountDate     int       `json:"account_date" xorm:"comment('账期日期')"`
	Merchandiser    string    `json:"merchandiser" xorm:"comment('跟单员')"` // 跟单员
	DepartmentName  string    `json:"department_name"`                    //部门
	Salesman        string    `json:"salesman" xorm:"comment('业务员')"`
	BigType         string    `json:"big_type" xorm:"comment('大分类')"`
	Type            string    `json:"type" xorm:"comment('细分类别')"`
	CooperationTime time.Time `json:"cooperation_time" xorm:"comment('合作时间')"`
	Version         int64     `json:"version" xorm:"comment('版本')"`
}

func (client Client) GetID() int64 {
	return client.Id
}

func (client Client) Get(id int64) interface{} {
	has, err := Engine.ID(id).Get(client)
	if err != nil {
		return err
	} else if !has {
		return nil
	} else {
		return client
	}
}

func (client Client) GetVersion(id int64) int64 {
	var version int64
	err := Engine.Table("client")
}

func (client Client) VersionComparison(data interface{}) bool {
	indata := data.(Client)
	client := client.Get(indata.Id)

	return false
}

func (client Client) IsExists(id int64) bool {
	return false
}
