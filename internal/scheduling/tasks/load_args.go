package tasks

import (
	"github.com/mitchellh/mapstructure"
	"time"
)

type LoadArgs struct {
	AdAccountId string     `json:",omitempty" mapstructure:"AdAccountId,omitempty"`
	BrandId     int32      `json:",omitempty" mapstructure:"BrandId,omitempty"`
	Before      *time.Time `json:",omitempty" mapstructure:"Before,omitempty"`
	Token       string     `json:",omitempty" mapstructure:"Token,omitempty"`
}

func (a *LoadArgs) FromArgs(m map[string]interface{}) {
	_ = mapstructure.Decode(m, a)
}
