package models

// Campaign
type CampaignInsight struct {
	CriterionId  int64
	ResourceName string
	Campaign     string
	DisplayName  string
	BidModifier  float32
	Negative     bool
}

func (CampaignInsight) TableName() string {
	return "g_campaign_insights"
}

// AdSet
type AdGroupInsight struct {
	Keyword      string
	AdGroupName  string
	CampaignName string
	Impressions  int64
	Clicks       int64
	Ctr          float32
	AverageCpc   float32
}

func (AdGroupInsight) TableName() string {
	return "g_ad_group_insights"
}
