package model

type User struct {
	ID       int64  `json:"id"`
	Name     string `binding:"required" json:"name"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u *User) UserDetailsWithoutPassword() map[string]interface{} {
	return map[string]interface{}{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
	}

}
