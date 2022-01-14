package models

import (
	"gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
)

type Ads struct {
	Id                  int64
	Name                string
	ResourceName        string
	UrlCollections      string
	DisplayUrl          string
	TrackingUrlTemplate string
	AddedByGoogleAds    bool
	FinalUrls           string
	FinalAppUrls        string
	FinalMobileUrls     string
	FinalUrlSuffix      string
}

func (Ads) TableName() string {
	return "ad_g_ads"
}

func (a *Ads) FromApi(ad *google_api.AdGroupAd) {
	a.Id = ad.Id
	a.Name = ad.Name
	a.ResourceName = ad.ResourceName
	a.UrlCollections = FromSliceToString(ad.UrlCollections)
	a.DisplayUrl = ad.DisplayUrl
	a.TrackingUrlTemplate = ad.TrackingUrlTemplate
	a.AddedByGoogleAds = ad.AddedByGoogleAds
	a.FinalUrls = FromSliceToString(ad.FinalUrls)
	a.FinalAppUrls = FromSliceToString(ad.FinalAppUrls)
	a.FinalMobileUrls = a.FinalMobileUrls
	a.FinalUrlSuffix = ad.FinalUrlSuffix
}
