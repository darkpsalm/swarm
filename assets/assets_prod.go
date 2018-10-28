// +build embed

// Package assets is generated by github.com/omeid/go-resources
package assets

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileSystem is an http.FileSystem implementation.
type FileSystem struct {
	files map[string]File
}

// String returns the content of the file as string.
func (fs *FileSystem) String(name string) (string, bool) {
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 ||
		strings.Contains(name, "\x00") {
		return "", false
	}

	file, ok := fs.files[name]

	if !ok {
		return "", false
	}

	return string(file.data), true
}

// Open implements http.FileSystem.Open
func (fs *FileSystem) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 ||
		strings.Contains(name, "\x00") {
		return nil, errors.New("http: invalid character in file path")
	}
	file, ok := fs.files[name]
	if !ok {
		files := []os.FileInfo{}
		for path, file := range fs.files {
			if strings.HasPrefix(path, name) {
				fi := file.fi
				files = append(files, &fi)
			}
		}

		if len(files) == 0 {
			return nil, os.ErrNotExist
		}

		//We have a directory.
		return &File{
			fi: FileInfo{
				isDir: true,
				files: files,
			}}, nil
	}
	file.Reader = bytes.NewReader(file.data)
	return &file, nil
}

// File implements http.File
type File struct {
	*bytes.Reader
	data []byte
	fi   FileInfo
}

// Close is a noop-closer.
func (f *File) Close() error {
	return nil
}

// Readdir implements http.File.Readdir
func (f *File) Readdir(count int) ([]os.FileInfo, error) {
	return nil, os.ErrNotExist
}

// Stat implements http.Stat.Readdir
func (f *File) Stat() (os.FileInfo, error) {
	return &f.fi, nil
}

// FileInfo implements the os.FileInfo interface.
type FileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
	sys     interface{}

	files []os.FileInfo
}

// Name implements os.FileInfo.Name
func (f *FileInfo) Name() string {
	return f.name
}

// Size implements os.FileInfo.Size
func (f *FileInfo) Size() int64 {
	return f.size
}

// Mode implements os.FileInfo.Mode
func (f *FileInfo) Mode() os.FileMode {
	return f.mode
}

// ModTime implements os.FileInfo.ModTime
func (f *FileInfo) ModTime() time.Time {
	return f.modTime
}

// IsDir implements os.FileInfo.IsDir
func (f *FileInfo) IsDir() bool {
	return f.isDir
}

// Readdir implements os.FileInfo.Readdir
func (f *FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return f.files, nil
}

// Sys returns the underlying value.
func (f *FileInfo) Sys() interface{} {
	return f.sys
}

var FS *FileSystem

func init() {
	FS = &FileSystem{
		files: map[string]File{
			"/assets/static/HotReload.js": File{
				data: []byte{
					0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x7b, 0x20, 0x53, 0x6f, 0x63,
					0x6b, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x7d, 0x20,
					0x66, 0x72, 0x6f, 0x6d, 0x20, 0x22, 0x2e, 0x2f, 0x53, 0x6f, 0x63, 0x6b,
					0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6a, 0x73, 0x22,
					0x3b, 0x0d, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x73, 0x63, 0x20,
					0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74,
					0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x28, 0x29, 0x3b, 0x0d, 0x0a, 0x73,
					0x63, 0x2e, 0x6f, 0x6e, 0x28, 0x65, 0x20, 0x3d, 0x3e, 0x20, 0x65, 0x2e,
					0x74, 0x79, 0x70, 0x65, 0x20, 0x3d, 0x3d, 0x20, 0x22, 0x72, 0x65, 0x6c,
					0x6f, 0x61, 0x64, 0x22, 0x20, 0x26, 0x26, 0x20, 0x77, 0x69, 0x6e, 0x64,
					0x6f, 0x77, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
					0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x28, 0x29, 0x29, 0x3b, 0x0d, 0x0a,
					0x73, 0x63, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x28, 0x29,
					0x3b, 0x0d, 0x0a, 
				},
				fi: FileInfo{
					name:    "HotReload.js",
					size:    159,
					modTime: time.Unix(0, 1540724656882958000),
					isDir:   false,
				},
			},"/assets/static/SocketClient.js": File{
				data: []byte{
					0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73,
					0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x69, 0x74, 0x74, 0x65,
					0x72, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e,
					0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x28, 0x29, 0x20, 0x7b,
					0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x68,
					0x69, 0x73, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73,
					0x20, 0x3d, 0x20, 0x5b, 0x5d, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6c, 0x69, 0x73,
					0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x63, 0x65, 0x72, 0x20,
					0x3d, 0x20, 0x5b, 0x5d, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6f, 0x6e, 0x20, 0x3d,
					0x20, 0x28, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x29, 0x20,
					0x3d, 0x3e, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6c,
					0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x75, 0x73,
					0x68, 0x28, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x29, 0x3b,
					0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x3b,
					0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x68,
					0x69, 0x73, 0x2e, 0x65, 0x6d, 0x69, 0x74, 0x20, 0x3d, 0x20, 0x28, 0x65,
					0x76, 0x65, 0x6e, 0x74, 0x29, 0x20, 0x3d, 0x3e, 0x20, 0x7b, 0x0d, 0x0a,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x2f, 0x2a, 0x2a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61,
					0x6e, 0x79, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x20, 0x6c,
					0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x20, 0x2a, 0x2f, 0x0d,
					0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
					0x65, 0x72, 0x73, 0x2e, 0x66, 0x6f, 0x72, 0x45, 0x61, 0x63, 0x68, 0x28,
					0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x20, 0x3d, 0x3e, 0x20,
					0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x28, 0x65, 0x76, 0x65,
					0x6e, 0x74, 0x29, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2f, 0x2a, 0x2a, 0x20, 0x43,
					0x6c, 0x65, 0x61, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x60, 0x6f, 0x6e,
					0x63, 0x65, 0x60, 0x20, 0x71, 0x75, 0x65, 0x75, 0x65, 0x20, 0x2a, 0x2f,
					0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65,
					0x6e, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x66, 0x6f,
					0x72, 0x45, 0x61, 0x63, 0x68, 0x28, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
					0x65, 0x72, 0x20, 0x3d, 0x3e, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
					0x65, 0x72, 0x28, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x29, 0x29, 0x3b, 0x0d,
					0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
					0x65, 0x72, 0x73, 0x4f, 0x6e, 0x63, 0x65, 0x72, 0x20, 0x3d, 0x20, 0x5b,
					0x5d, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x7d, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0d, 0x0a, 0x7d,
					0x0d, 0x0a, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x63, 0x6c, 0x61,
					0x73, 0x73, 0x20, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6c, 0x69,
					0x65, 0x6e, 0x74, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63,
					0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x28, 0x29,
					0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x3d,
					0x20, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
					0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x3b, 0x0d, 0x0a,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x73,
					0x74, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x20, 0x3d,
					0x20, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
					0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x20, 0x3d, 0x3d, 0x3d, 0x20, 0x22,
					0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x22, 0x20, 0x3f, 0x20, 0x22, 0x77,
					0x73, 0x73, 0x3a, 0x2f, 0x2f, 0x22, 0x20, 0x3a, 0x20, 0x22, 0x77, 0x73,
					0x3a, 0x2f, 0x2f, 0x22, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x20, 0x64, 0x6f, 0x6d,
					0x61, 0x69, 0x6e, 0x20, 0x3d, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
					0x6f, 0x6e, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x20,
					0x7c, 0x7c, 0x20, 0x22, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73,
					0x74, 0x22, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x75, 0x72, 0x6c, 0x20, 0x3d, 0x20,
					0x60, 0x24, 0x7b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x7d,
					0x24, 0x7b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x7d, 0x3a, 0x24, 0x7b,
					0x70, 0x6f, 0x72, 0x74, 0x7d, 0x2f, 0x5f, 0x5f, 0x73, 0x77, 0x61, 0x72,
					0x6d, 0x5f, 0x5f, 0x2f, 0x77, 0x73, 0x60, 0x3b, 0x0d, 0x0a, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x65,
					0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x20, 0x3d, 0x20, 0x6e, 0x65, 0x77,
					0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x69, 0x74, 0x74, 0x65,
					0x72, 0x28, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0d,
					0x0a, 0x20, 0x20, 0x20, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
					0x63, 0x74, 0x28, 0x29, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f,
					0x75, 0x74, 0x28, 0x28, 0x29, 0x20, 0x3d, 0x3e, 0x20, 0x74, 0x68, 0x69,
					0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x28, 0x29, 0x2c,
					0x20, 0x35, 0x30, 0x30, 0x30, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20,
					0x20, 0x7d, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6f, 0x6e, 0x28, 0x66,
					0x6e, 0x29, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x65, 0x6d, 0x69, 0x74, 0x74,
					0x65, 0x72, 0x2e, 0x6f, 0x6e, 0x28, 0x66, 0x6e, 0x29, 0x3b, 0x0d, 0x0a,
					0x20, 0x20, 0x20, 0x20, 0x7d, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63,
					0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x28, 0x29, 0x20, 0x7b, 0x0d, 0x0a,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x73,
					0x6f, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x28, 0x22, 0x25, 0x63, 0x43,
					0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x6f,
					0x20, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x61,
					0x74, 0x20, 0x22, 0x20, 0x2b, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x75,
					0x72, 0x6c, 0x2c, 0x20, 0x22, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20,
					0x23, 0x32, 0x33, 0x37, 0x61, 0x62, 0x65, 0x22, 0x29, 0x3b, 0x0d, 0x0a,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x65, 0x74, 0x54,
					0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x28, 0x28, 0x29, 0x20, 0x3d, 0x3e,
					0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x63, 0x6c, 0x69,
					0x65, 0x6e, 0x74, 0x20, 0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x57, 0x65,
					0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x28, 0x74, 0x68, 0x69, 0x73,
					0x2e, 0x75, 0x72, 0x6c, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x68, 0x69, 0x73,
					0x2e, 0x62, 0x69, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x28,
					0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x7d, 0x2c, 0x20, 0x30, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x7d, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x73, 0x65, 0x6e, 0x64, 0x28,
					0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x2c, 0x20, 0x64,
					0x61, 0x74, 0x61, 0x29, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x69, 0x66, 0x20, 0x28, 0x74, 0x68, 0x69, 0x73,
					0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x61, 0x64,
					0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x20, 0x3d, 0x3d, 0x3d, 0x20, 0x31,
					0x29, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x63, 0x6c,
					0x69, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x28, 0x4a, 0x53,
					0x4f, 0x4e, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x69, 0x66, 0x79,
					0x28, 0x7b, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x20, 0x65, 0x76,
					0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x2c, 0x20, 0x64, 0x61, 0x74,
					0x61, 0x3a, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x7c, 0x7c, 0x20, 0x7b,
					0x7d, 0x20, 0x7d, 0x29, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x7d, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d,
					0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2f, 0x2a, 0x2a, 0x20, 0x57, 0x69,
					0x72, 0x65, 0x73, 0x20, 0x75, 0x70, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73,
					0x6f, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
					0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x20, 0x74, 0x6f,
					0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x20,
					0x6f, 0x6e, 0x20, 0x6f, 0x75, 0x72, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74,
					0x20, 0x65, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x20, 0x2a, 0x2f, 0x0d,
					0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x45, 0x76, 0x65,
					0x6e, 0x74, 0x73, 0x28, 0x29, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x63, 0x6c,
					0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x6e, 0x6f, 0x70, 0x65, 0x6e, 0x20,
					0x3d, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x3d, 0x3e, 0x20, 0x7b,
					0x20, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x67,
					0x28, 0x22, 0x25, 0x63, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
					0x64, 0x22, 0x2c, 0x20, 0x22, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20,
					0x23, 0x32, 0x33, 0x37, 0x61, 0x62, 0x65, 0x22, 0x29, 0x3b, 0x20, 0x74,
					0x68, 0x69, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6f,
					0x6e, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x20, 0x3d, 0x20, 0x28, 0x65, 0x76,
					0x65, 0x6e, 0x74, 0x29, 0x20, 0x3d, 0x3e, 0x20, 0x74, 0x68, 0x69, 0x73,
					0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x28, 0x29,
					0x3b, 0x20, 0x7d, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
					0x20, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
					0x74, 0x2e, 0x6f, 0x6e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x3d, 0x20,
					0x28, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x29, 0x20, 0x3d, 0x3e, 0x20, 0x63,
					0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72,
					0x28, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20,
					0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x63,
					0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x6e, 0x6d, 0x65, 0x73, 0x73,
					0x61, 0x67, 0x65, 0x20, 0x3d, 0x20, 0x28, 0x65, 0x76, 0x65, 0x6e, 0x74,
					0x29, 0x20, 0x3d, 0x3e, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x64,
					0x61, 0x74, 0x61, 0x20, 0x26, 0x26, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e,
					0x65, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6d, 0x69, 0x74,
					0x28, 0x4a, 0x53, 0x4f, 0x4e, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x28,
					0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x29, 0x29,
					0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0d, 0x0a, 0x7d, 0x0d,
					0x0a, 
				},
				fi: FileInfo{
					name:    "SocketClient.js",
					size:    1837,
					modTime: time.Unix(0, 1540724656879968800),
					isDir:   false,
				},
			},"/assets/static/test-asset.js": File{
				data: []byte{
					0x61, 0x6c, 0x65, 0x72, 0x74, 0x28, 0x27, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
					0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x27, 0x29, 0x3b, 
				},
				fi: FileInfo{
					name:    "test-asset.js",
					size:    21,
					modTime: time.Unix(0, 1540714110706880800),
					isDir:   false,
				},
			},
		},
	}
}