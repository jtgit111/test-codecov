package testCodecov

import "testing"

func TestDiv(t *testing.T) {
    tests := []struct{a, b, expected int} {
        {2, 1, 2},
    }
    for _, test := range tests {
        res := Div(test.a, test.b)
        if res != test.expected {
            t.Errorf("Div(%d, %d) = %d; want %d", test.a, test.b, res, test.expected)
        }
    }
}

func TestDivThree(t *testing.T) {
    tests := []struct{a, b, c, expected int} {
        {8, 1, 2, 4},
    }
    for _, test := range tests {
        res := DivThree(test.a, test.b, test.c)
        if res != test.expected {
            t.Errorf("Div(%d, %d, %d) = %d; want %d", test.a, test.b, test.c, res, test.expected)
        }
    }
}

func TestDivFour(t *testing.T) {
    tests := []struct{a, b, c, d, expected int} {
        {8, 1, 2, 2, 2},
    }
    for _, test := range tests {
        res := DivFour1(test.a, test.b, test.c, test.d)
        if res != test.expected {
            t.Errorf("Div(%d, %d, %d, %d) = %d; want %d", test.a, test.b, test.c, test.d, res, test.expected)
        }
    }
}