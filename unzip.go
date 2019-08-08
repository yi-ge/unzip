package unzip

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Unzip - struct
type Unzip struct {
	Src  string
	Dest string
}

// New - Create a new Unzip.
func New(src string, dest string) Unzip {
	return Unzip{src, dest}
}

func writeSymbolicLink(filePath string, targetPath string) error {
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}

	err = os.Symlink(targetPath, filePath)
	if err != nil {
		return err
	}

	return nil
}

// Extract - Extract zip file.
func (uz Unzip) Extract() error {
	r, err := zip.OpenReader(uz.Src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(uz.Dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(uz.Dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			mode := f.FileHeader.Mode()
			if mode&os.ModeType == os.ModeSymlink {
				data, err := ioutil.ReadAll(rc)
				if err != nil {
					return err
				}
				writeSymbolicLink(path, string(data))
			} else {
				os.MkdirAll(filepath.Dir(path), f.Mode())
				outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err != nil {
					return err
				}
				defer func() {
					if err := outFile.Close(); err != nil {
						panic(err)
					}
				}()

				_, err = io.Copy(outFile, rc)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
