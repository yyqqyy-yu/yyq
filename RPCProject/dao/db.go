package dao

func NewDB() (db *Orm, cf func(), err error) {

	db = Minit("dao/db.yaml")
	cf = func() { _ = db.Close() }
	return

}
