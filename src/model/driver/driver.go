package driver

type DriverInfo struct {
	Name string `json:"name"`
	Id   int    `json:"id" xorm:"autoincr pk"`
}

func (this *DriverInfo) TableName() string {
	return "g_driver"
}
