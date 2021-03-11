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

import "fmt"

type Instruction interface {
	BinaryString() string
}

type HackFile struct {
	Instructions []Instruction
}

type AInstruction struct {
	lit  string // the raw assembly instruction pre-parsing
	addr int
}

func (a *AInstruction) BinaryString() string {
	return fmt.Sprintf("0%015b\n", a.addr)
}

type CInstruction struct {
	lit  string
	dest int // C instructions look like: `dest=comp;jump`; dest and jump are optional
	comp int
	jump int
}

func (c *CInstruction) BinaryString() string {
	return fmt.Sprintf("111%07b%03b%03b\n", c.comp, c.dest, c.jump)
}

//                     ADM
var t = 0b111_0111111_001_000
var t2 = 0b111_0101010_001_000

var t3 = 0b111_1110000_010_000
