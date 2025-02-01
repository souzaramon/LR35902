package sm83

var InstructionMap = map[byte]Instruction{
	0x00: {IK: IK_NOP, AM: AM_IMP},
	0x01: {IK: IK_LD, AM: AM_R_D16, R1: RK_BC},
	0x02: {IK: IK_LD, AM: AM_MR_R, R1: RK_BC, R2: RK_A},
	0x06: {IK: IK_LD, AM: AM_R_D8, R1: RK_B},
	0x08: {IK: IK_LD, AM: AM_D16_R, R1: RK_NONE, R2: RK_SP},
	0x0a: {IK: IK_LD, AM: AM_R_MR, R1: RK_A, R2: RK_BC},
	0x0e: {IK: IK_LD, AM: AM_R_D8, R1: RK_C},
	0x11: {IK: IK_LD, AM: AM_R_D16, R1: RK_DE},
	0x12: {IK: IK_LD, AM: AM_MR_R, R1: RK_DE, R2: RK_A},
	0x16: {IK: IK_LD, AM: AM_R_D8, R1: RK_D},
	0x1a: {IK: IK_LD, AM: AM_R_MR, R1: RK_A, R2: RK_DE},
	0x1e: {IK: IK_LD, AM: AM_R_D8, R1: RK_E},
	0x21: {IK: IK_LD, AM: AM_R_D16, R1: RK_HL},
	0x22: {IK: IK_LD, AM: AM_HLI_R, R1: RK_HL, R2: RK_A},
	0x26: {IK: IK_LD, AM: AM_R_D8, R1: RK_H},
	0x2a: {IK: IK_LD, AM: AM_R_HLI, R1: RK_A, R2: RK_HL},
	0x2e: {IK: IK_LD, AM: AM_R_D8, R1: RK_L},
	0x31: {IK: IK_LD, AM: AM_R_D16, R1: RK_SP},
	0x32: {IK: IK_LD, AM: AM_HLD_R, R1: RK_HL, R2: RK_A},
	0x36: {IK: IK_LD, AM: AM_MR_D8, R1: RK_HL},
	0x3a: {IK: IK_LD, AM: AM_R_HLD, R1: RK_A, R2: RK_HL},
	0x3e: {IK: IK_LD, AM: AM_R_D8, R1: RK_A},
	0x40: {IK: IK_NOP, AM: AM_IMP},
	0x41: {IK: IK_LD, AM: AM_R_R, R1: RK_B, R2: RK_C},
	0x42: {IK: IK_LD, AM: AM_R_R, R1: RK_B, R2: RK_D},
	0x43: {IK: IK_LD, AM: AM_R_R, R1: RK_B, R2: RK_E},
	0x44: {IK: IK_LD, AM: AM_R_R, R1: RK_B, R2: RK_H},
	0x45: {IK: IK_LD, AM: AM_R_R, R1: RK_B, R2: RK_L},
	0x46: {IK: IK_LD, AM: AM_R_MR, R1: RK_B, R2: RK_HL},
	0x47: {IK: IK_LD, AM: AM_R_R, R1: RK_B, R2: RK_A},
	0x48: {IK: IK_LD, AM: AM_R_R, R1: RK_C, R2: RK_B},
	0x49: {IK: IK_NOP, AM: AM_IMP},
	0x4a: {IK: IK_LD, AM: AM_R_R, R1: RK_C, R2: RK_D},
	0x4b: {IK: IK_LD, AM: AM_R_R, R1: RK_C, R2: RK_E},
	0x4c: {IK: IK_LD, AM: AM_R_R, R1: RK_C, R2: RK_H},
	0x4d: {IK: IK_LD, AM: AM_R_R, R1: RK_C, R2: RK_L},
	0x4e: {IK: IK_LD, AM: AM_R_MR, R1: RK_C, R2: RK_HL},
	0x4f: {IK: IK_LD, AM: AM_R_R, R1: RK_C, R2: RK_A},
	0x50: {IK: IK_LD, AM: AM_R_R, R1: RK_D, R2: RK_B},
	0x51: {IK: IK_LD, AM: AM_R_R, R1: RK_D, R2: RK_C},
	0x52: {IK: IK_NOP, AM: AM_IMP},
	0x53: {IK: IK_LD, AM: AM_R_R, R1: RK_D, R2: RK_E},
	0x54: {IK: IK_LD, AM: AM_R_R, R1: RK_D, R2: RK_H},
	0x55: {IK: IK_LD, AM: AM_R_R, R1: RK_D, R2: RK_L},
	0x56: {IK: IK_LD, AM: AM_R_MR, R1: RK_D, R2: RK_HL},
	0x57: {IK: IK_LD, AM: AM_R_R, R1: RK_D, R2: RK_A},
	0x58: {IK: IK_LD, AM: AM_R_R, R1: RK_E, R2: RK_B},
	0x59: {IK: IK_LD, AM: AM_R_R, R1: RK_E, R2: RK_C},
	0x5a: {IK: IK_LD, AM: AM_R_R, R1: RK_E, R2: RK_D},
	0x5b: {IK: IK_NOP, AM: AM_IMP},
	0x5c: {IK: IK_LD, AM: AM_R_R, R1: RK_E, R2: RK_H},
	0x5d: {IK: IK_LD, AM: AM_R_R, R1: RK_E, R2: RK_L},
	0x5e: {IK: IK_LD, AM: AM_R_MR, R1: RK_E, R2: RK_HL},
	0x5f: {IK: IK_LD, AM: AM_R_R, R1: RK_E, R2: RK_A},
	0x60: {IK: IK_LD, AM: AM_R_R, R1: RK_H, R2: RK_B},
	0x61: {IK: IK_LD, AM: AM_R_R, R1: RK_H, R2: RK_C},
	0x62: {IK: IK_LD, AM: AM_R_R, R1: RK_H, R2: RK_D},
	0x63: {IK: IK_LD, AM: AM_R_R, R1: RK_H, R2: RK_E},
	0x64: {IK: IK_NOP, AM: AM_IMP},
	0x65: {IK: IK_LD, AM: AM_R_R, R1: RK_H, R2: RK_L},
	0x66: {IK: IK_LD, AM: AM_R_MR, R1: RK_H, R2: RK_HL},
	0x67: {IK: IK_LD, AM: AM_R_R, R1: RK_H, R2: RK_A},
	0x68: {IK: IK_LD, AM: AM_R_R, R1: RK_L, R2: RK_B},
	0x69: {IK: IK_LD, AM: AM_R_R, R1: RK_L, R2: RK_C},
	0x6a: {IK: IK_LD, AM: AM_R_R, R1: RK_L, R2: RK_D},
	0x6b: {IK: IK_LD, AM: AM_R_R, R1: RK_L, R2: RK_E},
	0x6c: {IK: IK_LD, AM: AM_R_R, R1: RK_L, R2: RK_H},
	0x6d: {IK: IK_NOP, AM: AM_IMP},
	0x6e: {IK: IK_LD, AM: AM_R_MR, R1: RK_L, R2: RK_HL},
	0x6f: {IK: IK_LD, AM: AM_R_R, R1: RK_L, R2: RK_A},
	0x70: {IK: IK_LD, AM: AM_MR_R, R1: RK_HL, R2: RK_B},
	0x71: {IK: IK_LD, AM: AM_MR_R, R1: RK_HL, R2: RK_C},
	0x72: {IK: IK_LD, AM: AM_MR_R, R1: RK_HL, R2: RK_D},
	0x73: {IK: IK_LD, AM: AM_MR_R, R1: RK_HL, R2: RK_E},
	0x74: {IK: IK_LD, AM: AM_MR_R, R1: RK_HL, R2: RK_H},
	0x75: {IK: IK_LD, AM: AM_MR_R, R1: RK_HL, R2: RK_L},
	0x77: {IK: IK_LD, AM: AM_MR_R, R1: RK_HL, R2: RK_A},
	0x78: {IK: IK_LD, AM: AM_R_R, R1: RK_A, R2: RK_B},
	0x79: {IK: IK_LD, AM: AM_R_R, R1: RK_A, R2: RK_C},
	0x7a: {IK: IK_LD, AM: AM_R_R, R1: RK_A, R2: RK_D},
	0x7b: {IK: IK_LD, AM: AM_R_R, R1: RK_A, R2: RK_E},
	0x7c: {IK: IK_LD, AM: AM_R_R, R1: RK_A, R2: RK_H},
	0x7d: {IK: IK_LD, AM: AM_R_R, R1: RK_A, R2: RK_L},
	0x7e: {IK: IK_LD, AM: AM_R_MR, R1: RK_A, R2: RK_HL},
	0x7f: {IK: IK_NOP, AM: AM_IMP},
	0xC2: {IK: IK_JP, AM: AM_D16, CK: CK_NZ},
	0xC3: {IK: IK_JP, AM: AM_D16, CK: CK_NONE},
	0xCA: {IK: IK_JP, AM: AM_D16, CK: CK_Z},
	0xD2: {IK: IK_JP, AM: AM_D16, CK: CK_NC},
	0xDA: {IK: IK_JP, AM: AM_D16, CK: CK_C},
	0xea: {IK: IK_LD, AM: AM_D16_R, R1: RK_NONE, R2: RK_A},
	0xE9: {IK: IK_JP, AM: AM_R, R1: RK_HL, CK: CK_NONE},
	0xf8: {IK: IK_LD, AM: AM_HL_SPD8},
	0xf9: {IK: IK_LD, AM: AM_R_R, R1: RK_SP, R2: RK_HL},
	0xfa: {IK: IK_LD, AM: AM_R_A16, R1: RK_A},
}
