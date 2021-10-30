package spec_file

import (
	"io/fs"
)

type Root struct {
	Spec Spec `hcl:"spec,block"`
}

type Spec struct {
	Name         string        `hcl:"name,label"`
	Config       Config        `hcl:"config,block"`
	System       System        `hcl:"system,block"`
	Applications []Application `hcl:"application,block"`
}

type Config struct {
	OsDriver         DriverConfig `hcl:"os,block"`
	PackagesDriver   DriverConfig `hcl:"packages,block"`
	FilesystemDriver DriverConfig `hcl:"filesystem,block"`
	ServicesDriver   DriverConfig `hcl:"services,block"`
	FirewallDriver   DriverConfig `hcl:"firewall,block"`
}

type DriverConfig struct {
	Name string `hcl:"name,label"`
}

type System struct {
	OsOptions     []OsOption     `hcl:"os_option,block"`
	Users         []User         `hcl:"user,block"`
	Packages      []Package      `hcl:"package,block"`
	Directories   []Directory    `hcl:"directory,block"`
	Files         []File         `hcl:"file,block"`
	Services      []Service      `hcl:"service,block"`
	FirewallRules []FirewallRule `hcl:"firewall,block"`
}

type Application struct {
	Name          string         `hcl:"name,label"`
	OsOptions     []OsOption     `hcl:"os_option,block"`
	Users         []User         `hcl:"user,block"`
	Packages      []Package      `hcl:"package,block"`
	Directories   []Directory    `hcl:"directory,block"`
	Files         []File         `hcl:"file,block"`
	Services      []Service      `hcl:"service,block"`
	FirewallRules []FirewallRule `hcl:"firewall,block"`
}

type OsOption struct {
	Type  string `hcl:"type,label"`
	Name  string `hcl:"name,label"`
	Value string `hcl:"value"`
}

type User struct {
	Name   string `hcl:"name,label"`
	System bool   `hcl:"system,optional"`
}

type Package struct {
	Name    string `hcl:"name,label"`
	Removed bool   `hcl:"removed,optional"`
}

type Directory struct {
	Path        string      `hcl:"path,label"`
	User        string      `hcl:"user"`
	Group       string      `hcl:"group"`
	Permissions fs.FileMode `hcl:"permissions"`
}

type File struct {
	Path         string            `hcl:"path,label"`
	From         string            `hcl:"from"`
	IsTemplate   bool              `hcl:"template,optional"`
	TemplateVars map[string]string `hcl:"template_vars,optional"`
	IsCompressed bool              `hcl:"compressed,optional"`
	User         string            `hcl:"user"`
	Group        string            `hcl:"group"`
	Permissions  fs.FileMode       `hcl:"permissions"`
	Hash         string            `hcl:"hash,optional"`
}

type Service struct {
	Name     string `hcl:"name,label"`
	Disabled bool   `hcl:"disabled,optional"`
}

type FirewallRule struct {
	Name     string   `hcl:"name,label"`
	Id       string   `hcl:"identifier,optional"`
	Ports    []string `hcl:"ports"`
	Protocol string   `hcl:"protocol,optional"`
	Targets  []string `hcl:"targets"`
	IsOutput bool     `hcl:"is_output,optional"`
}
