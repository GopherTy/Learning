package interfaces

/*
	Go 语言使用接口，实现可选参数。
	以下示例设置学生的可选信息，电话和住址。
*/

// Options 可选参数接口
type Options interface {
	apply(*studentOptions)
}

type student struct {
	name string
	age  int
	ops  studentOptions
}

func newStudent(name string, age int, ops ...Options) student {
	s := student{
		name: name,
		age:  age,
	}

	for _, o := range ops {
		o.apply(&s.ops)
	}
	return s
}

type adaptOptions struct {
	f func(*studentOptions)
}

func (fdo *adaptOptions) apply(do *studentOptions) {
	fdo.f(do)
}

func newFuncStudentOptions(f func(*studentOptions)) *adaptOptions {
	return &adaptOptions{
		f: f,
	}
}

type studentOptions struct {
	phone   int
	address string
}

func setPhone(phone int) Options {
	return newFuncStudentOptions(func(o *studentOptions) {
		o.phone = phone
	})
}

func setAddress(address string) Options {
	return newFuncStudentOptions(func(o *studentOptions) {
		o.address = address
	})
}
