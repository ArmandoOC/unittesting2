package sum

import (
	"testing"
)

// func TestInts(t *testing.T) {
// 	//t.Errorf("this test failed")
// 	//t.Fatalf("this test failed and stopped the execution")

// 	s := Ints(1, 2, 3, 4, 5)
// 	if s != 15 {
// 		t.Errorf("sum of one to five should be 15; got %v", s)
// 	}

// 	s = Ints()
// 	if s != 0 {
// 		t.Errorf("sum of no numbers should be 0; got %v", s)
// 	}

// 	s = Ints(1, -1)
// 	if s != 0 {
// 		t.Errorf("sum of one plus minus one should be 0; got %v", s)
// 	}
// }

func TestInts(t *testing.T) {
	//table of tests
	tt := []struct {
		testName string
		numbers  []int
		sum      int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"no numbers", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}

	//Loop over the table of tests tt.   tc is test case
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			s := Ints(tc.numbers...)
			if s != tc.sum {
				t.Fatalf("sum of %v should be %v; got %v", tc.testName, tc.sum, s)
			}
		})

	}

}

// func TestFoo(t *testing.T) {
// 	t.Errorf("Fails")
// }

//go test
//go test -v
//go test -run Ints        con esto eliges qué test correr. Obvio, lo podemos hacer con el test explorer.
//go test -run Ints -v

//Run subtest
//  go test -v -run Ints/one     Con eso corren los que lleven la palabra one
//  go test -v -run Ints/numbers     Con eso corren los que lleven la palabra numbers
