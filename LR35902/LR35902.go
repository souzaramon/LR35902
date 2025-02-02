package LR35902

type MemoryLike interface {
	Read8(address uint16) byte
	Write8(address uint16, value byte)
	Read16(address uint16) uint16
	Write16(address uint16, value uint16)
}

type CPU struct {
	Cycles int
	R 		 Registers
	M      MemoryLike
}

func (cpu *CPU) ExecInstruction(opcode byte) {
	switch opcode {
		case 0x00, 0x40, 0x49, 0x52, 0x5b, 0x64, 0x6d, 0x7f: break
		case 0x01: LD_r16_n16(cpu, R_BC)
		case 0x02: LD_r16_A(cpu, R_BC)
		case 0x06: LD_r8_n8(cpu, R_B)
		case 0x08: LD_n16_SP(cpu, )
		case 0x0a: LD_A_r16(cpu, R_BC)
		case 0x0e: LD_r8_n8(cpu, R_C)
		case 0x11: LD_r16_n16(cpu, R_DE)
		case 0x12: LD_r16_A(cpu, R_DE)
		case 0x16: LD_r8_n8(cpu, R_D)
		case 0x1a: LD_A_r16(cpu, R_DE)
		case 0x1e: LD_r8_n8(cpu, R_E)
		case 0x21: LD_r16_n16(cpu, R_HL)
		case 0x22: LD_HLI_A(cpu, )
		case 0x26: LD_r8_n8(cpu, R_H)
		case 0x2a: LD_A_HLI(cpu, )
		case 0x2e: LD_r8_n8(cpu, R_L)
		case 0x31: LD_SP_n16(cpu, )
		case 0x32: LD_HLD_A(cpu, )
		case 0x36: LD_HL_n8(cpu, )
		case 0x3a: LD_A_HLD(cpu, )
		case 0x3e: LD_r8_n8(cpu, R_A)
		case 0x41: LD_r8_r8(cpu, R_B, R_C)
		case 0x42: LD_r8_r8(cpu, R_B, R_D)
		case 0x43: LD_r8_r8(cpu, R_B, R_E)
		case 0x44: LD_r8_r8(cpu, R_B, R_H)
		case 0x45: LD_r8_r8(cpu, R_B, R_L)
		case 0x46: LD_r8_HL(cpu, R_B)
		case 0x47: LD_r8_r8(cpu, R_B, R_A)
		case 0x48: LD_r8_r8(cpu, R_C, R_B)
		case 0x4a: LD_r8_r8(cpu, R_C, R_D)
		case 0x4b: LD_r8_r8(cpu, R_C, R_E)
		case 0x4c: LD_r8_r8(cpu, R_C, R_H)
		case 0x4d: LD_r8_r8(cpu, R_C, R_L)
		case 0x4e: LD_r8_HL(cpu, R_C)
		case 0x4f: LD_r8_r8(cpu, R_C, R_A)
		case 0x50: LD_r8_r8(cpu, R_D, R_B)
		case 0x51: LD_r8_r8(cpu, R_D, R_C)
		case 0x53: LD_r8_r8(cpu, R_D, R_E)
		case 0x54: LD_r8_r8(cpu, R_D, R_H)
		case 0x55: LD_r8_r8(cpu, R_D, R_L)
		case 0x56: LD_r8_HL(cpu, R_D)
		case 0x57: LD_r8_r8(cpu, R_D, R_A)
		case 0x58: LD_r8_r8(cpu, R_E, R_B)
		case 0x59: LD_r8_r8(cpu, R_E, R_C)
		case 0x5a: LD_r8_r8(cpu, R_E, R_D)
		case 0x5c: LD_r8_r8(cpu, R_E, R_H)
		case 0x5d: LD_r8_r8(cpu, R_E, R_L)
		case 0x5e: LD_r8_HL(cpu, R_E)
		case 0x5f: LD_r8_r8(cpu, R_E, R_A)
		case 0x60: LD_r8_r8(cpu, R_H, R_B)
		case 0x61: LD_r8_r8(cpu, R_H, R_C)
		case 0x62: LD_r8_r8(cpu, R_H, R_D)
		case 0x63: LD_r8_r8(cpu, R_H, R_E)
		case 0x65: LD_r8_r8(cpu, R_H, R_L)
		case 0x66: LD_r8_HL(cpu, R_H)
		case 0x67: LD_r8_r8(cpu, R_H, R_A)
		case 0x68: LD_r8_r8(cpu, R_L, R_B)
		case 0x69: LD_r8_r8(cpu, R_L, R_C)
		case 0x6a: LD_r8_r8(cpu, R_L, R_D)
		case 0x6b: LD_r8_r8(cpu, R_L, R_E)
		case 0x6c: LD_r8_r8(cpu, R_L, R_H)
		case 0x6e: LD_r8_HL(cpu, R_L)
		case 0x6f: LD_r8_r8(cpu, R_L, R_A)
		case 0x70: LD_HL_r8(cpu, R_B)
		case 0x71: LD_HL_r8(cpu, R_C)
		case 0x72: LD_HL_r8(cpu, R_D)
		case 0x73: LD_HL_r8(cpu, R_E)
		case 0x74: LD_HL_r8(cpu, R_H)
		case 0x75: LD_HL_r8(cpu, R_L)
		case 0x77: LD_HL_r8(cpu, R_A)
		case 0x78: LD_r8_r8(cpu, R_A, R_B)
		case 0x79: LD_r8_r8(cpu, R_A, R_C)
		case 0x7a: LD_r8_r8(cpu, R_A, R_D)
		case 0x7b: LD_r8_r8(cpu, R_A, R_E)
		case 0x7c: LD_r8_r8(cpu, R_A, R_H)
		case 0x7d: LD_r8_r8(cpu, R_A, R_L)
		case 0x7e: LD_r8_HL(cpu, R_A)
		case 0x80: ADD_A_r8(cpu, R_B)
		case 0x81: ADD_A_r8(cpu, R_C)
		case 0x82: ADD_A_r8(cpu, R_D)
		case 0x83: ADD_A_r8(cpu, R_E)
		case 0x84: ADD_A_r8(cpu, R_H)
		case 0x85: ADD_A_r8(cpu, R_L)
		case 0x87: ADD_A_r8(cpu, R_A)
		case 0x90: SUB_A_r8(cpu, R_B)
		case 0x91: SUB_A_r8(cpu, R_C)
		case 0x92: SUB_A_r8(cpu, R_D)
		case 0x93: SUB_A_r8(cpu, R_E)
		case 0x94: SUB_A_r8(cpu, R_H)
		case 0x95: SUB_A_r8(cpu, R_L)
		case 0x97: SUB_A_r8(cpu, R_A)
		case 0xa0: AND_A_r8(cpu, R_B)
		case 0xa1: AND_A_r8(cpu, R_C)
		case 0xa2: AND_A_r8(cpu, R_D)
		case 0xa3: AND_A_r8(cpu, R_E)
		case 0xa4: AND_A_r8(cpu, R_H)
		case 0xa5: AND_A_r8(cpu, R_L)
		case 0xa7: AND_A_r8(cpu, R_A)
		case 0xa8: XOR_A_r8(cpu, R_B)
		case 0xa9: XOR_A_r8(cpu, R_C)
		case 0xaa: XOR_A_r8(cpu, R_D)
		case 0xab: XOR_A_r8(cpu, R_E)
		case 0xac: XOR_A_r8(cpu, R_H)
		case 0xad: XOR_A_r8(cpu, R_L)
		case 0xaf: XOR_A_r8(cpu, R_A)
		case 0xb0: OR_A_r8(cpu, R_B)
		case 0xb1: OR_A_r8(cpu, R_C)
		case 0xb2: OR_A_r8(cpu, R_D)
		case 0xb3: OR_A_r8(cpu, R_E)
		case 0xb4: OR_A_r8(cpu, R_H)
		case 0xb5: OR_A_r8(cpu, R_L)
		case 0xb7: OR_A_r8(cpu, R_A)
		case 0xb8: CP_A_r8(cpu, R_B)
		case 0xb9: CP_A_r8(cpu, R_C)
		case 0xba: CP_A_r8(cpu, R_D)
		case 0xbb: CP_A_r8(cpu, R_E)
		case 0xbc: CP_A_r8(cpu, R_H)
		case 0xbd: CP_A_r8(cpu, R_L)
		case 0xbf: CP_A_r8(cpu, R_A)
		case 0xf9: LD_SP_HL(cpu, )
	}
}

func (cpu *CPU) ExecInstructionPrefixed(opcode byte) {
	switch opcode {
		case 0x00: break
	}
}

func (cpu *CPU) Step() int {
	cpu.Cycles = 0

	opcode := cpu.M.Read8(cpu.R.PC)
	cpu.R.PC += 1

	switch (opcode) {
		case 0xcb:
			opcode := cpu.M.Read8(cpu.R.PC)
			cpu.R.PC += 1
			
			cpu.ExecInstructionPrefixed(opcode)
		default:
			cpu.ExecInstruction(opcode)
	}

	return cpu.Cycles
}
