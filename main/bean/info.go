package bean

type Info struct {
	Name string `json:"name" form:"name"`
	Id   string `json:"id"   form:"id"`
	Sex  string `json:"sex"  form:"sex"`
	Age  int    `json:"age"  form:"age"`
}
