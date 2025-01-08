package testCodecov

import "testing"

func TestMul(t *testing.T) {
    tests := []struct{a, b, expected int} {
        {2, 1, 2},
    }
    for _, test := range tests {
        res := Mul(test.a, test.b)
        if res != test.expected {
            t.Errorf("Mul(%d, %d) = %d; want %d", test.a, test.b, res, test.expected)
        }
    }
}