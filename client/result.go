package client

// specification: https://julius.osdn.jp/juliusbook/ja/desc_module.html
type Result struct {
	Rank    string
	Score   string
	Gram    string
	Details []Detail
}

type Detail struct {
	Word    string
	ClassID string
	Phone   string
	CM      string
}
