package main

import "testing"

func TestCalculateDelay(t *testing.T) {
	cases := []struct {
		name    string
		fps     int
		want    int
		wantErr bool
	}{
		{
			name:    "returns correct value with 1 fps",
			fps:     1,
			want:    100,
			wantErr: false,
		},
		{
			name:    "returns correct value with 7 fps",
			fps:     7,
			want:    14,
			wantErr: false,
		},
		{
			name:    "returns correct value with 10 fps",
			fps:     10,
			want:    10,
			wantErr: false,
		},
		{
			name:    "returns correct value with 10 fps",
			fps:     30,
			want:    3,
			wantErr: false,
		},
		{
			name:    "fails with negative fps",
			fps:     -10,
			want:    0,
			wantErr: true,
		},
		{
			name:    "fails with 0 fps",
			fps:     0,
			want:    0,
			wantErr: true,
		},
		{
			name:    "fails with fps greater than maximum allowed",
			fps:     35,
			want:    0,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(s *testing.T) {
			got, err := calculateDelay(c.fps)

			if c.want != got {
				t.Errorf("calculateDelay(%d) = %d; want %d", c.fps, got, c.want)
			}

			gotErr := err != nil
			if c.wantErr != gotErr {
				t.Errorf("expected error with calculateDelay(%d) = %v; got error = %v", c.fps, c.wantErr, gotErr)
			}
		})
	}
}
