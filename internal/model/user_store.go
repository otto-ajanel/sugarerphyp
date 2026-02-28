package model

type UserStore struct {
	UserstoreId int `json:"userstore_id" gorm:"column:userstore_id;primaryKey"`
	UserId      int `json:"user_id" gorm:"column:user_id`
	StoreId     int `json:"store_id" gorm:"column:store_id`
	StateId     int `json:"state_id" gorm:column:state_id`
}

func (UserStore) TableName() string { return "user_store" }
