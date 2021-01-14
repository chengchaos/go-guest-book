package learning

import "testing"

func TestAdder(t *testing.T) {

	a := Add1()

	for i := 1; i <= 100; i++ {
		t.Logf("1 + 2 + ... + %d => %d\n",
			i,
			a(i))
	}
}

func TestAdd2(t *testing.T) {

	a := Add2(0)

	for i := 1; i <= 10; i++ {
		var s int
		s, a = a(i)

		t.Logf("1 + .. + %d => %d\n",
			i, s)
	}
}
