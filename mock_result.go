package tst

type Result struct {
	rowsAffected int64
	lastInsertID int64
	err          error
}

func NewResult(rowsAffected, lastInsertID int, err error) *Result {
	return new(Result{rowsAffected: int64(rowsAffected), lastInsertID: int64(lastInsertID), err: err})
}

func (r *Result) LastInsertId() (int64, error) {
	return r.lastInsertID, r.err
}

func (r *Result) RowsAffected() (int64, error) {
	return r.rowsAffected, r.err
}
