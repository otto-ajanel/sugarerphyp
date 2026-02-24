package model

// User representa la tabla users mínima usada para login.
type User struct {
	IDUser   int    `json:"id_user" gorm:"column:id_user;primaryKey"`
	Name     string `json:"name" gorm:"column:name"`
	Lastname string `json:lastname gorm:co"solumn:lastname"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"-" gorm:"column:password"`
	Active   bool   `json:"active" gorm:"column:active"`
	IdTenant int    `json:"id_tenant" gorm:"column:id_tenant"`
}

// TableName for gorm
func (User) TableName() string {
	return "users"
}

// Tenant minimal
type Tenant struct {
	IDTenant   int    `json:"id_tenant" gorm:"column:id_tenant;primaryKey"`
	NameTenant string `json:"name_tenant" gorm:"column:name_tenant"`
}

func (Tenant) TableName() string { return "tenants" }
