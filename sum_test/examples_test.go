package sum_test

import (
	"fmt"

	"github.com/ArmandoOC/unittesting2/sum"
)

func ExampleInts() {
	s := sum.Ints(1, 2, 3, 4, 5)
	fmt.Println("sum of one to five is", s)
}

//Los examples además de ser tests, la ventaja es
//que aparecen en la documentación:

//  godoc -http=:6060
