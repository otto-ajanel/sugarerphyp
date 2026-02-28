package request_dto

type StoreReqUser struct {
	Id        int    `json:"code"`
	StoreName string `json:name`
}

type UserReq struct {
	Email    string `json:email`
	Name     string `json:name`
	LastName string `json:Lastname`
	Password string `json:password`
	Phone    string `json:phone`
	Store    StoreReqUser
	TypeUser StoreReqUser
}

type UserReqUpdate struct {
	UserId   int    `json:user_id`
	Email    string `json:email`
	Name     string `json:name`
	LastName string `json:Lastname`
	Password string `json:password`
	Phone    string `json:phone`
	Store    StoreReqUser
	TypeUser StoreReqUser
}
