package testrdio

import (
	"testing"
)

func TestAddFriend(t *testing.T) {
	c := createClient(t)

	_, err := c.AddFriend("")
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestApproveFollower(t *testing.T) {
	c := createClient(t)

	_, err := c.ApproveFollower("")
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestCurrentUser(t *testing.T) {
	c := createClient(t)

	user, err := c.CurrentUser()
	if err != nil {
		t.Fatal(err)
	}

	if user.FirstName != "Myles" {
		t.Errorf("Expected FirstName 'Myles', received '%s'", user.FirstName)
	}
	if user.LastName != "Grant" {
		t.Errorf("Expected LastName 'Grant', received '%s'", user.LastName)
	}
}

func TestFindUserEmail(t *testing.T) {
	c := createClient(t)

	user, err := c.FindUserEmail("myles@mylesgrant.com")
	if err != nil {
		t.Fatal(err)
	}

	if user.FirstName != "Myles" {
		t.Errorf("Expected FirstName 'Myles', received '%s'", user.FirstName)
	}
	if user.LastName != "Grant" {
		t.Errorf("Expected LastName 'Grant', received '%s'", user.LastName)
	}
}

func TestFindUserVanityName(t *testing.T) {
	c := createClient(t)

	user, err := c.FindUserVanityName("myles")
	if err != nil {
		t.Fatal(err)
	}

	if user.FirstName != "Myles" {
		t.Errorf("Expected FirstName 'Myles', received '%s'", user.FirstName)
	}
	if user.LastName != "Grant" {
		t.Errorf("Expected LastName 'Grant', received '%s'", user.LastName)
	}
}

func TestHideFollower(t *testing.T) {
	c := createClient(t)

	_, err := c.HideFollower("")
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestRemoveFriend(t *testing.T) {
	c := createClient(t)

	_, err := c.RemoveFriend("")
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestUnapproveFollower(t *testing.T) {
	c := createClient(t)

	_, err := c.UnapproveFollower("")
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestUserFollowers(t *testing.T) {
	c := createClient(t)

	_, err := c.UserFollowers("")
	if err == nil {
		t.Fatal("Expected API error, got none")
	}

	users, err := c.UserFollowers("s3318")
	if err != nil {
		t.Fatal(err)
	}

	if len(users) == 0 {
		t.Error("Expected to find followers, but got none")
	}
}

func TestUserFollowing(t *testing.T) {
	c := createClient(t)

	_, err := c.UserFollowing("")
	if err == nil {
		t.Fatal("Expected API error, got none")
	}

	users, err := c.UserFollowing("s3318")
	if err != nil {
		t.Fatal(err)
	}

	if len(users) == 0 {
		t.Error("Expected to find following users, but got none")
	}
}

func TestUserHiddenFollowers(t *testing.T) {
	c := createClient(t)

	_, err := c.UserHiddenFollowers()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserPendingFollowers(t *testing.T) {
	c := createClient(t)

	_, err := c.UserPendingFollowers()
	if err != nil {
		t.Fatal(err)
	}
}
