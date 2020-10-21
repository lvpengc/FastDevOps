package foundation

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_ge(t *testing.T)  {

	d,err:=GenerateAvatar(122,122,"121")
	fmt.Println(string(d))
	//os.Create("test.png")
	ioutil.WriteFile("test1.png",d,0733)
	fmt.Println(err)

}
