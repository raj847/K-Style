package request

type Member struct {
	Username  string `validate:"required" json:"username"`
	Gender    string `validate:"required" json:"gender"`
	Skintype  string `validate:"required" json:"skintype"`
	Skincolor string `validate:"required" json:"skincolor"`
}
