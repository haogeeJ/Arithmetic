package main

func fakefenshu(v *FAL) {
	v.Nume += v.Num * v.Deno
	v.Num = 0
}

func realfenshu(v *FAL) {
	if v.Deno == 0 {
		Divedzero = true
		return
	}
	v.Num = v.Nume / v.Deno
	v.Nume %= v.Deno
}

func Add(v1 *FAL, v2 *FAL) *FAL {
	fakefenshu(v1)
	fakefenshu(v2)
	//fmt.Println(v1.Num, v1.Nume, v1.Deno)
	v1.Add(v2)
	realfenshu(v1)
	return v1
}

func Sub(v1 *FAL, v2 *FAL) *FAL {
	fakefenshu(v1)
	fakefenshu(v2)
	v1.Sub(v2)
	realfenshu(v1)
	return v1
}

func Mul(v1 *FAL, v2 *FAL) *FAL {
	fakefenshu(v1)
	fakefenshu(v2)
	v1.Mul(v2)
	realfenshu(v1)
	return v1
}

func Div(v1 *FAL, v2 *FAL) *FAL {
	fakefenshu(v1)
	fakefenshu(v2)
	v1.Div(v2)
	realfenshu(v1)
	return v1
}
