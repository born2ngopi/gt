package workspace

import (
	"mime"
	"os"
	"path"
	"path/filepath"
)

type workspace struct {
	RootPath string
	Files    []string
}

type Workspace interface {
	ListDir() *workspace
	ReadFile(file_path string) (*os.File, error)
	GetFileType(f *os.File) string
}

func Init(root_path string) Workspace {
	return &workspace{
		RootPath: root_path,
		Files:    []string{},
	}
}

func (w *workspace) ListDir() *workspace {

	if err := filepath.Walk(w.RootPath, func(file_path string, info os.FileInfo, err error) error {
		if file_path != w.RootPath && file_path != path.Join(w.RootPath, ".git") {
			w.Files = append(w.Files, file_path)
		}
		return nil
	}); err != nil {
		return nil
	}

	return w
}

func (w *workspace) ReadFile(file_path string) (*os.File, error) {
	return os.Open(file_path)
}

func (w *workspace) GetFileType(f *os.File) string {
	ext := filepath.Ext(f.Name())

	return mime.TypeByExtension(ext)
}
