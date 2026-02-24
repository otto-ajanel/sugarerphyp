package resquest_dto

type StoreReqUser struct {
	Id        int    `json:code`
	StoreName string `json:label`
}

type UserReq struct {
	Email    string `json:email`
	Name     string `json:name`
	LastName string `json:lastname`
	Password string `json:password`
	Phone    string `json:phone`
	Store    StoreReqUser
}
