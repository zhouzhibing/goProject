package mysql


type Entity struct {
	Id int `Id`
	Name string
	Argvalue string
}

func (this * Entity ) GetId() int{
	return this.Id
}

func (this * Entity ) GetName() string{
	return this.Name
}

func (this * Entity ) GetArgvalue() string{
	return this.Argvalue
}

func (this * Entity ) SetArgvalue(v string){
	this.Argvalue = v
}
func (this *Entity) Clone() interface{}{
	return new(Entity)
}