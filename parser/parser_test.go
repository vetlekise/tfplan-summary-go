package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ParseChanges(t *testing.T) {
	// TODO: Test JSON unmarshaling
	// TODO: Test 'no-op' filtering
	// TODO: Test 'replace' normalization
	// TODO: Test sorting

	data := []struct {
		name  string
		input []byte
		want  []ResourceDiff
	}{
		{
			name:  "filters no-op",
			input: []byte(`{"resource_changes":[{"address":"resource.1","change":{"actions":["no-op"]}}]}`),
			want:  nil,
		},
	}

	for _, d := range data {

		t.Run(d.name, func(t *testing.T) {
			got, err := ParseChanges(d.input)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if diff := cmp.Diff(d.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
