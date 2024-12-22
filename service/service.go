package service

import (
	"fmt"
	"strconv"
	"strings"
)

// func main() {
// 	fmt.Println(Calc("1+1*"))
// }

func Calc(expression string) (float64, error) {
	arr := strings.Split(expression, "")
	var stack []string
	var queue []string //выходная строка

	//Проверка
	count_digits := 0
	count_operators := 0
	count_brackets := 0
	for i := 0; i < len(arr); i++ {
		s := arr[i]

		if IsDigit(s) == false && IsOperation(s) == false && s != "(" && s != ")" {
			return 0, fmt.Errorf("invalid sybmol")
		}

		if IsDigit(s) {
			count_digits++
		}
		if IsOperation(s) {
			count_operators++
		}
		if s == "(" {
			count_brackets++
		}
		if s == ")" {
			count_brackets--
		}
		if count_brackets < 0 {
			return 0, fmt.Errorf("invalid input")
		}
	}

	if count_operators == count_digits {
		return 0, fmt.Errorf("invalid input")
	}

	//получим в обратной польской записи

	//цикл по всем символам
	for i := 0; i < len(arr); i++ {
		s := arr[i]

		if IsDigit(s) {
			//если число - помещаем в выходную строку
			queue = append(queue, s)
		}
		if IsOperation(s) {
			//оператор
			if len(stack) == 0 || stack[len(stack)-1] == "(" {
				//Если в стеке пусто, или в стеке открывающая скобка
				//добавляем оператор в стек
				stack = PushToStack(s, stack)
			} else if GetPriority(s) > GetPriority(stack[len(stack)-1]) {
				//Если входящий оператор имеет более высокий приоритет чем вершина stack
				//добавляем оператор в стек
				stack = PushToStack(s, stack)
			} else if GetPriority(s) <= GetPriority(stack[len(stack)-1]) {
				//Если оператор имеет более низкий или равный приоритет, чем в стеке,
				//выгружаем POP в очередь (QUEUE),
				//пока не увидите оператор с меньшим приоритетом или левую скобку на вершине (TOP),
				stack, queue = PopStackToQueue(s, stack, queue)
				// затем добавьте (PUSH) входящий оператор в стек (STACK).
				stack = PushToStack(s, stack)
			}
		} else if s == "(" {
			//Если входящий элемент является левой скобкой,
			// поместите (PUSH) его в стек (STACK).
			stack = PushToStack(s, stack)
		} else if s == ")" {
			//Если входящий элемент является правой скобкой,
			//выгружаем стек (POP) и добавляем его элементы в очередь (QUEUE),
			//пока не увидите левую круглую скобку.
			//Удалите найденную скобку из стека (STACK).
			stack, queue = PopStackToQueue(s, stack, queue)
		}

	}
	// В конце выражения выгрузите стек (POP) в очередь (QUEUE)
	stack, queue = PopStackToQueue("", stack, queue)

	//   посчитаем выражение
	//   queue_new := make([]string, cap(queue), cap(queue))
	stack_new := make([]float64, cap(stack), cap(queue))

	for i := 0; i < len(queue); i++ {
		curr := queue[i]
		curr_float64, _ := strconv.ParseFloat(strings.TrimSpace(curr), 64)
		if IsDigit(curr) {
			stack_new = append(stack_new, curr_float64)
		} else {
			a := stack_new[len(stack_new)-1]
			stack_new = RemoveItemInIntSlice(stack_new, len(stack_new)-1)

			b := stack_new[len(stack_new)-1]
			stack_new = RemoveItemInIntSlice(stack_new, len(stack_new)-1)

			stack_new = append(stack_new, Operation(b, a, curr))
		}

	}

	return stack_new[len(stack_new)-1], nil
}

func Operation(a, b float64, o string) float64 {
	if o == "*" {
		return a * b
	}
	if o == "/" {
		return a / b
	}
	if o == "+" {
		return a + b
	}
	if o == "-" {
		return a - b
	}
	return 0
}

func PushToStack(s string, stack []string) []string {
	stack = append(stack, s)
	return stack
}

func PopStackToQueue(s string, stack, queue []string) ([]string, []string) {
	//сделаем копию стека
	stack_temp := stack
	for i := 0; i < len(stack); i++ {
		stack_temp[i] = stack[i]
	}
	if s != "" {
		//если s != "" идем в обратном порядке до "(" или с меньшим приоритетом
		for i := len(stack) - 1; i >= 0; i-- {
			curr := stack[i]
			if GetPriority(curr) < GetPriority(s) || curr == "(" {
				//stop
				if curr == "(" {
					stack_temp = RemoveItemInSlice(stack_temp, i)
				}
				break
			} else {
				queue = append(queue, curr)
				stack_temp = RemoveItemInSlice(stack_temp, i)
			}
		}
	} else {
		for i := len(stack) - 1; i >= 0; i-- {
			curr := stack[i]
			queue = append(queue, curr)
			stack_temp = RemoveItemInSlice(stack_temp, i)
		}
	}

	return stack_temp, queue
}

func GetPriority(s string) int {
	if s == "*" || s == "/" {
		return 1
	}
	return 0
}

func IsOperation(s string) bool {
	// проверяет является ли сивол операцией
	if s == "*" || s == "/" || s == "+" || s == "-" {
		return true
	}
	return false
}

func IsDigit(s string) bool {
	// проверяет является ли сивол операцией
	if s == "0" || s == "1" || s == "2" || s == "3" || s == "4" || s == "5" || s == "6" || s == "7" || s == "8" || s == "9" {
		return true
	}
	return false
}

func RemoveItemInSlice(slice []string, ind int) []string {

	return append(slice[:ind], slice[ind+1:]...)
}

func RemoveItemInIntSlice(slice []float64, ind int) []float64 {

	return append(slice[:ind], slice[ind+1:]...)
}
