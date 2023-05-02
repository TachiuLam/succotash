package base

import "testing"

type p struct {
	Name string
	age  int
}

func NewP() *p {
	return &p{
		Name: "huahua",
		age:  4,
	}
}

func TestP(t *testing.T) {
	p1 := NewP()
	p2 := NewP()
	p3 := p1
	p1.age = 2
	t.Log("p1:", p1)
	t.Log("p2:", p2)
	t.Log("p3:", p3)
}
