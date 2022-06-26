package coverage

import (
	"os"
	"reflect"
	"testing"
	"time"
)

type (
	LessInputTest struct {
		people   People
		expected bool
	}
	MatrixInputTest struct {
		matrix   string
		expected Matrix
		err      string
	}
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
func TestPeople_Swap(t *testing.T) {
	people := People{{firstName: "John"}, {firstName: "Peter"}}
	peopleExpected := []string{"Peter", "John"}

	people.Swap(0, 1)
	if peopleExpected[0] != people[0].firstName && peopleExpected[1] != people[1].firstName {
		t.Errorf("func Swap(i, j int) does not work correctly")
	}
}

func TestPeople_Len(t *testing.T) {
	people := People{{firstName: "John"}, {firstName: "Peter"}}

	peopleLen := people.Len()
	if peopleLen != 2 {
		t.Errorf("func Len() does not work correctly")
	}
}

func TestPeople_Less(t *testing.T) {
	data := map[string]LessInputTest{
		"different people": {
			people: People{
				{
					firstName: "John",
					lastName:  "Black",
					birthDay:  time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Mary",
					lastName:  "Jane",
					birthDay:  time.Date(1993, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: true,
		},
		"equal birthday": {
			people: People{
				{
					firstName: "John",
					lastName:  "Black",
					birthDay:  time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Mary",
					lastName:  "Jane",
					birthDay:  time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: true,
		},
		"equal birthday and firstname": {
			people: People{
				{
					firstName: "Mary",
					lastName:  "Black",
					birthDay:  time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Mary",
					lastName:  "Jane",
					birthDay:  time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: false,
		},
		"equal everything": {
			people: People{
				{
					firstName: "Mary",
					lastName:  "Jane",
					birthDay:  time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Mary",
					lastName:  "Jane",
					birthDay:  time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: false,
		},
	}

	for name, input := range data {
		res := input.people.Less(0, 1)
		if res != input.expected {
			t.Errorf("%s: Actual result %t does not match the expected %t", name, res, input.expected)
		}
	}
}

func TestMatrix_Set(t *testing.T) {
	matrix, _ := New("1 2\n3 4")
	matrix.Set(1, 1, 100)
	expected := []int{1, 2, 3, 100}

	if reflect.DeepEqual(matrix.data, expected) == false {
		t.Errorf("Actual result %v does not match the expected %v", matrix.data, expected)
	}
}

func TestMatrix_Cols(t *testing.T) {
	matrix, _ := New("1 2 3\n4 5 6")
	expected := [][]int{{1, 4}, {2, 5}, {3, 6}}
	cols := matrix.Cols()

	if reflect.DeepEqual(cols, expected) == false {
		t.Errorf("Actual result %v does not match the expected %v", cols, expected)
	}
}

func TestMatrix_Rows(t *testing.T) {
	matrix, _ := New("1 2 3\n4 5 6")
	expected := [][]int{{1, 2, 3}, {4, 5, 6}}
	rows := matrix.Rows()

	if reflect.DeepEqual(rows, expected) == false {
		t.Errorf("Actual result %v does not match the expected %v", rows, expected)
	}
}

func TestNew(t *testing.T) {
	data := map[string]MatrixInputTest{
		"valid matrix": {
			matrix:   "1 2 3\n4 5 6",
			expected: Matrix{rows: 2, cols: 3, data: []int{1, 2, 3, 4, 5, 6}},
			err:      "",
		},
		"invalid matrix": {
			matrix: "1 2 3\n1",
			err:    "Rows need to be the same length",
		},
		"invalid elements": {
			matrix: "1 2 3\n3 2 b",
			err:    "strconv.Atoi: parsing \"b\": invalid syntax",
		},
	}

	for name, input := range data {
		mtrx, errs := New(input.matrix)
		if input.err != "" {
			if errs.Error() != input.err {
				t.Errorf("%s: error '%s' does not match the expected '%s'", name, errs, input.err)
			}
		} else {
			if mtrx.rows != input.expected.rows || mtrx.cols != input.expected.cols {
				t.Errorf("%s: actual matrix %v does not match expected matrix %v", name, mtrx, input.expected)
			}
		}
	}
}
