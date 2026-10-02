package dispatcher_test

import (
	"context"
	"errors"
	"testing"

	"fabriciolfj.github/study/internal/delivery"
	"fabriciolfj.github/study/internal/dispatcher"
)

type stubSender struct {
	sent []*delivery.Delivery
	err  error
}

func (s *stubSender) Send(_ context.Context, d *delivery.Delivery) error {
	s.sent = append(s.sent, d)
	return s.err
}

var _ delivery.Sender = (*stubSender)(nil)

func TestProcess(t *testing.T) {
	tests := []struct {
		name           string
		erroDoSender   error
		statusEsperado string
		tentativas     int
	}{
		{"sucesso na primeira", nil, "SEND", 0},
		{"falha incrementa", errors.New("timeout"), "", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stubSender{err: tt.erroDoSender}
			dp := dispatcher.New(s)
			d := delivery.New("m-1", "https://x", nil)

			err := dp.Process(context.Background(), d)

			if !errors.Is(err, tt.erroDoSender) {
				t.Errorf("Process() err = %v, want %v", err, tt.erroDoSender)
			}
			if d.Status != tt.statusEsperado {
				t.Errorf("Status = %q, want %q", d.Status, tt.statusEsperado)
			}
			if d.Attempts != tt.tentativas {
				t.Errorf("Attempts = %d, want %d", d.Attempts, tt.tentativas)
			}
			if len(s.sent) != 1 {
				t.Errorf("Send chamado %d vezes, want 1", len(s.sent))
			}
		})
	}
}
