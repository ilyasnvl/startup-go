package campaign

type CampaignFormatter struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	Name          string `json:"name"`
	ShortDesc     string `json:"short_description"`
	ImageURL      string `json:"image_url"`
	GoalAmount    int    `json:"goal_amount"`
	CurrentAmount int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDesc = campaign.ShortDesc
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func Formatcampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormater := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormater)
	}

	return campaignsFormatter
}
