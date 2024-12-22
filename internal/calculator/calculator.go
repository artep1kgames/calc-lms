package calculator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	r, err := regexp.Compile(`\([^()]*\)|\(([^()]*\([^()]*\)[^()]*)+\)`)
	if err != nil {
		return 0.0, err
	}

	matches := r.FindAllString(expression, -1)
	expression = r.ReplaceAllLiteralString(expression, "%f")
	matchesResults := []any{}
	for _, exp := range matches {
		resT, _ := Calc(exp[1 : len([]rune(exp))-1])
		matchesResults = append(matchesResults, resT)
	}
	expression = fmt.Sprintf(expression, matchesResults...)

	exp := strings.Split(expression, "")

	res := []string{}
	memory := ""
	for i, el := range exp {
		if NumberCheck(el) {
			memory += el
		} else if DotCheck(el) {
			memory += el
		} else if ActCheck(el) {
			if i == len(exp)-1 || ActCheck(exp[i-1]) {
				return 0.0, fmt.Errorf("Uncorrect input")
			}
			res = append(res, memory)
			memory = ""
			res = append(res, el)
		}
	}
	res = append(res, memory)
	return Calculate(res)
}

func DotCheck(symbol string) bool {
	r_symbol := []rune(symbol)[0]
	if r_symbol == 46 {
		return true
	}
	return false
}

func ActCheck(symbol string) bool {
	r_symbol := []rune(symbol)[0]
	if r_symbol == 42 || r_symbol == 43 || r_symbol == 45 || r_symbol == 47 {
		return true
	}
	return false
}

func NumberCheck(symbol string) bool {
	r_symbol := []rune(symbol)[0]
	if r_symbol >= 48 && r_symbol <= 57 {
		return true
	}
	return false
}

func Act(a, b float64, act string) float64 {
	switch act {
	case "/":
		return a / b
	case "*":
		return a * b
	case "+":
		return a + b
	case "-":
		return a - b
	}
	return 0.0
}

func Calculate(exp []string) (float64, error) {
	for _, act := range "*/+-" {
		for i := 0; i < len(exp); i++ {
			if string(act) == exp[i] {
				if i+1 >= len(exp) || i-1 < 0 || i == 0 {
					return 0.0, errors.New("Uncorrect input")
				}
				fs, _ := strconv.ParseFloat(exp[i-1], 64)
				ss, _ := strconv.ParseFloat(exp[i+1], 64)
				result := Act(fs, ss, string(act))
				var cExp []string
				cExp = append(cExp, exp[:i-1]...)
				cExp = append(cExp, fmt.Sprintf("%f", result))
				cExp = append(cExp, exp[i+2:]...)
				exp = cExp
				i--
			}
		}
	}
	return strconv.ParseFloat(exp[0], 64)
}

func main() {
	expression := "1 + 902 . 02 * 2"
	answer, _ := Calc(expression)
	fmt.Println(answer)
}
