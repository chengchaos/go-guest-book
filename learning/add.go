package learning

func Add1() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type IAdd func(int) (int, IAdd)

func Add2(base int) IAdd {
	return func(v int) (int, IAdd) {
		return base + v, Add2(base + v)
	}
}
