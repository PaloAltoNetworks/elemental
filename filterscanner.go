// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package elemental

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

// scanner scans a given input
type scanner struct {
	buf          bytes.Buffer
	isWhitespace checkRuneFunc
	isLetter     checkRuneFunc
	isDigit      checkRuneFunc
}

// newScanner returns an instance of a Scanner.
func newScanner(
	input string,
) *scanner {
	var buf bytes.Buffer
	_, _ = buf.WriteString(input)

	return &scanner{
		buf:          buf,
		isWhitespace: isWhitespace,
		isLetter:     isLetter,
		isDigit:      isDigit,
	}
}

// read returns the next rune or eof
func (s *scanner) read() rune {

	ch, _, err := s.buf.ReadRune()
	if err != nil {
		return runeEOF
	}
	return ch
}

// peekNextRune returns the next rune but does not read it.
func (s *scanner) peekNextRune() rune {

	if s.buf.Len() == 0 {
		return runeEOF
	}

	ch, _ := utf8.DecodeRune(s.buf.Bytes()[0:])
	return ch
}

// unread a previously read rune
func (s *scanner) unread() {
	_ = s.buf.UnreadRune()
}

// scan returns the next token and literal value.
func (s *scanner) scan() (parserToken, string) {

	ch := s.read()

	if s.isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	}

	if isOperatorStart(ch) {
		// Chack if the next run can create an operator
		nextCh := s.peekNextRune()
		if nextCh != runeEOF {
			if token, literal, ok := isOperator(ch, nextCh); ok {
				s.read() // read only if it has matched.
				return token, literal
			}
		}

		// Check if the current rune is an operator
		if token, literal, ok := isOperator(ch); ok {
			return token, literal
		}
	}

	if s.isLetter(ch) || s.isDigit(ch) {
		s.unread()
		return s.scanWord()
	}

	token, ok := runeToToken[ch]
	if !ok {
		return parserTokenILLEGAL, string(ch)
	}

	return token, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *scanner) scanWhitespace() (parserToken, string) {

	var buf bytes.Buffer
	_, _ = buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == runeEOF {
			break
		} else if !s.isWhitespace(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return parserTokenWHITESPACE, buf.String()
}

// scanWord consumes the current rune and all contiguous letters / digits.
func (s *scanner) scanWord() (parserToken, string) {

	var buf bytes.Buffer
	_, _ = buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == runeEOF {
			break
		} else if !s.isLetter(ch) && !s.isDigit(ch) {
			s.unread()
			break
		} else {
			if ch == '\\' {
				// Move forward
				ch = s.read()
			}

			if isOperatorStart(ch) {
				// Check if the next rune can create an operator
				nextCh := s.peekNextRune()
				if nextCh != runeEOF {
					if _, _, ok := isOperator(ch, nextCh); ok {
						s.unread() // unread ch rune
						break
					}
				}

				// Check if ch is an operator by itself
				if _, _, ok := isOperator(ch); ok {
					s.unread()
					break
				}
			}

			_, _ = buf.WriteRune(ch)
		}
	}

	return stringToToken(buf.String())
}

func stringToToken(output string) (parserToken, string) {

	upper := strings.ToUpper(output)

	if token, ok := wordToToken[upper]; ok {
		return token, output
	}

	if token, ok := operatorsToToken[upper]; ok {
		return token, output
	}

	return parserTokenWORD, output
}
