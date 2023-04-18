package observability

import (
	"context"
	"testing"
)

func TestAppendEvents(t *testing.T) {
	_ = AppendEvent(context.Background(), "key", "value")
}
