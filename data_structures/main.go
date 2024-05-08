package main

import (
	trees "data_structures/binaryTree"
	cl "data_structures/circularList"
	qe "data_structures/queue"
	st "data_structures/stack"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//using all the data structures
func generalDSTest() {

	fmt.Println("created empty int stack")
	var aStack st.Stack[int] = st.New[int]() // useless assignment, only for testing
	aStack.Push(1)
	fmt.Println("pushed: 1")
	aStack.Push(2)
	fmt.Println("pushed: 2")
	aStack.Push(122)
	fmt.Println("pushed: 122")
	fmt.Println("popped: ", aStack.Pop())
	t, _ := aStack.Top()
	fmt.Println("stack top is", t)
	fmt.Println("a stack", aStack)
	fmt.Println()

	fmt.Println("created empty int queue")
	var aQueue qe.Queue[int] = qe.New[int]() // useless assignment, only for testing
	aQueue.Enqueue(12)
	aQueue.Enqueue(124)
	aQueue.Enqueue(1)
	aQueue.Dequeue()
	fmt.Println("a queue", aQueue)
	fmt.Println()

	emptyStack := st.New[int]()
	fmt.Println("an empty stack", emptyStack)
	fmt.Println()

	emptyQueue := qe.New[int]()
	fmt.Println("an empty queue", emptyQueue)
	fmt.Println()

	fmt.Println("Circular Lists")
	var a cl.Circularlist
	a.Append(3)
	fmt.Println(a)
	a.Append(-3)
	a.Append(0)
	a.Append(1)
	a.Append(2)
	fmt.Println(a)
	fmt.Println(a.String2())

	fmt.Println("Print from zero:")
	a.PrintFromZero()

	fmt.Println()
	fmt.Println(a)
	fmt.Println("shift 3 (head) node forward 3 times")
	a.F2(0)
	p1 := a.GetPoiter(0)
	a.Move(p1)
	fmt.Println(a)
	fmt.Println("shift -3 node backwards 3 times")
	p2 := a.GetPoiter(2)
	a.Move(p2)
	fmt.Println(a)
	fmt.Println("3 is still the head of the list")
}

// stack exercise
// we assume that the input expr is legal or well formed postfix math expression
func postFixEvaluation(expr string) int {
	stack := st.New[int]()
	exprs := strings.Split(expr, " ")
	var op1, op2, res, num int
	for _, tk := range exprs {
		if unicode.IsDigit(rune(tk[0])) {
			num, _ = strconv.Atoi(tk)
			stack.Push(num)
		} else if tk[0] == '+' || tk[0] == '-' || tk[0] == '/' || tk[0] == '*' {
			op1 = stack.Pop()
			op2 = stack.Pop()
			switch rune(tk[0]) {
			case '+':
				res = op2 + op1
			case '-':
				res = op2 - op1
			case '*':
				res = op2 * op1
			case '/':
				res = op2 / op1
			}
			stack.Push(res)
		}

	}
	return stack.Pop()
}

// stack exercise
// we assume that the input expr is legal or well formed infix math expression
func infixEvalutation(expr string) string {
	exprs := strings.Split(expr, " ")
	stack := st.New[byte]()
	resexpr := ""
	for _, tk := range exprs {
		if unicode.IsDigit(rune(tk[0])) {
			resexpr += tk + " "
		} else if tk[0] == ')' {
			resexpr += string(stack.Pop()) + " "
		} else if tk[0] == '+' || tk[0] == '-' || tk[0] == '/' || tk[0] == '*' {
			stack.Push(tk[0])
		}
	}
	return resexpr
}

// stack tags exercise utility
func isOpenTag(tag string) bool {
	match, _ := regexp.MatchString("[<][a-z]+[>]", tag)
	return match
}

// stack tags exercise utility
func isCloseTag(tag string) bool {
	match, _ := regexp.MatchString("[<][\\/][a-z]+[>]", tag)
	return match
}

func isWellFormedDocument(doc string) (bool, error) {
	stack := st.New[string]()
	tags := strings.Split(doc, " ")

	var ptag string
	for i, tag := range tags {
		if isOpenTag(tag) {
			stack.Push(tag)
		} else {
			if isCloseTag(tag) {
				ptag = stack.Pop()
				if ptag != strings.Replace(tag, "/", "", -1) {
					return false, errors.New("error at pos " + fmt.Sprint(i))
				}
			} else {
				return false, errors.New("illegal token at pos " + fmt.Sprint(i))
			}
		}
	}
	// check that there is nothing inside the stack (change stack top fun to return error )
	_, e := stack.Top()
	var t string
	err := ""
	for e == nil {
		t = stack.Pop()
		err += " " + t
		_, e = stack.Top()
	}

	if len(err) > 0 {
		return false, errors.New("these tags remained open " + err)
	}

	return true, nil
}

func stackPlay() {
	expr1 := "5 3 - 2 *"
	expr2 := "2 5 3 - *"
	expr3 := "2 5 3 - * 92 + 8 * 78 -"
	fmt.Println("postfix math expression notation")
	fmt.Println(expr1, " = ", postFixEvaluation(expr1))
	fmt.Println(expr2, " = ", postFixEvaluation(expr2))
	fmt.Println(expr3, " = ", postFixEvaluation(expr3))
	fmt.Println()

	fmt.Println("correct document tags check")
	doc1 := "<a> <b> </b> </c>"
	r, e := isWellFormedDocument(doc1)
	fmt.Println(doc1, ": ", r, " ", e)

	doc2 := "<a> <b> </b> <c> <d> </d>"
	r, e = isWellFormedDocument(doc2)
	fmt.Println(doc2, ": ", r, " ", e)

	expr1 = "( ( 5 - 3 ) * 2 )"
	fmt.Println("infix expression notation conversion to postfix")
	fmt.Println(expr1, " = ", infixEvalutation(expr1))
}

func binTreePlay() {
	rootNode := trees.NewNode(78,
		trees.NewNode(54,
			nil,
			trees.NewNode(90,
				trees.NewNode(19,
					nil,
					nil),
				trees.NewNode(95,
					nil,
					nil))),
		trees.NewNode(21,
			trees.NewNode(16,
				trees.NewNode(5,
					nil,
					nil),
				nil),
			trees.NewNode(19,
				trees.NewNode(56,
					nil,
					nil),
				trees.NewNode(43,
					nil,
					nil))))
	tree := trees.NewTree(rootNode)

	fmt.Println(tree)
	fmt.Println("preorder   visit: ", tree.PreOrderVisit())
	fmt.Println("inorder    visit: ", tree.InOrderVisit())
	fmt.Println("postorder  visit: ", tree.PostOrderVisit())

	slice := []int{1, 3, 2, 13, 5, 98, 12, 34}
	fmt.Println()
	fmt.Println(slice)
	fmt.Println()
	treeFromSlice := trees.ArrayToTree(slice)
	fmt.Println(treeFromSlice)
}

func main() {
	fmt.Println("DATA STRUCTURES")
	generalDSTest()
	fmt.Println()

	fmt.Println("STACK EXERCISES")
	stackPlay()
	fmt.Println()

	fmt.Println("BINARY TREES")
	binTreePlay()
	fmt.Println()

}
