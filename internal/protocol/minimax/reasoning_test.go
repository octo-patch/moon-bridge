package minimax

import (
	"testing"

	"moonbridge/internal/format"
)

func TestReasoningMapsFlashToAdaptiveThinkingAndEffort(t *testing.T) {
	req := &format.CoreRequest{
		Model: "MiniMax-M3.1-Flash",
		Extensions: map[string]any{
			"openai": map[string]any{
				"reasoning": map[string]any{"effort": "high"},
			},
		},
	}
	thinking, effort, err := Reasoning(req, "adaptive")
	if err != nil {
		t.Fatal(err)
	}
	if thinking != "adaptive" || effort != "high" {
		t.Fatalf("got thinking=%q effort=%q", thinking, effort)
	}
}

func TestReasoningMapsM3ModesToThinkingType(t *testing.T) {
	for _, mode := range []string{"adaptive", "disabled"} {
		req := &format.CoreRequest{
			Model: "MiniMax-M3",
			Extensions: map[string]any{
				"openai": map[string]any{
					"reasoning": map[string]any{"effort": mode},
				},
			},
		}
		thinking, effort, err := Reasoning(req, "disabled")
		if err != nil {
			t.Fatal(err)
		}
		if thinking != mode || effort != "" {
			t.Fatalf("mode %q mapped to thinking=%q effort=%q", mode, thinking, effort)
		}
	}
}

func TestReasoningUsesNativeOutputAndThinkingControls(t *testing.T) {
	req := &format.CoreRequest{
		Model:    "MiniMax-M3.1-Flash",
		Thinking: &format.CoreThinkingConfig{Type: "adaptive"},
		Output:   &format.CoreOutputConfig{Effort: "low"},
	}
	thinking, effort, err := Reasoning(req, "adaptive")
	if err != nil {
		t.Fatal(err)
	}
	if thinking != "adaptive" || effort != "low" {
		t.Fatalf("got thinking=%q effort=%q", thinking, effort)
	}
}
