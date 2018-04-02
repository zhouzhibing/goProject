package entity

type User struct {
	Id int64
	UId string
	Scene_id int
	X int
	Y int
}

func (this * User) Clone() interface{} {
	return new(User)
}
