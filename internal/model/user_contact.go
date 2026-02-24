package model

type UserContact struct {
	UserContactId     int                    `json:usercontact_id gorm:"column:usercontact_id;autoIncrement"`
	Address           string                 `json:address gorm:"column:address"`
	UserId            int                    `json:user_id gorm:"column:user_id"`
	UserContactPhones map[string]interface{} `json:"usercontact_phones" gorm:"column:usercontact_phones;type:jsonb"`
}

func (UserContact) TableName() string { return "user_contact" }
