package service

type Baser interface {
	GetID() int64
	GetVersion() int64
	Get(id int64) interface{}
	VersionComparison(data interface{}) bool
	IsExists(id int64) bool
	Inserted(data interface{}) error
	Deleted(id int64) bool
}

func UpdateCheck(base Base) (updated, inserted []int64) {
	if base.IsExists(base.GetID()) {
		if !base.VersionComparison(base.GetID()) {
			updated = append(updated, base.GetID())
		}
	} else {
		inserted = append(inserted, base.GetID())
	}
	return
}
