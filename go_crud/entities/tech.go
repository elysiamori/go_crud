package entities

/*Define Variables*/
type Tech struct {
	Id          int64
	Name        string `validate:"required"`
	Job         string `validate:"required"`
	Programming string `validate:"required" label:"programming"`
	Date        string `validate:"required"`
}
