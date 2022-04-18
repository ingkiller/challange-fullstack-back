package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ingkiller/hackernews/graph/model"
	"github.com/ingkiller/hackernews/internal/client"
	"github.com/ingkiller/hackernews/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Id       int
	Name     string
	Username string
	Password string
	Website  string
	Email    string
}

type DataUser struct {
	User  User
	Token string
}

func GetAll() []*model.User {
	resp, err := client.MakeReq("https://jsonplaceholder.typicode.com/users")
	defer resp.Body.Close()
	if err != nil {
		fmt.Print("NewRequest: %v", err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var resObject []User
	json.Unmarshal(bodyBytes, &resObject)
	var result []*model.User
	for index, user := range resObject {
		result = append(result, &model.User{
			ID:       &resObject[index].Id,
			Name:     user.Name,
			Username: user.Username,
			Website:  user.Website,
			Email:    user.Email,
		})
	}
	return result
}

func (u User) GetUserIdByUsername(username string) (int, error) {
	return u.Id, nil
}

func GetUserById(id int) User {
	client := &http.Client{}
	userUrl := fmt.Sprint("https://jsonplaceholder.typicode.com/users/", id)
	req, err := http.NewRequest(http.MethodGet, userUrl, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("NewRequest: %v", err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var resObject User
	json.Unmarshal(bodyBytes, &resObject)
	return resObject
}

func (user *User) Authenticate() (DataUser, error) {

	if user.Username == user.Password {
		userUrl := fmt.Sprint("https://jsonplaceholder.typicode.com/users?username=", user.Username)
		resp, err := client.MakeReq(userUrl)
		if err != nil {
			fmt.Print("NewRequest: %v", err.Error())
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		var userResult []User
		json.Unmarshal(bodyBytes, &userResult)
		if len(userResult) > 0 {
			token, err := jwt.GenerateToken(user.Username)
			if err != nil {
				return DataUser{}, errors.New("no username")
			}
			return DataUser{Token: token, User: userResult[0]}, nil
		}
	}

	return DataUser{}, errors.New("no username")
}
func GetUserIdByUsername(username string) (int, error) {
	/*
		statement, err := database.Db.Prepare("select ID from Users WHERE Username = ?")
		if err != nil {
			log.Fatal(err)
		}
		row := statement.QueryRow(username)

		var Id int
		err = row.Scan(&Id)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Print(err)
			}
			return 0, err
		}

		return Id, nil
	*/
	return 1, nil
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}
