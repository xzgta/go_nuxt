package structs

type User struct {
	Id       string `form:"id_user" json:"id_user"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Token    string `form:"token" json:"token"`
}

type Note struct {
	Id    string `form:"id_note" json:"id_note"`
	Title string `form:"title" json:"title"`
	Note  string `form:"note" json:"note"`
}
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
