package testCodecov

import "testing"

func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {1, 2, 3},
        {2, 3, 5},
        {-1, 1, 0},
        {0, 0, 0},
    }

    for _, test := range tests {
        result := Add(test.a, test.b)
        if result != test.expected {
            t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
        }
    }
}

func TestSub(t *testing.T) {
    tests := []struct{a, b, expected int} {
        {2, 1, 1},
    }

    for _, test := range tests {
        res := Sub(test.a, test.b)
        if res != test.expected {
            t.Errorf("Sub(%d, %d) = %d; want %d", test.a, test.b, res, test.expected)
        }
    }
}
