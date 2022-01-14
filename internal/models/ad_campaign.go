package models

import (
	"database/sql"
	gapi "gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
	"time"
)

func FromSliceToString(slice []string) string {
	var resultString string
	for _, str := range slice {
		resultString += str + ","
	}
	size := len(resultString)
	if size > 0 && resultString[size-1] == ',' {
		resultString = resultString[:size-1]
	}
	return resultString
}

type AdCampaign struct {
	Id                          int64
	Name                        string
	TrackingUrlTemplate         string
	BaseCampaign                string
	CampaignBudget              string
	StartDate                   sql.NullTime
	EndDate                     sql.NullTime
	FinalUrlSuffix              string
	OptimizationScore           float64
	ResourceName                string
	Status                      string
	ServingStatus               string
	AdServingOptimizationStatus string
	AdvertisingChannelType      string
	AdvertisingChannelSubType   string
	AccessibleBiddingStrategy   string
	PaymentMode                 string
	Labels                      string
	RealTimeBiddingSetting      bool
	UrlCustomParameters         string
}

func (c *AdCampaign) FromApi(gads *gapi.Campaign) {
	c.Id = gads.Id
	c.Name = gads.Name
	c.TrackingUrlTemplate = gads.TrackingUrlTemplate
	c.BaseCampaign = gads.BaseCampaign
	c.CampaignBudget = gads.CampaignBudget

	startDate, err := time.Parse(time.RFC3339, gads.StartDate)
	if err == nil {
		c.StartDate = sql.NullTime{Time: startDate}
	}

	endDate, err := time.Parse(time.RFC3339, gads.EndDate)
	if err == nil {
		c.EndDate = sql.NullTime{Time: endDate}
	}

	c.FinalUrlSuffix = gads.FinalUrlSuffix
	c.OptimizationScore = float64(gads.OptimizationScore)
	c.ResourceName = gads.ResourceName
	c.Status = gads.Status.String()
	c.ServingStatus = gads.ServingStatus.String()
	c.AdServingOptimizationStatus = gads.AdServingOptimizationStatus.String()
	c.AdvertisingChannelType = gads.AdvertisingChannelType.String()
	c.AdvertisingChannelSubType = gads.AdvertisingChannelSubType.String()
	c.AccessibleBiddingStrategy = gads.AccessibleBiddingStrategy
	c.PaymentMode = gads.PaymentMode.String()
	c.RealTimeBiddingSetting = gads.RealTimeBiddingSetting
	c.Labels = FromSliceToString(gads.Labels)
	c.RealTimeBiddingSetting = gads.RealTimeBiddingSetting
	c.UrlCustomParameters = FromSliceToString(gads.UrlCustomParameters)
}

func (c *AdCampaign) TableName() string {
	return "ad_g_campaigns"
}
