package reader

import (
	"testing"
)

func Test_ReadPlan(t *testing.T) {
	data := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "invalid extension",
			input:   "plan.txt",
			wantErr: true,
		},
		{
			name:    "file not found",
			input:   "nonexistent.json",
			wantErr: true,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			_, err := ReadPlan(d.input)

			if (err != nil) != d.wantErr {
				t.Fatalf("got err=%v, wantErr=%v", err, d.wantErr)
			}
		})
	}
}
