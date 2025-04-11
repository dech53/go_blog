package flag

import (
	"errors"
	"fmt"
	"os"
	"server/global"
	"server/model/appTypes"
	"server/model/database"
	"server/utils"
	"syscall"

	"github.com/gofrs/uuid"
	"golang.org/x/term"
)

func Admin() error {
	var user database.User

	fmt.Print("Enter email: ")
	var email string
	_, err := fmt.Scanln(&email)
	if err != nil {
		return fmt.Errorf("failed to read email: %w", err)
	}
	user.Email = email

	fd := int(syscall.Stdin)
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer term.Restore(fd, oldState)

	fmt.Print("Enter password: ")
	password, err := readPassword()
	fmt.Println()
	if err != nil {
		return err
	}

	fmt.Print("Confirm password: ")
	rePassword, err := readPassword()
	fmt.Println()
	if err != nil {
		return err
	}

	if password != rePassword {
		return errors.New("passwords do not match")
	}

	if len(password) < 8 || len(password) > 20 {
		return errors.New("password length should be between 8 and 20 characters")
	}

	user.UUID = uuid.Must(uuid.NewV4())
	user.Username = global.Config.Website.Name
	user.Password = utils.BcryptHash(password)
	user.RoleId = appTypes.Admin
	user.Avatar = "/image/avatar.jpg"
	user.Address = global.Config.Website.Address

	if err := global.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func readPassword() (string, error) {
	var password string
	var buf [1]byte

	for {
		_, err := os.Stdin.Read(buf[:])
		if err != nil {
			return "", err
		}
		char := buf[0]

		if char == '\n' || char == '\r' {
			break
		}

		password += string(char)
	}

	return password, nil
}
