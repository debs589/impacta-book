package application

type Application interface {
	TearDown()
	SetUp() (err error)
	Run() (err error)
}
