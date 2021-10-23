package main

import (
	"container/list"
)

type Entry struct {
	kind  int
	value Value
}

func (e Entry) Len() int {
	return 0
}

type Problem struct {
	formula         *list.List
	postfixExpress  *list.List
	formulaTostring string
	answer          FAL
}

//func (p *Problem) Add(value Value, kind int) {
//	p.formula.PushFront(&Entry{
//		kind:  kind,
//		value: value,
//	})
//}

func (p *Problem) TransPostfixExpress() {
	var ll *list.List
	ll = new(list.List)
	for ele := p.formula.Front(); ele != nil; ele = ele.Next() {
		v := ele.Value.(*Entry)
		if v.kind == 1 {
			p.postfixExpress.PushBack(v)
		} else {
			symbol := v.value.(*Sign)
			if ll.Len() == 0 || symbol.s == '(' {
				ll.PushBack(v)
				continue
			}
			if symbol.s == ')' {
				for val := ll.Back(); val != nil; val = ll.Back() {
					x := val.Value.(*Entry).value.(*Sign)
					if x.s == '(' {
						ll.Remove(val)
						break
					}
					p.postfixExpress.PushBack(val.Value.(*Entry))
					ll.Remove(val)
				}
				continue
			}
			x := ll.Back().Value.(*Entry).value.(*Sign)
			if x.s == '(' {
				ll.PushBack(v)
				continue
			}
			if symbol.s == '-' || symbol.s == '+' {
				for val := ll.Back(); val != nil; val = ll.Back() {
					x := val.Value.(*Entry).value.(*Sign)
					if x.s != '(' {
						p.postfixExpress.PushBack(val.Value.(*Entry))
						ll.Remove(val)
					} else {
						ll.PushBack(v)
						break
					}
				}
				if ll.Len() == 0 {
					ll.PushBack(v)
				}
				continue
			}
			if symbol.s == '*' || symbol.s == '÷' {
				for val := ll.Back(); val != nil; val = ll.Back() {
					x := val.Value.(*Entry).value.(*Sign)
					if x.s == '÷' || x.s == '*' {
						p.postfixExpress.PushBack(val.Value.(*Entry))
						ll.Remove(val)
					} else {
						ll.PushBack(v)
						break
					}
				}
				if ll.Len() == 0 {
					ll.PushBack(v)
				}
			}
		}
	}
	for val := ll.Back(); val != nil; val = ll.Back() {
		p.postfixExpress.PushBack(val.Value.(*Entry))
		ll.Remove(val)
	}
}

func (p *Problem) Cal() *FAL {
	var ll *list.List
	//cnt := 0
	ll = new(list.List)
	for ele := p.postfixExpress.Front(); ele != nil; ele = ele.Next() {
		//cnt++
		kv := ele.Value.(*Entry)
		if kv.kind == 1 {
			ll.PushBack(kv)
		} else {
			//fmt.Println(ll.Len(), kv.value.(*Sign).s)
			ele1 := ll.Back()
			ll.Remove(ele1)
			ele2 := ll.Back()
			ll.Remove(ele2)
			kv1 := ele1.Value.(*Entry)
			kv2 := ele2.Value.(*Entry)
			fal1 := kv1.value.(*FAL)
			fal2 := kv2.value.(*FAL)
			sign := kv.value.(*Sign)
			//fmt.Println(fal1, fal2)
			switch sign.s {
			case '-':
				fal2 = Sub(fal2, fal1)
			case '+':
				fal2 = Add(fal2, fal1)
			case '*':
				fal2 = Mul(fal2, fal1)
			case '÷':
				fal2 = Div(fal2, fal1)
			}
			//fmt.Println(fal2)
			ll.PushBack(&Entry{
				kind:  1,
				value: fal2,
			})
		}
	}
	//fmt.Println(cnt)
	return ll.Back().Value.(*Entry).value.(*FAL)
}
