package js

func (Null) sealed() {}

func (v Null) Type() Type {
	return TypeNull
}

func (v Null) IsNull() bool {
	return true
}

func (v Null) IsBool() bool {
	return false
}

func (v Null) IsNumber() bool {
	return false
}

func (v Null) IsString() bool {
	return false
}

func (v Null) IsObject() bool {
	return false
}

func (v Null) IsArray() bool {
	return false
}

func (v Null) Bool() bool {
	panic("impossible")
}

func (v Null) Float64() float64 {
	panic("impossible")
}

func (v Null) Int64() int64 {
	panic("impossible")
}

func (v Null) String() string {
	panic("impossible")
}

func (v Null) Object() Object {
	panic("impossible")
}

func (v Null) Array() Array {
	panic("impossible")
}

func (Bool) sealed() {}

func (v Bool) Type() Type {
	return TypeBool
}

func (v Bool) IsNull() bool {
	return false
}

func (v Bool) IsBool() bool {
	return true
}

func (v Bool) IsNumber() bool {
	return false
}

func (v Bool) IsString() bool {
	return false
}

func (v Bool) IsObject() bool {
	return false
}

func (v Bool) IsArray() bool {
	return false
}

func (v Bool) Bool() bool {
	return bool(v)
}

func (v Bool) Float64() float64 {
	panic("impossible")
}

func (v Bool) Int64() int64 {
	panic("impossible")
}

func (v Bool) String() string {
	panic("impossible")
}

func (v Bool) Object() Object {
	panic("impossible")
}

func (v Bool) Array() Array {
	panic("impossible")
}

func (Number) sealed() {}

func (v Number) Type() Type {
	return TypeNumber
}

func (v Number) IsNull() bool {
	return false
}

func (v Number) IsBool() bool {
	return false
}

func (v Number) IsNumber() bool {
	return true
}

func (v Number) IsString() bool {
	return false
}

func (v Number) IsObject() bool {
	return false
}

func (v Number) IsArray() bool {
	return false
}

func (v Number) Bool() bool {
	panic("impossible")
}

func (v Number) Float64() float64 {
	return float64(v)
}

func (v Number) Int64() int64 {
	return int64(v)
}

func (v Number) String() string {
	panic("impossible")
}

func (v Number) Object() Object {
	panic("impossible")
}

func (v Number) Array() Array {
	panic("impossible")
}

func (String) sealed() {}

func (v String) Type() Type {
	return TypeString
}

func (v String) IsNull() bool {
	return false
}

func (v String) IsBool() bool {
	return false
}

func (v String) IsNumber() bool {
	return false
}

func (v String) IsString() bool {
	return true
}

func (v String) IsObject() bool {
	return false
}

func (v String) IsArray() bool {
	return false
}

func (v String) Bool() bool {
	panic("impossible")
}

func (v String) Float64() float64 {
	panic("impossible")
}

func (v String) Int64() int64 {
	panic("impossible")
}

func (v String) String() string {
	return string(v)
}

func (v String) Object() Object {
	panic("impossible")
}

func (v String) Array() Array {
	panic("impossible")
}

func (Object) sealed() {}

func (v Object) Type() Type {
	return TypeObject
}

func (v Object) IsNull() bool {
	return false
}

func (v Object) IsBool() bool {
	return false
}

func (v Object) IsNumber() bool {
	return false
}

func (v Object) IsString() bool {
	return false
}

func (v Object) IsObject() bool {
	return true
}

func (v Object) IsArray() bool {
	return false
}

func (v Object) Bool() bool {
	panic("impossible")
}

func (v Object) Float64() float64 {
	panic("impossible")
}

func (v Object) Int64() int64 {
	panic("impossible")
}

func (v Object) String() string {
	panic("impossible")
}

func (v Object) Object() Object {
	return Object(v)
}

func (v Object) Array() Array {
	panic("impossible")
}

func (Array) sealed() {}

func (v Array) Type() Type {
	return TypeArray
}

func (v Array) IsNull() bool {
	return false
}

func (v Array) IsBool() bool {
	return false
}

func (v Array) IsNumber() bool {
	return false
}

func (v Array) IsString() bool {
	return false
}

func (v Array) IsObject() bool {
	return false
}

func (v Array) IsArray() bool {
	return true
}

func (v Array) Bool() bool {
	panic("impossible")
}

func (v Array) Float64() float64 {
	panic("impossible")
}

func (v Array) Int64() int64 {
	panic("impossible")
}

func (v Array) String() string {
	panic("impossible")
}

func (v Array) Object() Object {
	panic("impossible")
}

func (v Array) Array() Array {
	return Array(v)
}

func (v Object) GetBool(field string) bool {
	return v.Get(field).Bool()
}

func (v Object) GetFloat64(field string) float64 {
	return v.Get(field).Float64()
}

func (v Object) GetInt64(field string) int64 {
	return v.Get(field).Int64()
}

func (v Object) GetString(field string) string {
	return v.Get(field).String()
}

func (v Object) GetObject(field string) Object {
	return v.Get(field).Object()
}

func (v Object) GetArray(field string) Array {
	return v.Get(field).Array()
}

func (v Array) GetBool(i int) bool {
	return v.Get(i).Bool()
}

func (v Array) GetFloat64(i int) float64 {
	return v.Get(i).Float64()
}

func (v Array) GetInt64(i int) int64 {
	return v.Get(i).Int64()
}

func (v Array) GetString(i int) string {
	return v.Get(i).String()
}

func (v Array) GetObject(i int) Object {
	return v.Get(i).Object()
}

func (v Array) GetArray(i int) Array {
	return v.Get(i).Array()
}
