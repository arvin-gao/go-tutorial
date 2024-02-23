package packages

import (
	"crypto/rand"
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	// The bcrypt cost range. (set default) < 4 <= Cost >=31 > (error), default = 10
	BcryptCost int
}

type User struct {
	Email          string
	HashedPassword []byte
	Salt           []byte
}

func TestCreateUserWithBcryptedPassword(t *testing.T) {
	conf := Config{
		BcryptCost: 3,
	}
	assert.Nil(t, CreateUser(&conf, "test@example.com", "password"))
}

func TestDe(t *testing.T) {
	// get from request input.
	plaintextPassword := "password"

	// get from database.
	cases := []struct {
		salt           []byte
		hashedPassword []byte
		hasErr         bool
	}{
		{
			salt:           []byte{78, 30, 26, 95, 92, 25, 131, 152, 243, 240},
			hashedPassword: []byte("$2a$10$H0k/TR6fK8DnMyxyGmz/WecclyLXmM1pk/XgtwkC8QaAcNpMw2JWq"),
			hasErr:         false,
		},
		{
			salt:           []byte{78, 30, 26, 95, 92, 25, 131, 152, 243, 240},
			hashedPassword: []byte("$2a$10$H0k/TR6fK8DnMyxyGmz/WecclyLXmM1pk/XgtwkC8QaAcNpMw2JWg"),
			hasErr:         true,
		},
	}

	for _, v := range cases {
		pwd := append(v.salt, []byte(plaintextPassword)...)
		err := bcrypt.CompareHashAndPassword(v.hashedPassword, pwd)

		if v.hasErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func CreateUser(conf *Config, email string, password string) error {
	// TODO: Check email is exists.
	salt, err := createSalt(10)
	if err != nil {
		return errors.New("create salt failed")
	}

	pwd := append(salt, []byte(password)...)
	hashPwd, _ := bcrypt.GenerateFromPassword(pwd, conf.BcryptCost)

	// TODO: Save a user metadata.
	saveUser := func(user *User) error {
		return nil
	}
	return saveUser(&User{
		Email:          email,
		HashedPassword: hashPwd,
		Salt:           salt,
	})
}

func createSalt(count int) ([]byte, error) {
	salt := make([]byte, count)
	_, err := rand.Read(salt)
	logrus.Info("salt:", salt)
	return salt, err
}

func Login(email string, password []byte) (token []byte, err error) {
	getUserByEmail := func(email string) *User {
		return &User{}
	}
	user := getUserByEmail(email)

	// Check the user password is corrected on there.

	if bcrypt.CompareHashAndPassword(user.HashedPassword, password) != nil {
		return nil, errors.New("password is not match")
	}

	// Make a token to response and returned on there.

	return
}
