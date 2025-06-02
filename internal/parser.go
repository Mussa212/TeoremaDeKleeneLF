package internal

var prec = map[TokenType]int{
	Union:  1,
	Concat: 2,
	Star:   3,
}

func ToPostfix(tokens []Token) []Token {
	var out, stack []Token
	for _, tok := range tokens {
		switch tok.Type {
		case Char, Lambda: // ahora Lambda va junto a Char
			out = append(out, tok)
		case Star:
			out = append(out, tok) // unario y de máxima prioridad
		case Union, Concat:
			for len(stack) > 0 && prec[stack[len(stack)-1].Type] >= prec[tok.Type] {
				out = append(out, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, tok)
		case LParen:
			stack = append(stack, tok)
		case RParen:
			for stack[len(stack)-1].Type != LParen {
				out = append(out, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// descartamos el '('
			stack = stack[:len(stack)-1]
		}
	}
	// vaciar la pila restante
	for len(stack) > 0 {
		out = append(out, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return out
}
