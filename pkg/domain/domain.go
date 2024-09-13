package domain

type Timeblock struct {
	title string
}

func NewTimeblock(title string) (*Timeblock) {
	return &Timeblock{title}
}
