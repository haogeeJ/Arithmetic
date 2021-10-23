package main

import (
	"container/list"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	m            = make(map[*FAL]int)
	problemList  []*Problem
	n            int
	r            int64
	exercisefile string
	answerfile   string
	Divedzero    bool
)

func randData(l *list.List, r int64, p *Problem) {
	var (
		num  int64
		nume int64
		deno int64
	)
	num = rand.Int63n(r) + 1
	deno = rand.Int63n(r) + 1
	if (rand.Intn(7) % 6) == 0 {
		nume = rand.Int63n(deno)
		if nume == 0 {
			deno = 1
		}
	} else {
		nume = 0
		deno = 1
	}
	if num != 0 || nume == 0 {
		p.formulaTostring += strconv.FormatInt(num, 10)
	}
	if nume != 0 {
		if num != 0 {
			p.formulaTostring += "'"
		}
		p.formulaTostring += strconv.FormatInt(nume, 10)
		p.formulaTostring += "/"
		p.formulaTostring += strconv.FormatInt(deno, 10)
	}
	//fmt.Println(num, nume, deno)
	f := Model(num, nume, deno)
	l.PushBack(&Entry{
		kind:  1,
		value: f,
	})
}

func getProblem() (*Problem, bool) {

	//运算符个数
	counts := rand.Intn(3) + 1

	//随机产生运算符，真分数和整数,拼接
	var (
		l    *list.List
		flag = false
		cnt  = 0
	)
	l = list.New()
	p := Problem{
		formula:        l,
		postfixExpress: list.New(),
	}
	for i := 0; i < counts; i++ {
		if !flag && rand.Intn(2)%2 == 0 && counts > 1 {
			p.formulaTostring += "("
			l.PushBack(&Entry{
				kind:  2,
				value: &Sign{s: '('},
			})
			flag = true
			cnt = 0
		}
		if flag {
			cnt++
		}
		randData(l, r, &p)
		if cnt == 2 && flag {
			flag = false
			p.formulaTostring += ")"
			l.PushBack(&Entry{
				kind:  2,
				value: &Sign{s: ')'},
			})
		}
		s := new(Sign)
		switch rand.Intn(4) {
		case 0:
			s.s = '+'
		case 1:
			s.s = '-'
		case 2:
			s.s = '×'
		case 3:
			s.s = '÷'
		}
		p.formulaTostring += string(s.s)
		l.PushBack(&Entry{
			kind:  2,
			value: s,
		})
	}
	randData(l, r, &p)
	if flag {
		p.formulaTostring += ")"
		l.PushBack(&Entry{
			kind:  2,
			value: &Sign{s: ')'},
		})
	}
	p.formulaTostring += "="
	//l.PushBack(&Entry{
	//	kind:  2,
	//	value: &Sign{s: '='},
	//})
	//fmt.Println(p.formulaTostring)
	p.TransPostfixExpress()
	Divedzero = false
	ret := p.Cal()
	if _, ok := m[ret]; ok || ret.Num < 0 || ret.Nume < 0 || ret.Deno < 0 {
		return nil, false
	} else {
		m[ret] = 1
		p.answer = *ret
		return &p, true
	}
}

func ParseArgs() {
	flag.Int64Var(&r, "r", 0, "操作数的取值范围。<r")
	flag.IntVar(&n, "n", 0, "要生成的题目数量。")
	flag.StringVar(&exercisefile, "e", "", "学生提交的作业。")
	flag.StringVar(&answerfile, "a", "", "答案文件。")
	flag.Parse()
}

func Usage() {
	argNum := len(os.Args)
	if argNum < 5 {
		fmt.Print(
			`
用法:  生成算式:main  [-n count] [-r number] 
	   阅卷模式:main [-a answerfile]  [-e exercisefile]
选项:
    -n count       要生成的题目数量。
    -r number        操作数的取值范围。<r。
    -a answerfile     答案文件。
	-e exercisefile 	学生提交的作业。
。	
`)
	}
}

func readLine(filename string) ([]*Problem, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		//fmt.Println(err.Error())
		return nil, err
	}
	s := make([]*Problem, 0)
	for _, list := range strings.Split(string(data), "\n") {
		l := strings.Split(list, ".")
		p := NewProblem()
		p.formulaTostring = l[1]
		p.TransStringToFormula()
		p.TransPostfixExpress()
		//fmt.Println(p.formula.Len())
		//for ele := p.formula.Front(); ele != nil; ele = ele.Next() {
		//	fmt.Println(ele.Value.(*Entry).value.(*FAL).String())
		//}
		p.answer = *p.Cal()

		s = append(s, p)
	}
	return s, nil
}

func writeGrade(exercise []*Problem, answer []string) {
	var (
		correct = make([]int, 0)
		wrong   = make([]int, 0)
	)
	for index, v := range exercise {
		//fmt.Println(v.formulaTostring, v.answer)
		if answer[index] == v.answer.String() {
			correct = append(correct, index+1)
		} else {
			wrong = append(wrong, index+1)
		}
	}
	fff, _ := os.OpenFile("grade.txt",
		os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_TRUNC, 0777) //读写模式打开，写入追加
	defer fff.Close()
	fff.Write([]byte("Correct:" + strconv.Itoa(len(correct)) + "("))
	for i := 0; i < len(correct)-1; i++ {
		fff.Write([]byte(strconv.Itoa(correct[i]) + ","))
	}
	if len(correct) > 0 {
		fff.Write([]byte(strconv.Itoa(correct[len(correct)-1]) + ")\n\n"))
	} else {
		fff.Write([]byte(")\n\n"))
	}
	fff.Write([]byte("Wrong:" + strconv.Itoa(len(wrong)) + "("))
	for i := 0; i < len(wrong)-1; i++ {
		fff.Write([]byte(strconv.Itoa(wrong[i]) + ","))
	}
	if len(wrong) > 0 {
		fff.Write([]byte(strconv.Itoa(wrong[len(wrong)-1]) + ")"))
	} else {
		fff.Write([]byte(")"))
	}
}

func generateProblems() {
	f, _ := os.OpenFile("exercise.txt",
		os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_TRUNC, 0777) //读写模式打开，写入追加
	defer f.Close()
	ff, _ := os.OpenFile("answer.txt",
		os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_TRUNC, 0777) //读写模式打开，写入追加
	defer ff.Close()
	for i := 0; i < n; {
		if problem, ok := getProblem(); ok {
			//fmt.Println(problem.postfixExpress.Len())
			problemList = append(problemList, problem)
			i++
			//fmt.Println(problem.formulaTostring, problem.answer)
			f.Write([]byte(strconv.Itoa(i) + "." + problem.formulaTostring))
			ff.Write([]byte(strconv.Itoa(i) + "." + problem.formulaTostring + problem.answer.String()))
			if i < n {
				f.Write([]byte("\n"))
				ff.Write([]byte("\n"))
			}
		}
	}
}

func check() {

	exercise, err := readLine(exercisefile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, err := ioutil.ReadFile("answerfile.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	answer := make([]string, 0)
	for _, list := range strings.Split(string(data), "\n") {
		l := strings.Split(list, ".")
		answer = append(answer, l[1])
	}
	writeGrade(exercise, answer)

}
func main() {
	s := time.Now()
	rand.Seed(time.Now().UnixNano())
	len := len(os.Args)
	//fmt.Println(len)
	if len < 5 {
		Usage()
	}
	ParseArgs()
	if n > 0 {
		generateProblems()
	}
	if exercisefile != "" && answerfile != "" {
		check()
	}
	//if exercisefile != "" && answerfile != "" {
	//
	//	exercise, err := readLine(exercisefile)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//
	//	data, err := ioutil.ReadFile("answerfile.txt")
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	answer := make([]string, 0)
	//	for _, list := range strings.Split(string(data), "\n") {
	//		l := strings.Split(list, ".")
	//		answer = append(answer, l[1])
	//	}
	//	writeGrade(exercise, answer)
	//}
	fmt.Println(time.Since(s))
}
