package internal

type TokenType int

const (
	Char   TokenType = iota // símbolos del alfabeto
	Lambda                  // 'λ' (vacío)
	Union                   // '+'
	Star                    // '*'
	LParen                  // '('
	RParen                  // ')'
	Concat                  // '.' ← insertado implícitamente
)

type Token struct {
	Type TokenType
	Val  rune
}

func Scan(input string) []Token {
	var out []Token
	prev := Token{Type: Union} // para que no inserte Concat al primer carácter

	for _, r := range input {
		var t Token
		switch r {
		case '+':
			t = Token{Union, r}
		case '*':
			t = Token{Star, r}
		case '(':
			t = Token{LParen, r}
		case ')':
			t = Token{RParen, r}
		case 'λ':
			t = Token{Lambda, r}
		default:
			t = Token{Char, r}
		}

		// Inserta concatenación implícita si hace falta
		if needsConcat(prev, t) {
			out = append(out, Token{Concat, '.'})
		}
		out = append(out, t)
		prev = t
	}
	return out
}

func needsConcat(a, b Token) bool {
	// literal, λ, cierre o estrella seguido de literal, λ o '('
	left := a.Type == Char || a.Type == Lambda || a.Type == RParen || a.Type == Star
	right := b.Type == Char || b.Type == Lambda || b.Type == LParen
	return left && right
}
