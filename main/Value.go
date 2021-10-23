package main

import "fmt"

type Value interface {
	Len() int
}

type FAL struct {
	Num  int64
	Nume int64
	Deno int64
}

type PositiveInt uint

type Sign struct {
	s int32
}

func (p PositiveInt) Len() int {
	return 0
}
func (s FAL) Len() int {
	return 0
}
func (s Sign) Len() int {
	return 0
}

// Format output
func (s FAL) String() string { // 格式化输出
	if s.Nume == 0 {
		return fmt.Sprintf("%v", s.Num)
	}
	if s.Num == 0 {
		return fmt.Sprintf("%v/%v", s.Nume, s.Deno)
	}
	return fmt.Sprintf("%v'%v/%v", s.Num, s.Nume, s.Deno)
}

// Model Create a score (molecular, denominator) with a denominator default of 1
func Model(nd ...int64) *FAL { // 创建一个分数(分子，分母)，分母默认为1
	var f *FAL
	f = new(FAL)
	//fmt.Println("after:", nd)
	f.Num = nd[0]
	f.Nume = nd[1]
	f.Deno = nd[2]

	return f
}

// Broadsheet  .阔张
func (s *FAL) broad(lcm int64) {
	if s.Deno <= 0 {
		Divedzero = true
		return
	}
	s.Nume = s.Nume * (lcm / s.Deno)
	s.Deno = lcm
}

// Compression Finishing .压缩 整理
func (s *FAL) offset() {
	if s.Nume == 0 {
		s.Deno = 1
		return
	}
	lcm := Gcd(s.Nume, s.Deno)
	s.Nume /= lcm
	s.Deno /= lcm
}

// Add Fraction addition
func (s *FAL) Add(f *FAL) *FAL {
	// Getting the Minimum Common Multiplier 获取最小公倍数
	lcm := Lcm(f.Deno, s.Deno)
	s.broad(lcm)
	f.broad(lcm)

	s.Nume += f.Nume
	s.offset()
	return s
}

// Sub fraction subtraction
func (s *FAL) Sub(f *FAL) *FAL {
	// Getting the Minimum Common Multiplier 获取最小公倍数
	lcm := Lcm(s.Deno, f.Deno)
	s.broad(lcm)
	f.broad(lcm)

	s.Nume -= f.Nume
	s.offset()
	return s
}

// Mul multiplication
func (s *FAL) Mul(f *FAL) *FAL { // 乘法
	s.Deno *= f.Deno
	s.Nume *= f.Nume
	s.offset()
	return s
}

// Div division
func (s *FAL) Div(f *FAL) *FAL { // 除法
	s.Mul(Model(f.Num, f.Deno, f.Nume))
	s.offset()
	return s
}

func Gcd(x int64, y int64) int64 {
	if y == 0 {
		return x
	} else {
		return Gcd(y, x%y)
	}
}

func Lcm(x int64, y int64) int64 {
	return x / Gcd(x, y) * y
}
