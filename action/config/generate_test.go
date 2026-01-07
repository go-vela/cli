// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"

	"github.com/spf13/afero"
	yaml "go.yaml.in/yaml/v3"

	"github.com/go-vela/cli/internal/output"
)

// Test constants for color configuration.
const (
	testThemeMonokai       = "monokai"
	testThemeMonokaiLight  = "monokailight"
	testThemeDracula       = "dracula"
	testThemeSolarizedDark = "solarized-dark"
	testFormatTerminal256  = "terminal256"
	testFormatTerminal16   = "terminal16"
	testFormatTerminal16m  = "terminal16m"
)

func TestConfig_Config_Generate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				GitHub: &GitHub{},
			},
		},
	}

	// run tests
	for _, test := range tests {
		// setup filesystem
		appFS = afero.NewMemMapFs()

		err := test.config.Generate()

		if test.failure {
			if err == nil {
				t.Errorf("Generate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Generate returned err: %v", err)
		}
	}
}

func TestConfig_Config_Generate_ColorTheme(t *testing.T) {
	// setup tests
	tests := []struct {
		name              string
		config            *Config
		expectTheme       bool
		expectThemeValue  string
		expectFormat      bool
		expectFormatValue string
	}{
		{
			name: "default theme not user specified - should not save theme",
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				GitHub: &GitHub{},
				Color: output.ColorOptions{
					Enabled:       true,
					Format:        testFormatTerminal256,
					Theme:         testThemeMonokai,
					ThemeLight:    testThemeMonokaiLight,
					UserSpecified: false,
				},
			},
			expectTheme:  false,
			expectFormat: false,
		},
		{
			name: "custom theme user specified - should save theme",
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				GitHub: &GitHub{},
				Color: output.ColorOptions{
					Enabled:       true,
					Format:        testFormatTerminal256,
					Theme:         testThemeDracula,
					ThemeLight:    testThemeMonokaiLight,
					UserSpecified: true,
				},
			},
			expectTheme:      true,
			expectThemeValue: testThemeDracula,
			expectFormat:     false,
		},
		{
			name: "default theme user specified - should save theme",
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				GitHub: &GitHub{},
				Color: output.ColorOptions{
					Enabled:       true,
					Format:        testFormatTerminal256,
					Theme:         testThemeMonokai,
					ThemeLight:    testThemeMonokaiLight,
					UserSpecified: true,
				},
			},
			expectTheme:      true,
			expectThemeValue: testThemeMonokai,
			expectFormat:     false,
		},
		{
			name: "custom format non-default - should save format",
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				GitHub: &GitHub{},
				Color: output.ColorOptions{
					Enabled:       true,
					Format:        testFormatTerminal16,
					Theme:         testThemeMonokai,
					ThemeLight:    testThemeMonokaiLight,
					UserSpecified: false,
				},
			},
			expectTheme:       false,
			expectFormat:      true,
			expectFormatValue: testFormatTerminal16,
		},
		{
			name: "custom format and theme both specified - should save both",
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				GitHub: &GitHub{},
				Color: output.ColorOptions{
					Enabled:       true,
					Format:        testFormatTerminal16m,
					Theme:         testThemeSolarizedDark,
					ThemeLight:    testThemeMonokaiLight,
					UserSpecified: true,
				},
			},
			expectTheme:       true,
			expectThemeValue:  testThemeSolarizedDark,
			expectFormat:      true,
			expectFormatValue: testFormatTerminal16m,
		},
		{
			name: "empty format - should not save format",
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				GitHub: &GitHub{},
				Color: output.ColorOptions{
					Enabled:       true,
					Format:        "",
					Theme:         testThemeMonokai,
					ThemeLight:    testThemeMonokaiLight,
					UserSpecified: false,
				},
			},
			expectTheme:  false,
			expectFormat: false,
		},
		{
			name: "default format terminal256 - should not save format",
			config: &Config{
				Action: "generate",
				File:   ".vela.yml",
				GitHub: &GitHub{},
				Color: output.ColorOptions{
					Enabled:       true,
					Format:        testFormatTerminal256,
					Theme:         testThemeMonokai,
					ThemeLight:    testThemeMonokaiLight,
					UserSpecified: false,
				},
			},
			expectTheme:  false,
			expectFormat: false,
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup filesystem
			appFS = afero.NewMemMapFs()

			err := test.config.Generate()
			if err != nil {
				t.Errorf("Generate returned err: %v", err)
				return
			}

			// read the generated file
			a := &afero.Afero{Fs: appFS}

			content, err := a.ReadFile(test.config.File)
			if err != nil {
				t.Errorf("Failed to read generated file: %v", err)
				return
			}

			// parse the YAML
			var configFile ConfigFile

			err = yaml.Unmarshal(content, &configFile)
			if err != nil {
				t.Errorf("Failed to unmarshal YAML: %v", err)
				return
			}

			// check theme
			if test.expectTheme {
				if configFile.ColorTheme == "" {
					t.Errorf("Expected color_theme to be saved, but it was not")
				}

				if configFile.ColorTheme != test.expectThemeValue {
					t.Errorf("ColorTheme = %v, want %v", configFile.ColorTheme, test.expectThemeValue)
				}
			} else {
				if configFile.ColorTheme != "" {
					t.Errorf("Expected color_theme to not be saved, but got: %v", configFile.ColorTheme)
				}
			}

			// check format
			if test.expectFormat {
				if configFile.ColorFormat == "" {
					t.Errorf("Expected color_format to be saved, but it was not")
				}

				if configFile.ColorFormat != test.expectFormatValue {
					t.Errorf("ColorFormat = %v, want %v", configFile.ColorFormat, test.expectFormatValue)
				}
			} else {
				if configFile.ColorFormat != "" {
					t.Errorf("Expected color_format to not be saved, but got: %v", configFile.ColorFormat)
				}
			}
		})
	}
}

func TestConfig_genBytes_ColorTheme(t *testing.T) {
	// setup tests
	tests := []struct {
		name         string
		config       *Config
		wantTheme    string
		wantFormat   string
		wantNoTheme  bool
		wantNoFormat bool
	}{
		{
			name: "user specified theme",
			config: &Config{
				GitHub: &GitHub{},
				Color: output.ColorOptions{
					Enabled:       true,
					Theme:         testThemeDracula,
					Format:        testFormatTerminal256,
					UserSpecified: true,
				},
			},
			wantTheme:    testThemeDracula,
			wantNoFormat: true, // default format shouldn't be saved
		},
		{
			name: "default theme not user specified",
			config: &Config{
				GitHub: &GitHub{},
				Color: output.ColorOptions{
					Enabled:       true,
					Theme:         testThemeMonokai,
					Format:        testFormatTerminal256,
					UserSpecified: false,
				},
			},
			wantNoTheme:  true,
			wantNoFormat: true,
		},
		{
			name: "custom format",
			config: &Config{
				GitHub: &GitHub{},
				Color: output.ColorOptions{
					Enabled:       true,
					Theme:         testThemeMonokai,
					Format:        testFormatTerminal16,
					UserSpecified: false,
				},
			},
			wantFormat:  testFormatTerminal16,
			wantNoTheme: true,
		},
		{
			name: "both theme and format custom",
			config: &Config{
				GitHub: &GitHub{},
				Color: output.ColorOptions{
					Enabled:       true,
					Theme:         testThemeSolarizedDark,
					Format:        testFormatTerminal16m,
					UserSpecified: true,
				},
			},
			wantTheme:  testThemeSolarizedDark,
			wantFormat: testFormatTerminal16m,
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bytes, err := genBytes(test.config)
			if err != nil {
				t.Errorf("genBytes returned err: %v", err)
				return
			}

			var configFile ConfigFile

			err = yaml.Unmarshal(bytes, &configFile)
			if err != nil {
				t.Errorf("Failed to unmarshal YAML: %v", err)
				return
			}

			if test.wantNoTheme {
				if configFile.ColorTheme != "" {
					t.Errorf("Expected no theme, but got: %v", configFile.ColorTheme)
				}
			} else if configFile.ColorTheme != test.wantTheme {
				t.Errorf("ColorTheme = %v, want %v", configFile.ColorTheme, test.wantTheme)
			}

			if test.wantNoFormat {
				if configFile.ColorFormat != "" {
					t.Errorf("Expected no format, but got: %v", configFile.ColorFormat)
				}
			} else if configFile.ColorFormat != test.wantFormat {
				t.Errorf("ColorFormat = %v, want %v", configFile.ColorFormat, test.wantFormat)
			}
		})
	}
}
