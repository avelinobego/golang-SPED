package value_test

import (
	"sped/internal/value"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	t.Log(value.Periodo(time.Now(), time.Now()))
}
