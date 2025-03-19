package types_test

import (
	"sped/internal/types"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	t.Log(types.Periodo(time.Now(), time.Now()))
}
