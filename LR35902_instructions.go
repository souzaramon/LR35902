package LR35902

// LD r8,r8
// Copy (aka Load) the value in register on the right into the register on the left.
func (cpu *CPU) LD_r8_r8(r1 string, r2 string) {
	cpu.SetRegister8(r1, cpu.GetRegister8(r2))
}

// LD r8,n8
// Copy the value n8 into register r8.
func (cpu *CPU) LD_r8_n8(r1 string) {
	cpu.SetRegister8(r1, cpu.Memory.Read8(cpu.Registers.PC))
	cpu.Registers.PC += 1
	cpu.Cycles += 4
}

// LD r16,n16
// Copy the value n16 into register r16.
func (cpu *CPU) LD_r16_n16(r1 string) {
	cpu.SetRegister16(r1, cpu.Memory.Read16(cpu.Registers.PC))
	cpu.Registers.PC += 2
	cpu.Cycles += 8
}

// LD [r16],A
// Copy the value in register A into the byte pointed to by r16.
func (cpu *CPU) LD_r16_A(r1 string) {
	addr := cpu.GetRegister16(r1)

	if r1 == R_C {
		addr |= 0xFF00
	}

	cpu.Memory.Write8(addr, cpu.GetRegister8(R_A))
}

// LD [HL],r8
// Copy the value in register r8 into the byte pointed to by HL.
func (cpu *CPU) LD_HL_r8(r2 string) {
	cpu.Memory.Write8(cpu.GetRegister16(R_HL), cpu.GetRegister8(r2))
}

// LD SP,n16
// Copy the value n16 into register SP.
func (cpu *CPU) LD_SP_n16() {
	cpu.SetRegister16(R_SP, cpu.Memory.Read16(cpu.Registers.PC))
	cpu.Registers.PC += 2
}

// LD A,[r16]
// Copy the byte pointed to by r16 into register A.
func (cpu *CPU) LD_A_r16(r2 string) {
	cpu.SetRegister8(R_A, cpu.Memory.Read8(cpu.GetRegister16(r2)))
}

// LD r8,[HL]
// Copy the value pointed to by HL into register r8.
func (cpu *CPU) LD_r8_HL(r1 string) {
	cpu.SetRegister8(r1, cpu.Memory.Read8(cpu.GetRegister16(R_HL)))
}
