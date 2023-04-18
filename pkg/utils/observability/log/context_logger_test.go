package log

import (
	"context"
	"nonoDemo/pkg/utils/observability"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func Test_ContextLogger(t *testing.T) {
	logfile := "./logs/ctx_logger.log"
	_ = os.Remove(logfile)
	baseLogger := NewLogger(Config{
		Writer: "os.File",
		Level:  "DEBUG",
		Path:   logfile,
		Format: "json",
	})
	ctx := context.Background()
	ctx = observability.AppendEvent(ctx, "mykey", "myvalue")
	ctxLogger := NewContextLogger(ctx, baseLogger)
	ctxLogger.Debug("_")
	ctxLogger.Info("_")
	ctxLogger.Warn("_", nil)
	ctxLogger.Error("_", nil)
	content, err := os.ReadFile(logfile)
	if err != nil {
		t.Fatal(err)
	}
	c := string(content)
	lines := strings.Split(c, "\n")
	if len(lines)-1 != 4 {
		t.Fatal("log count err")
	}
	log := make(map[string]interface{})
	_ = json.Unmarshal([]byte(lines[0]), &log)
	if log["level"].(string) != "DEBUG" || log["mykey"].(string) != "myvalue" {
		t.Fatal("DEBUG  log error")
	}
	_ = json.Unmarshal([]byte(lines[1]), &log)
	if log["level"].(string) != "INFO" || log["mykey"].(string) != "myvalue" {
		t.Fatal("INFO  log error")
	}
	_ = json.Unmarshal([]byte(lines[2]), &log)
	if log["level"].(string) != "WARN" || log["mykey"].(string) != "myvalue" {
		t.Fatal("WARN log error")
	}
	_ = json.Unmarshal([]byte(lines[3]), &log)
	if log["level"].(string) != "ERROR" || log["mykey"].(string) != "myvalue" {
		t.Fatal("ERROR log error")
	}
}
