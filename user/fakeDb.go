package user

import (
	"sort"

	"playgrounds.com/utils/utils"
)

type Database struct {
	Users map[string]User
}

func NewDb() *Database {
	return &Database{
		Users: map[string]User{},
	}
}

func (d *Database) GetAll() []User {
	users := []User{}
	keys := utils.KeyList(d.Users)
	sort.Strings(keys)

	for k := range keys {
		users = append(users, d.Users[keys[k]])
	}
	return users
}

func (d *Database) Get(id string) *User {
	if user, ok := d.Users[id]; ok {
		return &user
	}
	return nil
}

func (d *Database) Create(user User) {
	d.Users[user.ID] = user
}

func (d *Database) Update(id string, user User) {
	d.Users[id] = user
}

func (d *Database) Delete(id string) {
	delete(d.Users, id)
}
