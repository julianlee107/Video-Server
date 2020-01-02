package dbops

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func clearTables() {
	dbConn.Exec("truncate user")
}

func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	crypt := sha256.New()
	crypt.Write([]byte("123"))
	newPwd := fmt.Sprintf("%x",crypt.Sum(nil))
	if pwd != newPwd || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}