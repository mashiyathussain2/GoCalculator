package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	value := Add(4, 12)
	got := 10
	if value != got {
		t.Error("Expected value:", value, "\nValue got: ", got)
	}
}

func TestSubtract(t *testing.T) {
	value := Subtract(4, 6)
	got := -2
	if value != got {
		t.Error("Expected value:", value, "\nValue got: ", got)
	}
}

func TestMultiply(t *testing.T) {
	value := Multiply(4, 6)
	got := 24
	if value != got {
		t.Error("Expected value:", value, "\nValue got: ", got)
	}
}

func TestDivision(t *testing.T) {
	value := Division(-100, 10)
	got := 50
	if value != got {
		t.Error("Expected value:", value, "\nValue got: ", got)
	}
}

func TestSquare(t *testing.T) {
	value := Square(-10.44)
	got := float32(-100)
	if value != got {
		t.Error("Expected value:", value, "\nValue got: ", got)
	}
}

func TestAddWithTestify(t *testing.T) {
	x := 1
	y := 2
	value := 4
	got := Add(x, y)
	assert.Equal(t, got, value)
}

func TestSum(t *testing.T) {
	value := []struct {
		a, b, result int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{1, -1, 0},
		{2, 3, 5},
		{1000, 1000, 2000},
	}
	for _, v := range value {
		if got := Add(v.a, v.b); got != v.result {
			t.Errorf("Addition of (%d, %d) is %d, but the value is %d", v.a, v.b, got, v.result)
		}
	}
}

func TestName(t *testing.T) {
	value := Name("Mashiyat Hussain")
	got := "Mashiyat Hussain"
	if value != got {
		t.Error("Expected value:", value, "\nValue got: ", got)
	}
}

func TestNameAnother(t *testing.T) {
	value := []struct {
		a, result string
	}{
		{"Mashiyat", "Mashiyat"},
		{"Hussain", "Hussain"},
	}
	for _, v := range value {
		if got := Name(v.a); got != v.result {
			t.Errorf(v.a, got, v.result)
		}
	}
}

func TestPassword(t *testing.T) {
	value := []struct {
		a, result string
	}{
		{"12345", "12345"},
		{"Hussain", "Hussain"},
	}
	for _, v := range value {
		if got := Name(v.a); got != v.result {
			t.Errorf(v.a, got, v.result)
		}
	}
}
