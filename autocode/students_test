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
	firstPerson := Person{firstName: "Julia", lastName: "Kirillova", birthDay: time.Now()}
	secondPerson := Person{firstName: "Test", lastName: "Test", birthDay: time.Now()}
	people := People{firstPerson, secondPerson}

	people.Swap(0, 1)

	if people[0].firstName != "Test" || people[1].firstName != "Julia" {
		t.Errorf("func Swap() does not work correctly")
	}
}

func TestPeople_Swap2(t *testing.T) {
	person := Person{
		firstName: "Sara",
		lastName:  "Connor",
		birthDay:  time.Time{},
	}
	people := People{person, {}}

	people.Swap(0, 1)

	if people[0].firstName != "" || people[1].firstName != "Sara" {
		t.Errorf("func Swap() does not work correctly")
	}
}

func TestPeople_Len(t *testing.T) {
	people := People{{firstName: "John"}, {firstName: "Peter"}}

	peopleLen := people.Len()
	if peopleLen != 2 {
		t.Errorf("func Len() does not work correctly")
	}
}

func TestPeople_Len2(t *testing.T) {
	people := People{}

	peopleLen := people.Len()
	if peopleLen != 0 {
		t.Errorf("func Len() does not work correctly")
	}
}

func TestPeople_Less(t *testing.T) {
	data := map[string]LessInputTest{
		"different people": {
			people: People{
				{
					firstName: "Henry",
					lastName:  "White",
					birthDay:  time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Mary",
					lastName:  "Tate",
					birthDay:  time.Date(1985, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: true,
		},
		"equal firstname": {
			people: People{
				{
					firstName: "Henry",
					lastName:  "White",
					birthDay:  time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Henry",
					lastName:  "Tate",
					birthDay:  time.Date(1985, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: true,
		},
		"equal birthday": {
			people: People{
				{
					firstName: "Henry",
					lastName:  "White",
					birthDay:  time.Date(1985, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Mary",
					lastName:  "Tate",
					birthDay:  time.Date(1985, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: true,
		},
		"equal birthday and firstname": {
			people: People{
				{
					firstName: "Henry",
					lastName:  "Black",
					birthDay:  time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Henry",
					lastName:  "White",
					birthDay:  time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: true,
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

func TestMatrix_Set2(t *testing.T) {
	matrix, _ := New("1 2 3\n 2 3 4 \n 5 6 7")
	res := matrix.Set(2, -1, 4)

	if res != false {
		t.Errorf("Actual result %t does not match the expected %t", res, false)
	}
}

func TestMatrix_Set3(t *testing.T) {
	matrix, _ := New("1 2 3\n 2 3 4")
	res := matrix.Set(3, 2, 4)

	if res != false {
		t.Errorf("Actual result %t does not match the expected %t", res, false)
	}
}

func TestMatrix_Set4(t *testing.T) {
	matrix, _ := New("1 2 3\n 2 3 4 \n 5 6 7")
	res := matrix.Set(1, 4, 4)

	if res != false {
		t.Errorf("Actual result %t does not match the expected %t", res, false)
	}
}

func TestMatrix_Set5(t *testing.T) {
	matrix, _ := New("1 2 3\n 2 3 4 \n 5 6 7")
	res := matrix.Set(2, 2, 4)

	if res != true {
		t.Errorf("Actual result %t does not match the expected %t", res, true)
	}
}

func TestMatrix_Set6(t *testing.T) {
	matrix, _ := New("1 2 3\n 2 3 4 \n 5 6 7")
	res := matrix.Set(-1, 1, 4)

	if res != false {
		t.Errorf("Actual result %t does not match the expected %t", res, false)
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

func TestMatrix_Rows2(t *testing.T) {
	matrix, _ := New("1 2\n4 5")
	expected := [][]int{{1, 2}, {4, 5}}
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
		"empty matrix": {
			matrix: "",
			err:    "strconv.Atoi: parsing \"\": invalid syntax",
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
