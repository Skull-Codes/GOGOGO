package forms

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordUser struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
