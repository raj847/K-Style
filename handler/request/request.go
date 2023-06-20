package request

type Review struct {
	IDProduct  int    `json:"id_product"`
	IDMember   int    `json:"id_member"`
	DescReview string `json:"desc_review"`
}
