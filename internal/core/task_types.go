package core

type TaskType string

const (
	TaskLoadCampaignsInsights = TaskType("load_campaigns_insights")
	TaskLoadAdGroupsInsights  = TaskType("load_ad_groups_insights")
	TaskLoadAdsInsights       = TaskType("load_ads_insights")
	TaskSyncInsights          = TaskType("sync_insights")
	TaskLoadAdCampaigns       = TaskType("load_ad_campaigns")
	TaskLoadAdGroups          = TaskType("load_ad_groups")
	TaskLoadAds               = TaskType("load_ads")
	TaskLoadPages             = TaskType("load_pages")
	TaskSyncCampaigns         = TaskType("sync_campaigns")
)

var CodeByTaskType = map[TaskType]byte{
	TaskLoadCampaignsInsights: 1,
	TaskLoadAdGroupsInsights:  2,
	TaskLoadAdsInsights:       3,
	TaskSyncInsights:          4,
	TaskLoadAdCampaigns:       5,
	TaskLoadAds:               6,
	TaskLoadPages:             7,
	TaskSyncCampaigns:         8,
}

var TaskTypeByCode = map[byte]TaskType{
	1: TaskLoadCampaignsInsights,
	2: TaskLoadAdGroupsInsights,
	3: TaskLoadAdsInsights,
	4: TaskSyncInsights,
	5: TaskLoadAdCampaigns,
	6: TaskLoadAds,
	7: TaskLoadPages,
	8: TaskSyncCampaigns,
}
