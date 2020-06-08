package dog

import "testing"

func TestYearsTwo(t *testing.T) {
	x := YearsTwo(4)
	if x != 28 {
		t.Error("Expected", 28, "found", x)
	}
}
func TestYears(t *testing.T) {
	type test struct {
		age int
		ans int
	}
	tests := []test{
		test{10, 70},
		test{20, 140},
		test{2, 14},
		test{5, 35},
	}
	for _, v := range tests {
		x := Years(v.age)
		if x != v.ans {
			t.Error("Expected", v.ans, "found", x)
		}
	}

}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
