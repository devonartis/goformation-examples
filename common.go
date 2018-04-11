package main

type instanceParams struct {

	// Struct to declare instance paramaters

	Description   string
	Type          string
	Default       string
	AllowedValues []string
}

type keyParam struct {
	Description           string
	Type                  string
	ConstraintDescription string
}

type sshLocParam struct {
	Description           string
	Type                  string
	ConstraintDescription string
	MinLength             string
	MaxLength             string
	Default               string
	AllowedPattern        string
}
