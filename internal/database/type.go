package database

type detectServices struct {
	Id       int    `gorm:"column:id;AUTO_INCREMENT"`
	Url      string `json:"url"`
	Status   int    `json:"status"`
	IsOnline int    `json:"is_online"`
}

type results struct {
	Id         int    `gorm:"column:id;AUTO_INCREMENT"`
	PicUrl     string `json:"pic_url"`
	UpdateTime string `json:"update_time"`
	Results    string `json:"results"`
	Count0     int    `json:"count0"`
	Count1     int    `json:"count1"`
	Count2     int    `json:"count2"`
	Count3     int    `json:"count3"`
	Count4     int    `json:"count4"`
}
