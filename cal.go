// test1
package main

import (
	"bufio"
	"fmt"
	"os"
)

//保存表达式中的元素，flag为true标记为运算数，false标记为运算符号,运算符号按照type数值存储
type Part struct {
	flag bool
	num  float64
}

//声明表达式元素的栈
type Stack struct {
	data [100]Part
	top  int
}

//栈的初始化
func (s *Stack) InitStack() {
	s.top = -1
}

//入栈
func (s *Stack) Push(part Part) {
	s.top++
	s.data[s.top] = part
}

//出栈
func (s *Stack) Pop() {
	s.top--
}

//判断栈是否为空，空栈返回true
func (s Stack) Empty() bool {
	if s.top == -1 {
		return true
	} else {
		return false
	}
}

//返回栈中元素个数
func (s Stack) Size() int {
	return (s.top + 1)
}

//返回栈顶元素
func (s Stack) Top() Part {
	return s.data[s.top]
}

func judge(a Part, b Part) bool {
	tmpa := a.num
	tmpb := b.num
	//type ()+-*/ 40,41,43,45,42,47
	switch tmpa {
	case 43, 45:
		if tmpb == 42 || tmpb == 47 {
			return true
		}
	case 40:
		return true
	}
	return false
}

//将一个字符串处理成Part数组的中缀表达式
func ChangeToPart(exp string, a []Part) int {
	top := -1
	num := -1.0
	//type ()+-*/ 40,41,43,45,42,47
	//type 0~9 48~57
	l := len(exp)
	for i := 0; i < l; i++ {
		switch exp[i] {
		case 40, 41, 43, 45, 42, 47:
			if num != -1.0 {
				top++
				a[top].flag = true
				a[top].num = num
				num = -1.0
			}
			top++
			tmp := float64(exp[i])
			a[top].num = tmp
			a[top].flag = false
		default:
			if num == -1.0 {
				num = 0
			}
			tmp := float64(exp[i])
			num = 10*num + tmp - 48
		}
	}
	if num != -1 {
		top++
		a[top].flag = true
		a[top].num = num
		num = -1
	}
	return (top + 1)
}

//将中缀表达式处理为后缀表达式，返回后缀表达式的长度
func InToPost(a []Part, nums int) int {
	var stack Stack
	newnums := -1
	stack.InitStack()
	for i := 0; i < nums; i++ {
		tmp := a[i]
		if tmp.flag {
			newnums++
			a[newnums] = a[i]
		} else {
			//type ()+-*/ 40,41,43,45,42,47
			switch a[i].num {
			case 40:
				stack.Push(a[i])
			case 41:
				for !stack.Empty() {
					top := stack.Top()
					if top.num == 40 {
						stack.Pop()
						break
					}
					newnums++
					a[newnums] = top
					stack.Pop()
				}
			default:
				for !stack.Empty() {
					top := stack.Top()
					if top.num == 40 || judge(top, a[i]) {
						break
					}
					newnums++
					a[newnums] = top
					stack.Pop()
				}
				stack.Push(a[i])
			}
		}
	}
	for !stack.Empty() {
		newnums++
		a[newnums] = stack.Top()
		stack.Pop()
	}
	return newnums
}

//计算后缀表达式
func cal(a []Part, nums int) float64 {
	var stack Stack
	stack.InitStack()
	for i := 0; i <= nums; i++ {
		tmp := a[i].num
		if a[i].flag {
			stack.Push(a[i])
		} else {
			a := stack.Top()
			stack.Pop()
			b := stack.Top()
			stack.Pop()
			var s Part
			s.flag = true
			//type ()+-*/ 40,41,43,45,42,47
			switch tmp {
			case 43:
				s.num = a.num + b.num
				stack.Push(s)
			case 45:
				s.num = a.num - b.num
				stack.Push(s)
			case 42:
				s.num = a.num * b.num
				stack.Push(s)
			case 47:
				s.num = a.num * 1.0 / b.num
				stack.Push(s)
			}
		}
	}
	result := stack.Top().num
	return result
}
func main() {
	str := bufio.NewScanner(os.Stdin)
	str.Scan()
	exp := str.Text()
	var partsA [100]Part
	var PartsA []Part = partsA[:]
	nums := ChangeToPart(exp, PartsA)
	nums = InToPost(PartsA, nums)
	fmt.Println(cal(PartsA, nums))
}
