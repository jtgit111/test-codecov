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

func TestAddThree(t *testing.T) {
    tests := []struct {
        a, b, c, expected int
    }{
        {1, 2, 3, 6},
    }

    for _, test := range tests {
        result := AddThree(test.a, test.b, test.c)
        if result != test.expected {
            t.Errorf("Add(%d, %d, %d) = %d; want %d", test.a, test.b, test.c, result, test.expected)
        }
    }
}

func TestAddFour(t *testing.T) {
    tests := []struct {
        a, b, c, d, expected int
    }{
        {1, 2, 3, 4, 10},
    }

    for _, test := range tests {
        result := AddFour(test.a, test.b, test.c, test.d)
        if result != test.expected {
            t.Errorf("Add(%d, %d, %d, %d) = %d; want %d", test.a, test.b, test.c, test.d, result, test.expected)
        }
    }
}