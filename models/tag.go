package models

import (
	"baseapi/models/request"
	"fmt"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name  string `json:"name"`
	State *int   `json:"state" binding:"oneof=0 1"`
}

func GetTags(info request.PageInfo, maps interface{}) (err error, list []Tag, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	fmt.Println(db)
	err = db.Model(&Tag{}).Where(maps).Count(&total).Error
	if err != nil {
		return err, nil, total
	} else {
		err = db.Where(maps).Offset(offset).Limit(limit).Find(&list).Error
	}
	return err, list, total
}

func ExistTagByName(name string) bool {
	var count int64
	db.Model(&Tag{}).Where("name = ?", name).Count(&count)
	return count > 0
}

func AddTag(name string, state *int) (tag Tag) {
	tag = Tag{
		Name:  name,
		State: state,
	}
	db.Save(&tag)
	return
}
