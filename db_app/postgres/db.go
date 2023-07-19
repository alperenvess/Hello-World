package postgres

type PsqlStore struct {
}

func (ps *PsqlStore) GetEmployee(id int) string {
	return "Psql"
}
