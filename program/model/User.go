package model

type User struct {
	Id           int    `gorm:"column:id;primary_key;auto_increment;"`
	NickName     string `gorm:"column:nickname;unique;type:varchar(100);index:name;not null"`
	Avatar       string `gorm:"column:avatar;type:varchar(255);not null"`
	password     string `gorm:"column:password;type:varchar(255);not null"`
	Phone        string `gorm:"column:phone;type:char(11);unique;index:phone;not null"`
	RegisterTime string `gorm:"column:register_time;type:datetime;not null"`
	Level        int    `gorm:"column:level;type:smallint;not null;default:1"`
	Score int 			`gorm:"column:score;not null;default:1"`
	GameCard int `gorm:column:game_card;not null;default:10`
}
