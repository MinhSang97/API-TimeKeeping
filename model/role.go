package model

type Role int

const (
	ADMIN Role = iota
	MEMBER
	ADMIN1
	ADMIN2
)

func (r Role) String() string {
	return []string{"ADMIN", "MEMBER"}[r]
}
