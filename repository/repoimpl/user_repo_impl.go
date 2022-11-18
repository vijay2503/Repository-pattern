package repoimpl

import (
	"database/sql"
	"errors"
	"fmt"
	models "go-postgres/model"
	repo "go-postgres/repository"
	m "net/mail"
	"strconv"
)

type UserRepoImpl struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) repo.UserRepo {
	return &UserRepoImpl{
		Db: db,
	}
}
func (u *UserRepoImpl) Create(Userst models.User)error {
	err := Insert(u, Userst)
	if err != nil {
		return err
	}
return nil
}

func (u *UserRepoImpl) Select() (users []models.User) {
	rows, err := u.Db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Email)
		if err != nil {
			break
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	if len(users) == 0 {
		fmt.Println("The Table is Empty")
		return
	}
	return users
}

func Insert(u *UserRepoImpl, user models.User) error {
	insertStatement := `
	INSERT INTO users (id, name, gender, email)
	VALUES ($1, $2, $3, $4)`
	_, err := u.Db.Exec(insertStatement, user.ID, user.Name, user.Gender, user.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Record added: ", user)
	return nil
}
func (u *UserRepoImpl) Delete(delId int) error {
	deleteStmt := `Delete from users where id=` + strconv.Itoa(delId)
	sqlresult, err := u.Db.Exec(deleteStmt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	affectRows, rr := sqlresult.RowsAffected()
	if rr != nil {
		fmt.Print(rr)
		return rr
	}
	if affectRows == 0 {
		return errors.New("User-Id Not found")
	}
	return nil
}
func (u *UserRepoImpl) Update(userId *int, user *models.User) error {
	use := models.User{}
	rows, err := u.Db.Exec("SELECT * FROM users where id=" + strconv.Itoa(*userId))
	if err != nil {
		fmt.Println(err)
		return err
	}
	n, _ := rows.RowsAffected()
	if n == 0 {
		fmt.Println("Invalid User Id ")
		return errors.New("Invalid User-ID")
	}
	mail, Err := u.Db.Query("SELECT * FROM users where id=$1", *userId)
	if Err != nil {
		return Err
	}
	for mail.Next() {
		err := mail.Scan(&use.ID, &use.Name, &use.Gender, &use.Email)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	if use.Email == user.Email {
		fmt.Println(" Previous mail & updated mail are same")
		return errors.New("Updation Mail and Previous Mail are same")
	}
	fmt.Println("line 121", user.Email)
	_, err = m.ParseAddress(user.Email)
	if err != nil {
		fmt.Println("Invalid Email")
		return errors.New("In-valid mail")
	}
	updateStmt := `update "users" set "email"=$1 where "id"=$2`
	_, e := u.Db.Exec(updateStmt, user.Email, *userId)
	if e != nil {
		return e
	}
	fmt.Println("Successfully Updated")
	return nil
}
