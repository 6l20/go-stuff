package main

import "testing"

func BenchmarkImmutablePerson(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = immutableCreatePerson()
    }
}

func BenchmarkMutablePerson(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = mutableCreatePerson()
    }
}