/*База:
Реализовать свой стек(Stack)
Методы которые должны быть обязательно:
Pop()
Push()
IsEmpty()
Size()
Clear()*/

package main

import (
	"fmt"
	"os"
)

type Stack []int

func main() {
	var stack Stack
	var length int
	var elem int
	fmt.Printf("Введите длину стека:")
	fmt.Scanf("%d/n", &length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &elem)
		stack.Push(elem)
	}
	if stack.IsEmpty() {
		fmt.Println("Стек пустой, попробуйте ввести ненулевую длинну стека")
		os.Exit(1)
	}
	fmt.Printf("Размер стека равен %d\n\r", stack.Size())
	fmt.Printf("Последний помещённый в стек элемент равен %d\n\r", stack.Pop())
	stack = stack.Clear()
	if stack.IsEmpty() {
		fmt.Printf("Стек был успешно очищен")
	} else {
		fmt.Printf("Ошибка: стек не был очищен")
	} //проверка
}

func (s *Stack) Size() int { //узнаёт размер стека
	return len(*s)
}

func (s *Stack) IsEmpty() bool { //проверяет пустой стек или нет
	return len(*s) == 0 //возврат true  или false в зависимости от истинности выражения
}

func (s *Stack) Push(item int) { //Кладёт в стек элемент
	*s = append(*s, item)
}
func (s *Stack) Pop() int { //достаёт верхний элемент из стека и удаляет его
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}
func (s Stack) Clear() Stack { //очистка стека
	s = nil
	return s
}
