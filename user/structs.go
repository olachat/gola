package user

type Name struct {
	val string
}
type Nick string
type Age int

func (o *Name) SetName(v string) {
	o.val = v
}

func (o *Name) GetName() string {
	return o.val
}

func (o *Age) SetAge(v int) {
	*o = Age(v)
}

func Run() interface{} {
	var q struct {
		Nick
	}
	return q
}
