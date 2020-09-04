package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `form:"name" json:"name" binding`
	Password  string `form:"password" json:"password" binding`
	Followers []User `gorm:"many2many:user_followers;"`
	Followees []User `gorm:"many2many:user_followers;joinForeignKey:FollowerId;"`
	Tags      []Tag  `gorm:"many2many:user_tags"`
	// Followers []User `gorm:"foreignkey:user_id;associationforeignkey:follower_id;many2many:user_followers;"`
	// Contacts  []User `gorm:"foreignkey:user_id;associationforeignkey:contact_id;many2many:user_contacts;"`

	IsActive *bool `json:"is_active"`
}

func AddUser(name string, password string) (user User) {
	user = User{
		Name:     name,
		Password: password,
	}
	db.Save(&user)
	return
}

func Follow(user User) bool {
	currentUser := db.First(&User{}, 1)
	var followers []User
	fmt.Println("Get Here")
	fmt.Println(currentUser)
	db.Model(&currentUser).Association("Followers")
	fmt.Println(followers)
	return true
}

func CheckAuth(name, password string) uint {
	var user User
	db.Select("id").Where("name = ? and password = ?", name, password).First(&user)
	fmt.Println(user)
	return user.ID
}
