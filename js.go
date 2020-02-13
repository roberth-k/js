package js

//go:generate ./generate.sh methods.go

const (
	TypeNull Type = iota + 1
	TypeBool
	TypeNumber
	TypeString
	TypeObject
	TypeArray
)

type Type int

type Value interface {
	sealed()

	Type() Type

	IsNull() bool
	IsBool() bool
	IsString() bool
	IsNumber() bool
	IsObject() bool
	IsArray() bool

	Bool() bool
	String() string
	Int64() int64
	Float64() float64
	Object() Object
	Array() Array
}

type Null struct{}

func (Null) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

type Bool bool

type Number float64

type String string

type Object map[string]Value

func (v Object) Get(field string) Value {
	return v[field]
}

type Array []Value

func (v Array) Get(i int) Value {
	return v[i]
}

var (
	_ Value = new(Null)
	_ Value = new(Bool)
	_ Value = new(Number)
	_ Value = new(String)
	_ Value = new(Object)
	_ Value = new(Array)
)
