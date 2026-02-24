package model

type UserStore struct {
	UserStoreId int `json:userstore_id gorm:"column:userstore_id;autoIncrement"`
	UserId      int `json:user_id gorm:"column:user_id`
	StoreId     int `json:store_id gorm:"column:store_id`
}

func (UserStore) TableName() string { return "user_store" }
