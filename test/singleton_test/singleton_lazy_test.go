package singleton

import (
	"github.com/bmizerany/assert"
	"testing"
	"tsaiilin.com/singleton"
)

func TestLazySingleton(t *testing.T) {
	assert.Equal(t, singleton.GetInstanceLazy(), singleton.GetInstanceLazy())
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetInstanceLazy() != singleton.GetInstanceLazy() {
				b.Errorf("test fail")
			}
		}
	})
}