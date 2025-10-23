package main

import (
	"fmt"
)

func main() {
	/*var inp string
	fmt.Println("Введите число: ")
	fmt.Scanln("%d", &inp)
	N, err := strconv.Atoi(inp)
	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
	} else {
		fmt.Println(N) // Вывод: 123
	}*/
	N := 1238 //не успел разобраться с тем, как вводить число с консоли:(
	if N >= 12307 {
		fmt.Println("error; write a number less than 12307")
	}
	for N < 12307 {
		if N < 0 {
			N = N * (-1)
		} else if N%7 == 0 {
			N = N * 39
		} else if N%9 == 0 {
			N = N*13 + 1
		} else {
			N = (N + 2) * 3
		}
		if N%7 == 0 && N%13 == 0 {
			fmt.Println("service error")
			break
		} else {
			N++
		}
	}
	fmt.Println(N)
	fmt.Println(optional(N))
}
func optional(n int) string {
	num1 := n % 10                          //единицы
	num2 := (n%100 - num1) / 10             //десятки
	num3 := (n%1000 - n%100) / 100          //сотни
	num4 := (n%10000 - n%1000) / 1000       //тысячи
	num5 := (n%100000 - n%10000) / 10000    //десятки тысяч
	num6 := (n%1000000 - n%100000) / 100000 //сотни тысяч
	var str string
	switch num6 {
	case 0:
		{
			str += ""
		}
	case 1:
		{
			str += "Сто "
		}
	case 2:
		{
			str += "Двести "
		}
	case 3:
		{
			str += "Триста "
		}
	case 4:
		{
			str += "Четыреста "
		}
	} //max(N)<500000
	switch num5 {
	case 0:
		{
			str += " "
		}
	case 1:
		{
			if num4 == 0 {
				str += "десять "
			}
		}
	case 2:
		{
			str += "двадцать "
		}
	case 3:
		{
			str += "Тридцать "
		}
	case 4:
		{
			str += "сорок "
		}
	case 5:
		{
			str += "пятьдесят "
		}
	case 6:
		{
			str += "шестьдесят "
		}
	case 7:
		{
			str += "семьдесят "
		}
	case 8:
		{
			str += "восемьдесят "
		}
	case 9:
		{
			str += "девяносто "
		}
	}

	switch num4 {
	case 0:
		{
			str += "тысяч "
		}
	case 1:
		{
			if num5 == 1 {
				str += "одиннадцать тысяч"
			} else {
				str += "одна тысяча "
			}
		}
	case 2:
		{
			if num5 == 1 {
				str += "двенадцать тысяч"
			} else {
				str += "две тысячи "
			}
		}
	case 3:
		{
			if num5 == 1 {
				str += "тринадцать тысяч"
			} else {
				str += "три тысячи "
			}
		}
	case 4:
		{
			if num5 == 1 {
				str += "четырнадцать тысяч"
			} else {
				str += "четыре тысячи "
			}
		}
	case 5:
		{
			if num5 == 1 {
				str += "пятнадцать тысяч"
			} else {
				str += "пять тысяч "
			}
		}
	case 6:
		{
			if num5 == 1 {
				str += "шестнадцать тысяч"
			} else {
				str += "шесть тысяч "
			}
		}
	case 7:
		{
			if num5 == 1 {
				str += "семнадцать тысяч"
			} else {
				str += "семь тысяч "
			}
		}
	case 8:
		{
			if num5 == 1 {
				str += "восемнадцать тысяч"
			} else {
				str += "восемь тысяч "
			}
		}
	case 9:
		{
			if num5 == 1 {
				str += "девятнадцать тысяч"
			} else {
				str += "девять тысяч "
			}
		}
	}

	switch num3 {
	case 0:
		{
		}
	case 1:
		{
			str += "сто "
		}
	case 2:
		{
			str += "двести "
		}
	case 3:
		{
			str += "Триста "
		}
	case 4:
		{
			str += "четыреста "
		}
	case 5:
		{
			str += "пятьтсот "
		}
	case 6:
		{
			str += "шестьсот "
		}
	case 7:
		{
			str += "семьсот "
		}
	case 8:
		{
			str += "восемьсот "
		}
	case 9:
		{
			str += "девятьсот "
		}
	}

	switch num2 {
	case 0:
		{

		}
	case 1:
		{
			if num1 == 0 {
				str += "десять "
			}
		}
	case 2:
		{
			str += "двадцать "
		}
	case 3:
		{
			str += "Тридцать "
		}
	case 4:
		{
			str += "сорок "
		}
	case 5:
		{
			str += "пятьдесят "
		}
	case 6:
		{
			str += "шестьдесят "
		}
	case 7:
		{
			str += "семьдесят "
		}
	case 8:
		{
			str += "восемьдесят "
		}
	case 9:
		{
			str += "девяносто "
		}
	}

	switch num1 {
	case 0:
		{
		}
	case 1:
		{
			if num2 == 1 {
				str += "одиннадцать"
			} else {
				str += "один"
			}
		}
	case 2:
		{
			if num2 == 1 {
				str += "двенадцать"
			} else {
				str += "два"
			}
		}
	case 3:
		{
			if num2 == 1 {
				str += "тринадцать"
			} else {
				str += "Три"
			}
		}
	case 4:
		{
			if num2 == 1 {
				str += "четырнадцать"
			} else {
				str += "четыре"
			}
		}
	case 5:
		{
			if num2 == 1 {
				str += "пятнадцать"
			} else {
				str += "пять"
			}
		}
	case 6:
		{
			if num2 == 1 {
				str += "шестнадцать"
			} else {
				str += "шесть"
			}
		}
	case 7:
		{
			if num2 == 1 {
				str += "семнадцать"
			} else {
				str += "семь"
			}
		}
	case 8:
		{
			if num2 == 1 {
				str += "восемнадцать"
			} else {
				str += "восемь"
			}
		}
	case 9:
		{
			if num2 == 1 {
				str += "девятнадцать"
			} else {
				str += "девять"
			}
		}
	}
	return str
}
