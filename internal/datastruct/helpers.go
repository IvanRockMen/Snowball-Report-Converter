package datastruct

type Bool struct {
	Exists bool
	Value  bool
}

func NewBool(value bool) *Bool {
	return &Bool{
		Exists: true,
		Value:  value,
	}
}
