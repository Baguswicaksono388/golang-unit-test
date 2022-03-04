package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Bagus",
			request: "Bagus",
		},
		{
			name:    "Wicaksono",
			request: "Wicaksono",
		},
		{
			name:    "BagusWicaksono",
			request: "Bagus Wicaksono",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Bagus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bagus")
		}
	})

	b.Run("Wicaksono", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Wicaksono")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bagus")
	}
}

func BenchmarkHelloWorldWicaksono(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Wicaksono")
	}
}

// Konsep table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Bagus",
			request:  "Bagus",
			expected: "Hello Bagus",
		},
		{
			name:     "Wicaksono",
			request:  "Wicaksono",
			expected: "Hello Wicaksono",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before Unit Test")

	m.Run()

	// after
	fmt.Println("After Unit Test")
}

func TestSubTest(t *testing.T) {
	t.Run("Bagus", func(t *testing.T) {
		result := HelloWorld("Bagus")
		require.Equal(t, "Hello Bagus", result, "Result must be 'Hello Bagus")
	})
	t.Run("Wicaksono", func(t *testing.T) {
		result := HelloWorld("Wicaksono")
		require.Equal(t, "Hello Wicaksono", result, "Result must be 'Hello Wicaksono")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" { //darwin for Mac OS
		t.Skip("Cannot run on Windows")
	}

	result := HelloWorld("Bagus")
	require.Equal(t, "Hello Bagus", result, "Result must be 'Hello Bagus")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bagus")
	assert.Equal(t, "Hello Bagus", result, "Result must be 'Hello Bagus")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bagus")
	require.Equal(t, "Hello Bagus", result, "Result must be 'Hello Bagus")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bagus")

	if result != "Hello Bagus" {
		// error
		// panic("Result is not 'Hello Bagus'")
		// t.Fail()
		t.Error("Result must be 'Hello Bagus'")
	}
	fmt.Println("TestHelloWorld done")
}

func TestHelloBagus(t *testing.T) {
	result := HelloWorld("Bagus")

	if result != "Hello Bagus" {
		// t.FailNow()
		t.Fatal("Result must be 'Hello Bagus'")
	}

	fmt.Println("TestHelloBagus done")
}
