package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	g := Years(10)
	if g != 70 {
		t.Error("get", g, "want", 70)
	}
}

func TestYearTwo(t *testing.T) {
	g := YearsTwo(10)
	if g != 70 {
		t.Error("get", g, "want", 70)
	}
}

func ExampleYears() {
	y := 5
	dy := Years(y)
	fmt.Println(dy)
	// Output:
	// 35
}

func ExampleYearsTwo() {
	y := 7
	dy := YearsTwo(y)
	fmt.Println(dy)
	// Output:
	// 49
}

func BenchmarkYears(b *testing.B) {
	y := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Years(y)
	}
}

func BenchmarkYearTwo(b *testing.B) {
	y := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Years(y)
	}
}
