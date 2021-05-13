// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package taskresourceattribute

import (
	"fmt"

	"github.com/spf13/pflag"
)

// GetPFlagSet will return strongly types pflags for all fields in AttrUpdateConfig and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg AttrUpdateConfig) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("AttrUpdateConfig", pflag.ExitOnError)
	cmdFlags.StringVar(&(DefaultUpdateConfig.AttrFile), fmt.Sprintf("%v%v", prefix, "attrFile"), DefaultUpdateConfig.AttrFile, "attribute file name to be used for updating attribute for the resource type.")
	return cmdFlags
}
