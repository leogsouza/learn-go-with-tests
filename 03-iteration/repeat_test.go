package iteration

import (
    "testing"
    "fmt"
)


func TestRepeat(t *testing.T) {
    
    repeated := Repeat("a", 3)
    expected := "aaa"
    
    if (repeated != expected) {
        t.Errorf("expected '%s' got '%s", expected, repeated)
    }
}

func BenchmarkRepeat(b *testing.B) {
    for i:= 0; i < b.N; i++ {
        Repeat("a", 3)
    }
}

func ExampleRepeat() {
    repeat := Repeat("la", 3)
    
    fmt.Println(repeat)
    // Output: lalala
    
}