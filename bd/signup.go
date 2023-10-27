package bd

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goestoreuser/models"
	"goestoreuser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Start of the registration")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	instant := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FetchMySQL() + "')"
	fmt.Println(instant)

	_, err = Db.Exec(instant)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Successful execution")
	return nil

}
