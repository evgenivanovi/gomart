package auth

/* __________________________________________________ */

type Credentials struct {
	username string
	password string
	hashed   bool
}

func (c Credentials) Username() string {
	return c.username
}

func (c Credentials) Password() string {
	return c.password
}

func (c Credentials) Hashed() bool {
	return c.hashed
}

func (c Credentials) WithPassword(password string) *Credentials {
	return credentials(c.Username(), password, false)
}

func (c Credentials) WithHashedPassword(password string) *Credentials {
	return credentials(c.Username(), password, true)
}

func (c Credentials) WithHash(hash func(string) string) *Credentials {
	return credentials(c.Username(), hash(c.password), true)
}

func NewCredentials(username string, password string) *Credentials {
	return credentials(username, password, false)
}

func NewHashedCredentials(username string, password string) *Credentials {
	return credentials(username, password, true)
}

func credentials(username string, password string, hashed bool) *Credentials {
	return &Credentials{
		username: username,
		password: password,
		hashed:   hashed,
	}
}

/* __________________________________________________ */
