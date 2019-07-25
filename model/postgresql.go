package model

type ABanks struct {
	IID      string `gorm:"Column:iid"`
	Name     string `gorm:"Column:name"`
	NickName string `gorm:"Column:nickname"`
}

func (ABanks) TableName() string {
	return "banks"
}

type ACards struct {
	BankID   string `gorm:"Column:bankid"`
	Name     string `gorm:"Column:name"`
	NickName string `gorm:"Column:nickname"`
}

func (ACards) TableName() string {
	return "cards"
}

type AEpays struct {
	Name      string `gorm:"Column:name"`
	NickName  string `gorm:"Column:nickname"`
	ImageName string `gorm:"Column:imagename"`
}

func (AEpays) TableName() string {
	return "epays"
}

type AScenarios struct {
	Name     string `gorm:"Column:name"`
	NickName string `gorm:"Column:nickname"`
}

func (AScenarios) TableName() string {
	return "scenarios"
}

type ATickets struct {
	Name      string `gorm:"Column:name"`
	NickName  string `gorm:"Column:nickname"`
	ShortName string `gorm:"Column:shortname"`
}

func (ATickets) TableName() string {
	return "tickets"
}
