// Package minimax maps MiniMax reasoning controls to the two supported wire protocols.
package minimax

import (
	"fmt"

	"moonbridge/internal/format"
)

// Reasoning returns wire thinking.type and effort for the exact supported models.
// Other models retain their existing adapter behavior. Native Core controls take
// precedence over Responses extensions. M3 thinking modes are not effort levels.
func Reasoning(req *format.CoreRequest, defaultThinking string) (thinking, effort string, err error) {
	if req.Model != "MiniMax-M3.1-Flash" && req.Model != "MiniMax-M3" {
		return "", "", nil
	}
	openai, _ := req.Extensions["openai"].(map[string]any)
	reasoning, _ := openai["reasoning"].(map[string]any)
	effort, _ = reasoning["effort"].(string)
	thinking, _ = reasoning["type"].(string)
	if req.Output != nil {
		effort = req.Output.Effort
	}
	if req.Thinking != nil {
		thinking = req.Thinking.Type
	}
	if req.Model == "MiniMax-M3" {
		// Accept previously advertised presets without sending them as effort.
		if thinking == "" && (effort == "adaptive" || effort == "disabled") {
			thinking = effort
		}
		if thinking == "" {
			thinking = defaultThinking
		}
		if thinking != "adaptive" && thinking != "disabled" {
			return "", "", fmt.Errorf("MiniMax-M3 thinking type must be adaptive or disabled")
		}
		return thinking, "", nil
	}
	if thinking != "" && thinking != "adaptive" {
		return "", "", fmt.Errorf("MiniMax-M3.1-Flash requires adaptive thinking")
	}
	if effort == "" {
		effort = "max"
	}
	switch effort {
	case "low", "medium", "high", "xhigh", "max":
		return "adaptive", effort, nil
	default:
		return "", "", fmt.Errorf("unsupported MiniMax-M3.1-Flash reasoning effort %q", effort)
	}
}
