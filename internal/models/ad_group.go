package models

import (
	gapi "gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
	"strconv"
	"strings"
)

type AdGroup struct {
	Id                       int64
	Name                     string
	BaseAdGroup              string
	TrackingUrlTemplate      string
	CampaignId               int64
	CpcBidMicros             int64
	CpmBidMicros             int64
	TargetCpaMicros          int64
	CpvBidMicros             int64
	TargetCpmMicros          int64
	TargetRoas               float64
	PercentCpcBidMicros      int64
	FinalUrlSuffix           string
	EffectiveTargetCpaMicros int64
	EffectiveTargetRoas      float64
	ResourceName             string
	Labels                   string
}

func (c *AdGroup) getCompaignId(str string) int64 {
	s := strings.Split(str, "/")
	i, err := strconv.ParseInt(s[len(s)-1], 10, 64)
	if err != nil {
		return -1
	}
	return i
}

func (c *AdGroup) FromApi(gads *gapi.AdGroup) {
	c.Id = gads.Id
	c.Name = gads.Name
	c.BaseAdGroup = gads.BaseAdGroup
	c.TrackingUrlTemplate = gads.TrackingUrlTemplate
	c.CampaignId = c.getCompaignId(gads.Campaign)
	c.CpcBidMicros = gads.CpcBidMicros
	c.CpmBidMicros = gads.CpmBidMicros
	c.TargetCpaMicros = gads.TargetCpaMicros
	c.CpvBidMicros = gads.CpvBidMicros
	c.TargetCpmMicros = gads.TargetCpmMicros
	c.TargetRoas = float64(gads.TargetRoas)
	c.PercentCpcBidMicros = gads.PercentCpcBidMicros
	c.FinalUrlSuffix = gads.FinalUrlSuffix
	c.EffectiveTargetCpaMicros = gads.EffectiveTargetCpaMicros
	c.EffectiveTargetRoas = float64(gads.EffectiveTargetRoas)
	c.ResourceName = gads.ResourceName
	c.Labels = FromSliceToString(gads.Labels)
}

func (AdGroup) TableName() string {
	return "ad_g_groups"
}
