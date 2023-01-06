package main

func isValid(s string) bool {
	m := make(map[rune]rune)

	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'

	stack := make([]rune, 0)

	for _, symbol := range s {
		if symbol == '(' || symbol == '[' || symbol == '{' {
			stack = append(stack, symbol)
			continue
		}

		if symbol == ')' || symbol == ']' || symbol == '}' {

			if len(stack) == 0 {
				return false
			}

			if m[symbol] != stack[len(stack)-1] {
				return false
			}

			stack = stack[0 : len(stack)-1]
		}

	}

	return len(stack) == 0
}
