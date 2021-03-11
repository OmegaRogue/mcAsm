/*
 *  mcAsm - a minecraft Assembler
 *     Copyright (C) 2021 OmegaRogue
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */

package app

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	scanner      Scanner
	symbols      map[string]int
	instructions []Instruction
	nAddr        int // next available address
}

func (p *Parser) Init(src []byte) {
	p.scanner.Init(src)
	p.nAddr = 16 // address [0:15] are reserved
	p.symbols = map[string]int{
		"R0": 0, "R1": 1, "R2": 2, "R3": 3, "R4": 4, "R5": 5, "R6": 6, "R7": 7, "R8": 8,
		"R9": 9, "R10": 10, "R11": 11, "R12": 12, "R13": 13, "R14": 14, "R15": 15,
		"SCREEN": 16384, "KBD": 24576,
		"SP": 0, "LCL": 1, "ARG": 2, "THIS": 3, "THAT": 4,
	}
}

func (p *Parser) Parse() HackFile {
loop:
	for {
		tok, lit := p.scanner.Scan()
		switch tok {
		case EOF:
			break loop // break out of the loop not just the switch
		case LABEL:
			p.symbols[lit] = len(p.instructions)
		case A_INSTRUCTION:
			p.instructions = append(p.instructions, &AInstruction{lit: lit})
		case C_INSTRUCTION:
			p.instructions = append(p.instructions, &CInstruction{lit: lit})
		}
	}

	for _, instr := range p.instructions {
		switch i := instr.(type) {
		case *AInstruction:
			p.parseAInstruction(i)
		case *CInstruction:
			p.parseCInstruction(i)
		}
	}
	return HackFile{Instructions: p.instructions}
}

func (p *Parser) parseAInstruction(instr *AInstruction) {
	fmt.Println("A: " + instr.lit)
	a, err := strconv.ParseInt(instr.lit, 10, 15)
	if err != nil {
		panic(err)
	}
	instr.addr = int(a)
	// TODO
}

func (p *Parser) parseCInstruction(instr *CInstruction) {
	lit := instr.lit
	s := strings.Split(instr.lit, "=")
	if len(s) == 2 {
		instr.dest = 0b000
		if strings.Contains(s[0], "A") {
			instr.dest |= 0b100
		}
		if strings.Contains(s[0], "D") {
			instr.dest |= 0b010
		}
		if strings.Contains(s[0], "M") {
			instr.dest |= 0b001
		}
		lit = s[1]
	}
	s = strings.Split(lit, ";")
	if len(s) == 2 {
		switch s[1] {
		case "JGT":
			instr.jump = 0b001
			break
		case "JEQ":
			instr.jump = 0b010
			break
		case "JGE":
			instr.jump = 0b011
			break
		case "JLT":
			instr.jump = 0b100
			break
		case "JNE":
			instr.jump = 0b101
			break
		case "JLE":
			instr.jump = 0b110
			break
		case "JMP":
			instr.jump = 0b111
			break
		}
		lit = s[0]
	}

	fmt.Printf("C: %v, %v\n", instr.lit, len(s))
	// TODO
}
