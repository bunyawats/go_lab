package doer

type Doer interface {
	DoSomething(int, string) error
	SaySomething(string2 string) string
}
