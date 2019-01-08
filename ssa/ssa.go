package ssa

type register struct {
  num int
}

func (v *register) setNum(num int)            { v.num = num }

type Call struct {
    register
    name string
}

func (c *Call) Name() string {
  return c.name
}

type Value interface {
  Name() string
}
