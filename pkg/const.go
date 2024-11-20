package pkg

type UserRole string

const (
	Owner    UserRole = "owner"
	Admin    UserRole = "admin"
	Operator UserRole = "operator"
	Provider UserRole = "provider"
	Manager  UserRole = "manager"
)
