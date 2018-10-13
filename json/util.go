package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

const conf_path = "conf.toml"

type Config struct {
	Report_dir  string
	Path_dir    string
	Port        string
	Report_tmpl string
	Path_tmpl   string
}

func Load_config(c *Config) error {
	p, e := os.Executable()

	if e != nil {
		return e
	}

	sep := string(os.PathSeparator)
	exe_dir := strings.Split(filepath.Dir(p), sep)

	_, e = toml.DecodeFile(Complete_path(exe_dir, conf_path, sep), c)

	if c.Report_dir == "" {
		c.Report_dir = "reports"
	}
	c.Report_dir = Complete_path(exe_dir, c.Report_dir, sep)
	if c.Path_dir == "" {
		c.Path_dir = "paths"
	}
	c.Path_dir = Complete_path(exe_dir, c.Path_dir, sep)
	if c.Port == "" {
		c.Port = ":80"
	}
	if c.Report_tmpl == "" {
		c.Report_tmpl = strings.Join([]string{"tmpls", "report.tmpl"}, sep)
	}
	c.Report_tmpl = Complete_path(exe_dir, c.Report_tmpl, sep)
	if c.Path_tmpl == "" {
		c.Path_tmpl = strings.Join([]string{"tmpls", "path.tmpl"}, sep)
	}
	c.Path_tmpl = Complete_path(exe_dir, c.Path_tmpl, sep)

	return nil
}

func Complete_path(exe_dir []string, path, sep string) string {
	pp := append_list(exe_dir, strings.Split(path, sep))
	return strings.Join(pp, sep)
}

func append_list(a, b []string) []string {
	p := a
	for _, x := range b {
		p = append(p, x)
	}
	return p
}
