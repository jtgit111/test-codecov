package testCodecov

import "testing"

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

func TestSubThree(t *testing.T) {
    tests := []struct{a, b, c, expected int} {
        {4, 1, 1, 2},
    }

    for _, test := range tests {
        res := SubThree(test.a, test.b, test.c)
        if res != test.expected {
            t.Errorf("Sub(%d, %d, %d) = %d; want %d", test.a, test.b, test.c, res, test.expected)
        }
    }
}

func TestSubFour(t *testing.T) {
    tests := []struct{a, b, c, d, expected int} {
        {4, 1, 1, 2, 0},
    }

    for _, test := range tests {
        res := SubFour(test.a, test.b, test.c, test.d)
        if res != test.expected {
            t.Errorf("Sub(%d, %d, %d, %d) = %d; want %d", test.a, test.b, test.c, test.d, res, test.expected)
        }
    }
}