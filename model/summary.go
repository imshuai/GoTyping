package model

//ArticalSummary 定义文章摘要类
type ArticalSummary struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	Summary    string   `json:"summary"`
	Category   Category `json:"category"`
	CreateTime string   `json:"create_time"`
}

//GenSummary 生成文章摘要索引
func (atl *Artical) GenSummary() ArticalSummary {
	return ArticalSummary{
		ID:      atl.ID,
		Name:    atl.Name,
		Summary: atl.Summary,
		Category: func() Category {
			cat := new(Category)
			db.ID(atl.Category).Get(cat)
			return *cat
		}(),
		CreateTime: atl.CreateTime.Format("Monday, 02 Jan 2006"),
	}
}

//GetSummariesWithPageID 根据文章索引页获得摘要
func GetSummariesWithPageID(pid int) []ArticalSummary {
	count := (new(Artical)).Count()
	start := 0
	if start = (pid - 1) * 12; start < 0 || int64(start) > count {
		start = 0
	}
	atls := make([]*Artical, 0)
	db.Asc("id").Limit(12, start).Find(&atls)
	atlsums := make([]ArticalSummary, 0)
	for _, v := range atls {
		atlsums = append(atlsums, v.GenSummary())
	}
	return atlsums
}
