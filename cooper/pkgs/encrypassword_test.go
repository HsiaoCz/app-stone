package pkg

import "testing"

func TestEncryPassword(t *testing.T) {
   password:=EncryPassword("asdf1234")
   t.Logf("password hash : %v\n",password)
}
