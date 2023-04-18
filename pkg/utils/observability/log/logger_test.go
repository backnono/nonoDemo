package log

import (
	"io"
	"io/ioutil"
	"os"
	path2 "path"
	"strings"
	"testing"
	"time"
)

func TestLogger_FileWriter(t *testing.T) {
	cfg := Config{
		Writer: "os.File",
		Level:  "debug",
		Path:   "./logs/test.log",
		Format: "text",
	}
	logger := NewLogger(cfg)
	logger.Info("test")
	f, err := os.Open(cfg.Path)
	if err != nil {
		t.Error(err)
	}
	content, err := io.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(content), "test") {
		t.Error("log content is not correct")
	}
	_ = f.Close()
	_ = os.Remove(cfg.Path)
}

func TestLogger_FileWriter_WithDate(t *testing.T) {
	cfg := Config{
		Writer: "os.File",
		Level:  "debug",
		Path:   "./logs_date/",
		Format: "text",
	}
	logger := NewLogger(cfg)
	logger.Info("test", "a", "abc")
	f, err := os.Open(path2.Join(cfg.Path, time.Now().Format("2006-01-02")+".log"))
	if err != nil {
		t.Error(err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(content), "test") {
		t.Error("log content is not correct")
	}
	_ = os.Remove(path2.Join(cfg.Path, time.Now().Format("2006-01-02")+".log"))
}
