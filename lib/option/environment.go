package option

import (
	"context"
	"time"

	"github.com/mithrandie/go-text/color"
)

const (
	XDGConfigHomeEnvName   = "XDG_CONFIG_HOME"
	DefaultXDGConfigDir    = ".config"
	CSVQConfigDir          = "csvq"
	EnvFileName            = "csvq_env.json"
	PreloadCommandFileName = "csvqrc"

	HiddenPrefix = '.'
)

type Environment struct {
	DatetimeFormat       []string            `json:"datetime_format"`
	Timezone             *string             `json:"timezone"`
	AnsiQuotes           *bool               `json:"ansi_quotes"`
	InteractiveShell     InteractiveShell    `json:"interactive_shell"`
	EnvironmentVariables map[string]string   `json:"environment_variables"`
	Palette              color.PaletteConfig `json:"palette"`
}

func NewEnvironment(ctx context.Context, defaultWaitTimeout time.Duration, retryDelay time.Duration) (*Environment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Environment) Merge(e2 *Environment) { _ = "STUB: not implemented"; return }

type InteractiveShell struct {
	HistoryFile      string `json:"history_file"`
	HistoryLimit     *int   `json:"history_limit"`
	Prompt           string `json:"prompt"`
	ContinuousPrompt string `json:"continuous_prompt"`
	Completion       *bool  `json:"completion"`
	KillWholeLine    *bool  `json:"kill_whole_line"`
	ViMode           *bool  `json:"vi_mode"`
}

func (e *Environment) Load(ctx context.Context, defaultWaitTimeout time.Duration, retryDelay time.Duration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func GetSpecialFilePath(filename string) []string { _ = "STUB: not implemented"; return nil }

func GetHomeDirFilePath(filename string) string { _ = "STUB: not implemented"; return "" }

func GetCSVQConfigDirFilePath(filename string) string { _ = "STUB: not implemented"; return "" }

func GetConfigDirFilePath(filename string) string { _ = "STUB: not implemented"; return "" }

func GetCurrentDirFilePath(filename string) string { _ = "STUB: not implemented"; return "" }
