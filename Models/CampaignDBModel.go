package Models

import "one_test_case/DBConfig"

func GetAllCampaigns(campaigns *[]Campaign) (err error) {
	if err = DBConfig.DB.Find(&campaigns, "status = 1").Error; err != nil {
		return err
	}
	return nil
}

func CreateCampaign(campaign *Campaign) (err error) {
	if err = DBConfig.DB.Create(campaign).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCampaign(campaign *Campaign) (err error) {
	if err = DBConfig.DB.Save(campaign).Error; err != nil {
		return err
	}
	return nil
}

func GetCampaignById(campaign *Campaign, id int32) (err error) {
	if err = DBConfig.DB.Where("id = ?", id).First(campaign).Error; err != nil {
		return err
	}
	return nil
}

func GetCampaignByName(campaign *Campaign, name string) (err error) {
	if err = DBConfig.DB.Where("name = ?", name).First(campaign).Error; err != nil {
		return err
	}
	return nil
}
