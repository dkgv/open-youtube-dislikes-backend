package database

func IsNoRowError(err error) bool {
	return err != nil && err.Error() == "sql: no rows in result set"
}
