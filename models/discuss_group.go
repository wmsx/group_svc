package models

import "gorm.io/gorm"

type DiscussGroup struct {
	Model
	Name   string `gorm:"type:varchar(48);not null"`
	PostId int64  `gorm:"type:bigint;not null"`
	Cover  string `gorm:"type:varchar(256)"`
}

func CreateDiscussGroup(name, cover string, postId int64) error {
	discussGroup := DiscussGroup{
		Name:   name,
		Cover:  cover,
		PostId: postId,
	}
	if err := db.Create(&discussGroup).Error; err != nil {
		return err
	}
	return nil
}

func GetAllDiscussGroups() ([]*DiscussGroup, error) {
	var discussGroups []*DiscussGroup
	err := db.Find(&discussGroups).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return discussGroups, nil
}

func GetDiscussGroupById(id int64) (*DiscussGroup, error) {
	var discussGroup DiscussGroup
	err := db.Where("id = ? and deleted_at is null", id).First(&discussGroup).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &discussGroup, nil
}
