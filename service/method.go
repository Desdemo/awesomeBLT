package service

type Base interface {
	GetID() int64
	VersionComparison(id int64) bool
	IsExists(id int64) bool
}

type Demo struct {
	Id int64
}

func (demo *Demo) GetID() int64  {
	return demo.Id
}

func (demo *Demo) VersionComparison(id int64) bool  {
	return false
}

func (demo *Demo) IsExists(id int64) bool {
	return false
}



func UpdateCheck(base Base) (updated,inserted []int64){
	if base.IsExists(base.GetID()) {
		if !base.VersionComparison(base.GetID()) {
			updated = append(updated, base.GetID())
		}
	} else {
		inserted = append(inserted, base.GetID())
	}

	return
}



