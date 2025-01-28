package sm83

import (
	"fmt"
	"os"
)

type CPU struct {
	Cycles int

	CurrentOpcode      byte
	CurrentInstruction Instruction

	Data         uint16
	DestData     uint16
	DestIsMemory bool

	Registers Registers

	Memory interface {
		Read8(address uint16) byte
		Write8(address uint16, value byte)
		Read16(address uint16) uint16
		Write16(address uint16, value uint16)
	}
}

type Registers struct {
	A  byte
	F  byte
	B  byte
	C  byte
	D  byte
	E  byte
	H  byte
	L  byte
	PC uint16
	SP uint16
}

func (cpu *CPU) Print(message string, a ...any) {
	fmt.Printf("[PC 0x%04X] ", cpu.Registers.PC)
	fmt.Printf(message, a...)
	fmt.Printf("\n")
}

func (cpu *CPU) PrintAndDie(message string, a ...any) {
	cpu.Print(message, a...)
	os.Exit(1)
}

func (cpu *CPU) ReadRegister(rk RegisterKind) uint16 {
	switch rk {
	case RK_A:
		return uint16(cpu.Registers.A)
	case RK_F:
		return uint16(cpu.Registers.F)
	case RK_B:
		return uint16(cpu.Registers.B)
	case RK_C:
		return uint16(cpu.Registers.C)
	case RK_D:
		return uint16(cpu.Registers.D)
	case RK_E:
		return uint16(cpu.Registers.E)
	case RK_H:
		return uint16(cpu.Registers.H)
	case RK_L:
		return uint16(cpu.Registers.L)
	case RK_SP:
		return cpu.Registers.SP
	case RK_PC:
		return cpu.Registers.PC
	case RK_AF:
		return (uint16(cpu.Registers.A) << 8) | uint16(cpu.Registers.F)
	case RK_BC:
		return (uint16(cpu.Registers.B) << 8) | uint16(cpu.Registers.C)
	case RK_DE:
		return (uint16(cpu.Registers.D) << 8) | uint16(cpu.Registers.E)
	case RK_HL:
		return (uint16(cpu.Registers.H) << 8) | uint16(cpu.Registers.L)
	default:
		cpu.PrintAndDie("Unknown register (%s)", rk)
		return 0
	}
}

func (cpu *CPU) SetRegister(rk RegisterKind, value uint16) {
	switch rk {
	case RK_A:
		cpu.Registers.A = uint8(value & 0xFF)
		return
	case RK_F:
		cpu.Registers.F = uint8(value & 0xFF)
		return
	case RK_B:
		cpu.Registers.B = uint8(value & 0xFF)
		return
	case RK_C:
		cpu.Registers.C = uint8(value & 0xFF)
		return
	case RK_D:
		cpu.Registers.D = uint8(value & 0xFF)
		return
	case RK_E:
		cpu.Registers.E = uint8(value & 0xFF)
		return
	case RK_H:
		cpu.Registers.H = uint8(value & 0xFF)
		return
	case RK_L:
		cpu.Registers.L = uint8(value & 0xFF)
		return
	case RK_SP:
		cpu.Registers.SP = value
		return
	case RK_PC:
		cpu.Registers.PC = value
		return
	case RK_AF:
		cpu.Registers.A = uint8(value)
		cpu.Registers.F = uint8(value >> 8)
		return
	case RK_BC:
		cpu.Registers.B = uint8(value)
		cpu.Registers.C = uint8(value >> 8)
		return
	case RK_DE:
		cpu.Registers.D = uint8(value)
		cpu.Registers.E = uint8(value >> 8)
		return
	case RK_HL:
		cpu.Registers.H = uint8(value)
		cpu.Registers.L = uint8(value >> 8)
		return
	default:
		cpu.PrintAndDie("Unknown register (%s)", rk)
		return
	}
}

// bits: {z, n, h, c}
func (cpu *CPU) WriteFlags(bits [4]int) {
	for i := 3; i >= 0; i-- {
		bit := bits[i]

		if bit != -1 {
			cpu.Registers.F = SetNthBit(cpu.Registers.F, 7-i, bit > 0)
		}

	}
}

func (cpu *CPU) CheckCondition() bool {
	flag_z := GetNthBit(cpu.Registers.F, 7)
	flag_c := GetNthBit(cpu.Registers.F, 4)

	switch cpu.CurrentInstruction.CK {
	case CK_NONE:
		return true
	case CK_C:
		return flag_c
	case CK_NC:
		return !flag_c
	case CK_Z:
		return flag_z
	case CK_NZ:
		return !flag_z
	default:
		return false
	}
}

func (cpu *CPU) FetchInstruction() {
	currentOpcode := cpu.Memory.Read8(cpu.Registers.PC)
	currentInstruction, exists := InstructionMap[currentOpcode]

	if !exists {
		cpu.PrintAndDie("unknown instruction (0x%02X)", currentOpcode)
	} else {
		cpu.Print(
			"opcode (0x%02X), kind (%s), am (%s)",
			currentOpcode,
			currentInstruction.IK,
			currentInstruction.AM,
		)
	}

	cpu.CurrentOpcode = currentOpcode
	cpu.CurrentInstruction = currentInstruction
}

func (cpu *CPU) FetchData() {
	switch cpu.CurrentInstruction.AM {
	case AM_IMP:
		return
	case AM_R:
		cpu.Data = cpu.ReadRegister(cpu.CurrentInstruction.R1)
		return
	case AM_R_R:
		cpu.Data = cpu.ReadRegister(cpu.CurrentInstruction.R2)
		return
	case AM_R_D8:
		cpu.Data = uint16(cpu.Memory.Read8(cpu.Registers.PC))
		cpu.Cycles += 4
		cpu.Registers.PC += 1
		return
	case AM_R_D16, AM_D16:
		lo := uint16(cpu.Memory.Read8(cpu.Registers.PC))
		hi := uint16(cpu.Memory.Read8(cpu.Registers.PC + 1))

		cpu.Data = lo | (hi << 8)
		cpu.Registers.PC += 2
		cpu.Cycles += 8
		return
	case AM_MR_R:
		cpu.Data = cpu.ReadRegister(cpu.CurrentInstruction.R2)
		cpu.DestData = cpu.ReadRegister(cpu.CurrentInstruction.R1)
		cpu.DestIsMemory = true

		if cpu.CurrentInstruction.R1 == RK_C {
			cpu.DestData |= 0xFF00
		}
		return
	case AM_R_MR:
		address := cpu.ReadRegister(cpu.CurrentInstruction.R2)

		if cpu.CurrentInstruction.R1 == RK_C {
			address |= 0xFF00
		}

		cpu.Data = uint16(cpu.Memory.Read8(address))
		cpu.Cycles += 4
		return
	case AM_R_HLI:
		cpu.Data = cpu.ReadRegister(cpu.CurrentInstruction.R2)
		cpu.SetRegister(RK_HL, cpu.ReadRegister(RK_HL)+1)
		cpu.Cycles += 4
		return
	case AM_R_HLD:
		cpu.Data = cpu.ReadRegister(cpu.CurrentInstruction.R2)
		cpu.SetRegister(RK_HL, cpu.ReadRegister(RK_HL)-1)
		cpu.Cycles += 4
		return
	case AM_HLI_R:
		// TODO
	case AM_HLD_R:
		// TODO
	case AM_R_A8:
		// TODO
	case AM_A8_R:
		// TODO
	case AM_HL_SPR:
		// TODO
	case AM_D8:
		// TODO
	case AM_A16_R:
		// TODO
	case AM_D16_R:
		// TODO
	case AM_MR_D8:
		// TODO
	case AM_MR:
		// TODO
	case AM_R_A16:
		// TODO
	default:
		cpu.PrintAndDie("unknown addressing mode (%s)", cpu.CurrentInstruction.AM)
	}
}

func (cpu *CPU) Execute() {
	switch cpu.CurrentInstruction.IK {
	case IK_NOP:
		cpu.Registers.PC++
		cpu.Cycles += 4
		return
	case IK_JP:
		if cpu.CheckCondition() {
			cpu.Registers.PC = cpu.Data
			cpu.Cycles += 4
			return
		}
		cpu.Registers.PC++
		return
	default:
		cpu.PrintAndDie("instruction kind (%s) not implemented", cpu.CurrentInstruction.IK)
	}
}

func (cpu *CPU) Step() int {
	cpu.Cycles = 0

	cpu.FetchInstruction()
	cpu.FetchData()
	cpu.Execute()

	return cpu.Cycles
}
