package service

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"
)

func TestConfigLogsAsOneGroup(t *testing.T) {
	var buf bytes.Buffer
	NewLogger(&buf, slog.LevelInfo).Info("config", "config", Config{Addr: "127.0.0.1:9000", LogLevel: slog.LevelWarn})

	var record struct {
		Config map[string]string `json:"config"`
	}
	if err := json.Unmarshal(buf.Bytes(), &record); err != nil {
		t.Fatalf("not JSON: %v: %q", err, buf.String())
	}
	want := map[string]string{"addr": "127.0.0.1:9000", "log_level": "ERROR"}
	if len(record.Config) != len(want) {
		t.Fatalf("config = %v, want %v", record.Config, want)
	}
	for k, v := range want {
		if record.Config[k] != v {
			t.Errorf("config[%q] = %q, want %q", k, record.Config[k], v)
		}
	}
}
