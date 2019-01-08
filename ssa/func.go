package ssa

func Panics(instr interface{}, v int) {
    switch instr.(type) {
    case Value:
      instr.(interface {
        setNum(int)
      }).setNum(v)
    }
}

type setNumable interface {
  setNum(int)
}

func DoesNotPanic(instr interface{}, v int) {
    switch instr.(type) {
    case Value:
      instr.(setNumable).setNum(v)
    }
}
