package GoLearn

import "fmt"

type GoUserHelper interface {
	GoUserGenerate()
	GoUserUpdate()
}

type AwemeGoUser struct {
	TestTag int8
}

type TikTokGoUser struct {
	TestTag int8
}

func (goUser AwemeGoUser) GoUserGenerate() {
	fmt.Println("aweme gen")
}

func (goUser AwemeGoUser) GoUserUpdate() {
	fmt.Println("aweme  update")
	fmt.Println(goUser.TestTag)
}

func (goUser TikTokGoUser) GoUserGenerate() {
	fmt.Println("TT gen")
}

func (goUser TikTokGoUser) GoUserUpdate() {
	fmt.Println("TT  update")
	fmt.Println(goUser.TestTag)
}

func BuildTccClient() GoUserHelper {
	a := 2
	if a == 1 {
		return TikTokGoUser{TestTag: 2}
	} else {
		return AwemeGoUser{TestTag: 1}
	}
}
