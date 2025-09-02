package util

import "testing"

func TestCopyMap(t *testing.T) {
	dst := make(map[string]any)
	src := make(map[string]any)

	t.Run("case", func(t *testing.T) {
		src["key"] = "value"
		CopyMap(dst, src)
		if dst["key"] != "value" {
			t.Errorf("dst key not copied")
		}
	})

}
