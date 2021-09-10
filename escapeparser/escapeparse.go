package escapeparser

import (
	"errors"
	"io"
)

func Tokenise(buf string) ([]Token, error) {
	s := &stringStream{buf, len(buf), 0}
	tokens := []Token{}

	pg := tokenGen{LXAny}

	for escape := false; !escape; {
		// log.Println("TokeniseLoop")
		cc, err := pg.Lex(s)
		tokens = append(tokens, cc)
		if err != nil {
			break
		}
		// switch cc.Type {
		// case TokWIPNone:
		// 	return tokens, ErrNoneToken
		// case TokWIPEsc:
		// 	return tokens, ErrDanglingEscape
		// case TokEnd:
		// 	escape = true
		// default:
		// 	//nothing
		// }
	}

	return tokens, nil
}

var (
	ErrDanglingEscape = errors.New(`Dangling Escape`)
	ErrNoneToken      = errors.New(`Tokens???`)
)

type TokenType = int

const (
	TokWIPNone TokenType = iota
	TokEnd
	TokWIPEsc
	TokEscVar
	TokEscLit
	TokStr
	TokSpace
)

type Token struct {
	Type  TokenType
	Value string
	Start int
}

type LexState = int

const (
	LXAny LexState = iota
	LXEsc
	LXEscVar
	LXStr
)

type tokenGen struct {
	State LexState
}

func (g *tokenGen) Lex(s Stream) (Token, error) {
	var cc Token
	for {
		part, pos, err := s.Next()
		//log.Printf("LexLoop: `%s`, %d \n State: %v\n Token: %+v\n ERR:: %v\n", part, pos, g.State, cc, err)
		if err != nil {
			if errors.As(err, &io.EOF) {
				return cc, err
			}
		}

		switch g.State {
		case LXAny:
			switch part {
			case "$":
				// if cc.Value != ""{
				// 	g.Back(1)
				// NOPE That isnt gonna work...
				// }
				cc = Token{Type: TokWIPEsc, Value: "", Start: pos}
				g.State = LXEsc
			default:
				cc.Value += part
			}
		case LXEsc:
			if part == "$" {
				cc.Type = TokEscLit
				return cc, nil
			} else {
				cc.Type = TokEscVar
				cc.Value += part
			}
			break
		case LXEscVar:
			if part == "$" {
				return cc, nil
			} else {
				cc.Value += part
			}
		default:
			//BRRR
			panic("LexDefault")
			// return cc, nil
		}

	}
	panic("LexEnd")
	return cc, nil
}

type Stream interface {
	Next() (Piece string, Postition int, EOF error)
	Back(int)
	// Peek() string
}

type stringStream struct {
	buffer string
	len    int
	pos    int
}

func (sa *stringStream) Next() (string, int, error) {
	if sa.pos < sa.len {
		s, p := string(sa.buffer[sa.pos]), sa.pos
		sa.pos += 1
		return s, p, nil
	}
	return "", sa.pos, io.EOF
}

func (sa *stringStream) Back(n int) {
	sa.pos -= n
}
