package attendant

import "time"

type HrEmployee struct {
	SerialID  uint      `json:"SerialID"`
	EmplID    string    `json:"EmplID"`
	CardID    string    `json:"CardID"`
	DeptID    string    `json:"DeptID"`
	EmplName  string    `json:"EmplName"`
	Sex       string    `json:"Sex"`
	EntryDate time.Time `json:"EntryDate"`
}

func (he *HrEmployee) TableName() string {
	return "HrEmployee"
}

type HrAttendantRecord struct {
	SerialID uint   `json:"SerialID"`
	EmplID   string `json:"EmplID"`
	RecDate  string `json:"RecDate"`
	RecTime  string `json:"RecTime"`
	OperDate string `json:"OperDate"`
}

func (hr *HrAttendantRecord) TableName() string {
	return "AtdRecord"
}
