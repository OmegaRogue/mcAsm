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

type Scanner struct {
	src      []byte
	ch       rune // current character
	offset   int  // character offset
	rdOffset int  // reading offset
}

const bom = 0xFEFF

func (s *Scanner) Init(src []byte) {
	s.src = src
	s.ch = ' '
	s.offset = 0
	s.rdOffset = 0

	s.next()
	if s.ch == bom {
		s.next()
	}
}

func (s *Scanner) next() {
	if s.rdOffset < len(s.src) {
		s.offset = s.rdOffset
		s.ch = rune(s.src[s.rdOffset])
		s.rdOffset++
	} else {
		s.offset = len(s.src)
		s.ch = -1 // eof
	}
}

func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.next()
	}
}

func (s *Scanner) scanComment() string {
	s.next()
	offs := s.offset
	for s.ch != '\n' && s.ch >= 0 {
		s.next()
	}
	return string(s.src[offs:s.offset])
}

func (s *Scanner) scanLine() string {
	offs := s.offset
	for s.ch != '\n' && s.ch != '\r' && s.ch >= 0 && s.ch != ' ' {
		s.next()
	}
	return string(s.src[offs:s.offset])
}

func (s *Scanner) scanLabel() string {
	offs := s.offset
	for {
		ch := s.ch
		if ch == '\n' || ch == '\r' || ch < 0 {
			break
		}
		s.next()
		if ch == ')' {
			break
		}
	}
	return string(s.src[offs : s.offset-1])
}

func isCInstruction(ch rune) bool {
	return ch == '0' || ch == '1' || ch == '-' || ch == '!' || ch == 'A' || ch == 'D' || ch == 'M'
}

func (s *Scanner) Scan() (tok Token, lit string) {
	s.skipWhitespace()

	switch ch := s.ch; {
	case isCInstruction(ch):
		tok = C_INSTRUCTION
		lit = s.scanLine()
	default:
		s.next() // always make progress
		switch ch {
		case -1:
			tok = EOF
		case '/':
			if s.ch == '/' {
				tok = COMMENT
				lit = s.scanComment()
			}
		case '(':
			tok = LABEL
			lit = s.scanLabel()
		case '@':
			tok = A_INSTRUCTION
			lit = s.scanLine()
		default:
			tok = ILLEGAL
		}
	}
	return
}
