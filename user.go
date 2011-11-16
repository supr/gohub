package gohub

func (u *User) Equals(w User) bool {
	return u.Login == w.Login
}
