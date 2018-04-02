package entity

type User struct {
	UId string
	Id int64
	Scene_id int
	X int
	Y int
	Robot bool
}

func (this * User) Clone() interface{} {
	return new(User)
}
