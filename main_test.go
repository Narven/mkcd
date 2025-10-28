package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpandTilde(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		checkFn func(t *testing.T, result string)
	}{
		{
			name:    "simple tilde",
			input:   "~",
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				home, err := os.UserHomeDir()
				require.NoError(t, err)
				assert.Equal(t, home, result)
			},
		},
		{
			name:    "tilde with path",
			input:   "~/test",
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				home, err := os.UserHomeDir()
				require.NoError(t, err)
				expected := filepath.Join(home, "test")
				assert.Equal(t, expected, result)
			},
		},
		{
			name:    "tilde with nested path",
			input:   "~/test/nested/path",
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				home, err := os.UserHomeDir()
				require.NoError(t, err)
				expected := filepath.Join(home, "test", "nested", "path")
				assert.Equal(t, expected, result)
			},
		},
		{
			name:    "no tilde",
			input:   "test/path",
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.Equal(t, "test/path", result)
			},
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.Equal(t, "", result)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := expandTilde(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				tt.checkFn(t, result)
			}
		})
	}
}

func TestResolvePath(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name    string
		input   string
		wantErr bool
		checkFn func(t *testing.T, result string)
	}{
		{
			name:    "relative path",
			input:   "testdir",
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.True(t, filepath.IsAbs(result))
				assert.Contains(t, result, "testdir")
			},
		},
		{
			name:    "absolute path",
			input:   tmpDir,
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.True(t, filepath.IsAbs(result))
				assert.Equal(t, tmpDir, result)
			},
		},
		{
			name:    "tilde expansion",
			input:   "~/test",
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.True(t, filepath.IsAbs(result))
				home, err := os.UserHomeDir()
				require.NoError(t, err)
				expected := filepath.Join(home, "test")
				assert.Equal(t, expected, result)
			},
		},
		{
			name:    "nested relative path",
			input:   "a/b/c",
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.True(t, filepath.IsAbs(result))
				assert.Contains(t, result, "a")
				assert.Contains(t, result, "b")
				assert.Contains(t, result, "c")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := resolvePath(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, filepath.IsAbs(result))
				tt.checkFn(t, result)
			}
		})
	}
}

func TestValidateOrCreateDir(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name          string
		setupFn       func(t *testing.T) string
		wantErr       bool
		errorContains string
		checkFn       func(t *testing.T, result string)
	}{
		{
			name: "create new directory",
			setupFn: func(t *testing.T) string {
				return filepath.Join(tmpDir, "newdir")
			},
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.True(t, filepath.IsAbs(result))
				info, err := os.Stat(result)
				require.NoError(t, err)
				assert.True(t, info.IsDir())
			},
		},
		{
			name: "create nested directory",
			setupFn: func(t *testing.T) string {
				return filepath.Join(tmpDir, "nested", "deep", "path")
			},
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.True(t, filepath.IsAbs(result))
				info, err := os.Stat(result)
				require.NoError(t, err)
				assert.True(t, info.IsDir())
			},
		},
		{
			name: "existing directory",
			setupFn: func(t *testing.T) string {
				existing := filepath.Join(tmpDir, "existing")
				err := os.MkdirAll(existing, 0755)
				require.NoError(t, err)
				return existing
			},
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.True(t, filepath.IsAbs(result))
				info, err := os.Stat(result)
				require.NoError(t, err)
				assert.True(t, info.IsDir())
			},
		},
		{
			name: "path exists but is a file",
			setupFn: func(t *testing.T) string {
				filePath := filepath.Join(tmpDir, "notadir")
				err := os.WriteFile(filePath, []byte("content"), 0644)
				require.NoError(t, err)
				return filePath
			},
			wantErr:       true,
			errorContains: "exists but is not a directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testPath := tt.setupFn(t)

			result, err := validateOrCreateDir(testPath)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.True(t, filepath.IsAbs(result))
				if tt.checkFn != nil {
					tt.checkFn(t, result)
				}
			}
		})
	}
}

func TestRunMkcd(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name          string
		input         string
		wantErr       bool
		errorContains string
		checkFn       func(t *testing.T, result string)
	}{
		{
			name:    "create simple directory",
			input:   filepath.Join(tmpDir, "simple"),
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.True(t, filepath.IsAbs(result))
				info, err := os.Stat(result)
				require.NoError(t, err)
				assert.True(t, info.IsDir())
			},
		},
		{
			name:    "create nested directory",
			input:   filepath.Join(tmpDir, "a", "b", "c"),
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.True(t, filepath.IsAbs(result))
				info, err := os.Stat(result)
				require.NoError(t, err)
				assert.True(t, info.IsDir())
			},
		},
		{
			name:    "existing directory",
			input:   tmpDir,
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.True(t, filepath.IsAbs(result))
				assert.Equal(t, tmpDir, result)
			},
		},
		{
			name:    "relative path",
			input:   "relativedir",
			wantErr: false,
			checkFn: func(t *testing.T, result string) {
				assert.True(t, filepath.IsAbs(result))
				assert.Contains(t, result, "relativedir")
			},
		},
		{
			name: "file exists instead of directory",
			input: func() string {
				filePath := filepath.Join(tmpDir, "file")
				os.WriteFile(filePath, []byte("content"), 0644)
				return filePath
			}(),
			wantErr:       true,
			errorContains: "exists but is not a directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := runMkcd(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.True(t, filepath.IsAbs(result))
				if tt.checkFn != nil {
					tt.checkFn(t, result)
				}
			}
		})
	}
}

func TestRunMkcdWithTilde(t *testing.T) {
	home, err := os.UserHomeDir()
	require.NoError(t, err)

	testPath := filepath.Join(home, "mkcd_test_dir")
	// Clean up in case it exists from a previous test
	os.RemoveAll(testPath)
	defer os.RemoveAll(testPath)

	result, err := runMkcd("~/mkcd_test_dir")
	require.NoError(t, err)
	assert.Equal(t, testPath, result)

	// Verify directory was created
	info, err := os.Stat(testPath)
	require.NoError(t, err)
	assert.True(t, info.IsDir())
}
