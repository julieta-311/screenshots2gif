package animate

import (
	"context"
	"testing"
)

func TestAnimate(t *testing.T) {
	cases := []struct {
		name string
		loop bool
		fps  int
	}{
		{
			name: "succeeds with single loop",
			loop: false,
			fps:  4,
		},
		{
			name: "succeeds with infinite loop",
			loop: true,
			fps:  2,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(s *testing.T) {
			err := Animate(context.TODO(), "testdata", c.loop, c.fps)
			if err != nil {
				t.Errorf("got unexpected error: %v", err)
			}
		})
	}
}
