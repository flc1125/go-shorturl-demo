package models

type MarketCommunity struct {
	ID        uint   `gorm:"primaryKey"`
	Type      int    `json:"type"`
	Acid      int    `json:"acid"`
	Regionid  int    `json:"regionid"`
	Sectionid int    `json:"sectionid"`
	Stationid string `json:"stationid"`
	Name      string `json:"name"`
}

type Tabler interface {
	TableName() string
}

func (MarketCommunity) TableName() string {
	return "market_community"
}
