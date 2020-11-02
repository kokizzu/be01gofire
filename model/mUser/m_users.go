package mUser

import (
	"cloud.google.com/go/firestore"
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
)

const Collection = `users`

type User struct {
	Email string 
	Pass  string 
}

func (u *User) Insert(c *firestore.Client) error {
	users := c.Collection(Collection)
	user := users.Doc(hashOf(u.Email))
	snap, err := user.Get(context.Background())
	if err == nil  && snap.Exists() {
		return errors.New(`user already exists`)
	}
	_, err = user.Set(context.Background(), u)
	return err
}

func (u *User) FindByEmail(c *firestore.Client) error {
	users := c.Collection(Collection)
	user := users.Doc(hashOf(u.Email))
	snap, err := user.Get(context.Background())
	fmt.Println(err)
	if err != nil {
		return err
	}
	return snap.DataTo(u)
}

func hashOf(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (u *User) HashPassword() {
	u.Pass = hashOf(u.Pass)
}

func (u *User) CheckPass(pass string) bool {
	return u.Pass == hashOf(pass)
}
