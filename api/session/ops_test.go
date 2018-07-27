package session

import (
	"fmt"
	"testing"
)

var tempsid string

func TestMain(m *testing.M) {
	m.Run()
}

func TestSessionWorkFlow(t *testing.T) {
	t.Run("InsertSessionId", testGenerateNewSessionId)
	t.Run("LoadAllSession", testLoadSessionFromDB)
	t.Run("ExpiredSession", testIsSessionExpired)
}

func testGenerateNewSessionId(t *testing.T) {
	sid := GenerateNewSessionId("skyone")
	tempsid = sid
	fmt.Printf("sid :%s\n", sid)
}

func testLoadSessionFromDB(t *testing.T) {
	m := LoadSessionsFromDB()
	m.Range(func(k, v interface{}) bool {
		fmt.Printf("key= %s, value=%s\n", k, v)
		return true
	})
}

func testIsSessionExpired(t *testing.T) {
	username, isExpired := IsSessionExpired(tempsid)
	fmt.Printf("username = %s, isExpired = %t\n", username, isExpired)
}
