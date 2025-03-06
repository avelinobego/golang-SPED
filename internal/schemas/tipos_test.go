package schemas_test

import (
	"sped/internal/schemas"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	t.Log(schemas.Periodo(time.Now(), time.Now()))
}
