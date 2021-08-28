package model

type PaginationCondition struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type TopicListCondition struct {
	PaginationCondition
	//time_format:"unixNano" msgpack:"createTime"
	//Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1" json:"birthday" url:"birthday" msgpack:"birthday"`
	CreateTime int64 `json:"createTime" form:"createTime"` // 创建时间 时间戳
	NodeId     int64 `json:"nodeId"`                       // 类别
	Recommend  bool  `json:"recommend"`                    // 是否推荐
	Cursor     int64 `json:"cursor"`
}
