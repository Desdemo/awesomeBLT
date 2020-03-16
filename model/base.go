package model

import (
	"errors"
	"time"
)

type Base struct {
	Id       int64     `xorm:"pk autoincr comment('唯一ID')" json:"id"`
	CreateAt time.Time `xorm:"created comment('创建时间')"`
	UpdateAt time.Time `xorm:"updated comment('更新时间')"`
	Version  int64     `json:"version" xorm:"comment('版本')"`
}

func (b Base) GetID() int64 {
	return b.Id
}

func (b Base) GetVersion() int64 {
	return b.Version
}

func (b Base) Get(id int64) interface{} {
	has, err := Engine.Id(id).Get(b)
	if err != nil || !has {
		return nil
	} else {
		return b
	}
}

func (b Base) VersionComparison(data interface{}) bool {
	if data == nil {
		return false
	} else {
		input := data.(Base)
		if output := b.Get(input.Id); output == nil {
			return false
		} else {
			if output.(Base).Version != input.Version {
				return false
			} else {
				return true
			}
		}
	}
}

func (b Base) IsExists(id int64) bool {
	has, err := Engine.Id(id).Exist(b)
	if !has || err != nil {
		return false
	} else {
		return true
	}
}

func (b Base) Inserted(data interface{}) error {
	affected, err := Engine.Insert(b)
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("插入数据失败")
	} else {
		return nil
	}
}

func (b Base) Deleted(id int64) bool {

}
