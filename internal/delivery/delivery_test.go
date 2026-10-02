package delivery_test

import (
	"errors"
	"testing"

	"fabriciolfj.github/study/internal/delivery"
)

func TestIsRetryable(t *testing.T) {
	tests := []struct {
		name     string
		falhas   int
		esperado bool
	}{
		{"nunca tentou", 0, true},
		{"quarta falha ainda tenta", 4, true},
		{"quinta falha esgota", 5, false},
		{"acima do limite", 9, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := delivery.New("m-1", "https://x", nil)

			for i := 0; i < tt.falhas; i++ {
				d.Fail(errors.New(tt.name))
			}

			if got := d.IsRetryable(); got != tt.esperado {
				t.Errorf("isRetryable() = %v, want %v (attempts=%d)", got, tt.esperado, tt.falhas)
			}
		})
	}
}
