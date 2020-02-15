package cp1

import (
	"io"
	"os"
	"path/filepath"
)

// CopyDir copies an entire source directory to the dest directory.
//
// This is akin to `cp -a src/* dest/`
//
// We copy the directory here rather than jumping out to a shell so we can
// support multiple operating systems.
func CopyDir(source string, dest string) error {

	// get properties of source dir
	si, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, si.Mode())
	if err != nil {
		return err
	}

	d, err := os.Open(source)
	if err != nil {
		return err
	}
	defer d.Close()

	objects, err := d.Readdir(-1)

	for _, obj := range objects {

		sp := filepath.Join(source, "/", obj.Name())

		dp := filepath.Join(dest, "/", obj.Name())

		if obj.IsDir() {
			err = CopyDir(sp, dp)
			if err != nil {
				return err
			}
		} else {
			// perform copy
			err = CopyFile(sp, dp)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// CopyFile copies a source file to a destination.
//
// It follows symbolic links and retains modes.
func CopyFile(source string, dest string) error {
	ln, err := os.Readlink(source)
	if err == nil {
		return os.Symlink(ln, dest)
	}
	s, err := os.Open(source)
	if err != nil {
		return err
	}

	defer s.Close()

	d, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer d.Close()

	_, err = io.Copy(d, s)
	if err != nil {
		return err
	}

	si, err := os.Stat(source)
	if err != nil {
		return err
	}
	err = os.Chmod(dest, si.Mode())

	return err
}
