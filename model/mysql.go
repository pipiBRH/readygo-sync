package model

import "time"

type Banks struct {
	ID           int     `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	BankCode     string  `gorm:"UNIQUE_INDEX;NOT NULL"`
	BankName     string  `gorm:"NOT NUL"`
	BankNickname string  `gorm:"NOT NUL"`
	Cards        []Cards `gorm:"ForeignKey:BankID;AssociationForeignKey:ID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Banks) TableName() string {
	return "banks"
}

type Cards struct {
	ID           int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	BankID       int    `gorm:"INDEX:bank_id;NOT NULL;UNIQUE_INDEX:unique_cardname"`
	CardName     string `gorm:"NOT NULL;UNIQUE_INDEX:unique_cardname"`
	CardNickname string `gorm:"NOT NULL"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Cards) TableName() string {
	return "cards"
}

type Epays struct {
	ID           int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	EpayName     string `gorm:"UNIQUE_INDEX;NOT NULL"`
	EpayNickname string `gorm:"NOT NULL"`
	ImageName    string `gorm:"NOT NULL"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Epays) TableName() string {
	return "epays"
}

type Scenarios struct {
	ID               int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	ScenarioName     string `gorm:"UNIQUE_INDEX;NOT NULL"`
	ScenarioNickname string `gorm:"NOT NULL"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (Scenarios) TableName() string {
	return "scenarios"
}

type Tickets struct {
	ID              int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	TicketName      string `gorm:"UNIQUE_INDEX;NOT NULL"`
	TicketNickname  string `gorm:"NOT NULL"`
	TicketShortname string `gorm:"NOT NULL"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (Tickets) TableName() string {
	return "tickets"
}
