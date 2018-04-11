package main

// Paramaters default cloudformation paramaters
// Parameters enable you to input custom values to your template each time you create or update a stack.
// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html
type Paramaters struct {
	AllowedValues         []string `json:"AllowedValues,omitempty"`
	MaxValue              int      `json:"MaxValue,omitempty"`
	MinValue              int      `json:"MinValue,omitempty"`
	Description           string   `json:"Description,omitempty"`
	Type                  string
	ConstraintDescription string `json:"ConstraintDescription,omitempty"`
	MinLength             string `json:"MinLength,omitempty"`
	MaxLength             string `json:"MaxLength,omitempty"`
	Default               string `json:"Default,omitempty"`
	AllowedPattern        string `json:"AllowedPattern,omitempty"`
	NoEcho                bool   `json:"NoEcho,omitempty"`
	String                string `json:"String,omitempty"`
	Number                int    `json:"Number,omitempty"`
}
