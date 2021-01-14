package learning

import (
	"bufio"
	"fmt"
	"os"
)

func TryDefer() {

	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	fmt.Println(4)
}

func Fibonacci() func() int {
	//return func() int {
	//    a, b = b, a + b
	//    return a
	//}
	return func() int {
		return 1
	}
}

func WriteFile(filename string) {

	defer func() {

		r := recover()
		if r != nil {
			if err, ok := r.(error); ok {
				fmt.Printf("Error occurred: %v\n", err)
			} else {
				panic(r)
			}
		}
	}()

	file, err := os.Create(filename)
	if err != nil {
		//panic(err)
		fmt.Println("Error =>:", err.Error())

		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}

		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, i)
	}
}
