package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ParseChanges(t *testing.T) {
	data := []struct {
		name    string
		input   []byte
		want    []ResourceDiff
		wantErr bool
	}{
		{
			name:  "filters no-op",
			input: []byte(`{"resource_changes":[{"address":"resource.1","change":{"actions":["no-op"]}}]}`),
			want:  nil,
		},
		{
			name:  "replace normalization",
			input: []byte(`{"resource_changes":[{"address":"resource.something['replace1']","change":{"actions":["create","delete"]}},{"address":"resource.something['replace2']","change":{"actions":["delete","create"]}}]}`),
			want: []ResourceDiff{
				{Action: "replace", Address: "resource.something['replace1']"},
				{Action: "replace", Address: "resource.something['replace2']"},
			},
		},
		{
			name:    "invalid json",
			input:   []byte(`not json`),
			want:    nil,
			wantErr: true,
		},
		{
			name:  "empty plan",
			input: []byte(`{"resource_changes":[]}`),
			want:  nil,
		},
		{
			name:  "sorting",
			input: []byte(`{"resource_changes":[{"address":"resource.1","change":{"actions":["create"]}},{"address":"resource.2","change":{"actions":["delete"]}},{"address":"resource.3","change":{"actions":["update"]}}]}`),
			want: []ResourceDiff{
				{Action: "update", Address: "resource.3"},
				{Action: "delete", Address: "resource.2"},
				{Action: "create", Address: "resource.1"},
			},
		},
	}

	for _, d := range data {

		t.Run(d.name, func(t *testing.T) {
			got, err := ParseChanges(d.input)

			if (err != nil) != d.wantErr {
				t.Fatalf("got err=%v, wantErr=%v", err, d.wantErr)
			}

			if !d.wantErr {
				if diff := cmp.Diff(d.want, got); diff != "" {
					t.Errorf("mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
