package pkg

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestGetAvatar(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
	picture := GetPicture("AVATAR")
	t.Logf("your avatar is : %v\n", picture)
}
