package cropper

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
)

// Cropper type.
type Cropper struct {
	Exec string
	Dir  string
}

// New returns a new cropper.
func New(exec, dir string) (*Cropper, error) {
	return &Cropper{exec, dir}, nil
}

// Screenshot generates a screenshot from site.
func (c *Cropper) Screenshot(ctx context.Context, uri string, width, height int) error {
	args, err := screenshotArgs(c.Exec, uri, width, height)
	if err != nil {
		return err
	}
	// Generate dir for preview.
	checksum := hashURL(uri, width, height)
	fdir := dirPath(c.Dir, checksum)
	if err := os.MkdirAll(fdir, os.ModePerm); err != nil {
		return err
	}
	// Invoke headless chrome and save 'screenshot.png' file.
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Dir = fdir
	return cmd.Run()
}

// Exists return true if screenshot exists on cache dir.
func (c *Cropper) Exists(uri string, width, height int) bool {
	check := hashURL(uri, width, height)
	d := dirPath(c.Dir, check)
	fname := previewName(d)
	_, err := os.Stat(fname)
	return !os.IsNotExist(err)
}

// Copy will read file in cache dir to w.
func (c *Cropper) Copy(w io.Writer, uri string, width, height int) error {
	checksum := hashURL(uri, width, height)
	fdir := dirPath(c.Dir, checksum)
	fname := previewName(fdir)
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	if _, err := io.Copy(w, f); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func hashURL(uri string, width, height int) string {
	s := fmt.Sprintf("%s,w=%d,h=%d", uri, width, height)
	return base64.URLEncoding.EncodeToString([]byte(s))
}

func dirPath(dir, s string) string {
	return filepath.Clean(dir + "/" + s)
}

func previewName(dir string) string {
	return dirPath(dir, "screenshot.png")
}

// screenshotArgs returns the arguments for exec.
func screenshotArgs(exec, uri string, width, height int) ([]string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if u.Hostname() == "" {
		return nil, errors.New("invalid host name")
	}
	if u.Scheme == "" {
		return nil, errors.New("invalid scheme")
	}

	if width < 1 {
		return nil, errors.New("width must be greater than zero")
	}
	if height < 1 {
		return nil, errors.New("height must be greater than zero")
	}
	wh := fmt.Sprintf("--window-size=%d,%d", width, height)
	args := []string{exec, "--headless", "--disable-gpu", "--hide-scrollbars", "--screenshot", wh, uri}
	return args, nil
}
