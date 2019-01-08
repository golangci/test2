package ssa

// numberRegisters assigns numbers to all SSA registers
// (value-defining Instructions) in f, to aid debugging.
// (Non-Instruction Values are named at construction.)
//
func NumberRegisters(instr interface{}, v int) {
    switch instr.(type) {
    case Value:
      instr.(interface {
        setNum(int)
      }).setNum(v)
    }
}
