// USERNAME, GENDER,SKINTYPE, SKINCOLOR, DESC_REVIEW, JUMLAH_LIKE_REVIEW
package response

import "kstyle-test/entity"

type Reviews struct {
	Username         string `json:"username"`
	Gender           string `json:"gender"`
	Skintype         string `json:"skintype"`
	Skincolor        string `json:"skincolor"`
	DescReview       string `json:"desc_review"`
	JumlahLikeReview int    `json:"jumlah_like_review"`
}
type Product struct {
	ID          uint      `json:"id"`
	NameProduct string    `json:"name_product"`
	Price       float64   `json:"price"`
	Reviews     []Reviews `json:"reviews"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	DeletedAt   string    `json:"deleted_at"`
}

func BuildProducts(products []entity.Product, list []entity.ReviewLikesCount) (res []Product) {
	for _, v := range products {
		product := Product{
			ID:          v.ID,
			NameProduct: v.NameProduct,
			Price:       v.Price,
			CreatedAt:   v.CreatedAt.String(),
			UpdatedAt:   v.UpdatedAt.String(),
			DeletedAt:   v.DeletedAt.Time.String(),
		}
		for _, rev := range v.Reviews {
			review := Reviews{
				Username:   rev.Member.Username,
				Gender:     rev.Member.Gender,
				Skintype:   rev.Member.Skintype,
				Skincolor:  rev.Member.Skincolor,
				DescReview: rev.DescReview,
			}
			for _, revlike := range list {
				if revlike.ReviewID == rev.ID {
					review.JumlahLikeReview = revlike.LikesCount
				}
			}
			product.Reviews = append(product.Reviews, review)
		}

		res = append(res, product)
	}
	return res
}

func BuildProduct(product entity.Product, list []entity.ReviewLikesCount) (res Product) {
	temp := Reviews{}
	res.ID = product.ID
	res.NameProduct = product.NameProduct
	res.Price = product.Price
	for _, v := range product.Reviews {
		temp.Username = v.Member.Username
		temp.Gender = v.Member.Gender
		temp.Skintype = v.Member.Skintype
		temp.Skincolor = v.Member.Skincolor
		temp.DescReview = v.DescReview
		for _, v1 := range list {
			if v.ID == v1.ReviewID {
				temp.JumlahLikeReview = v1.LikesCount
			}
		}
		res.Reviews = append(res.Reviews, temp)
	}
	res.CreatedAt = product.CreatedAt.String()
	res.UpdatedAt = product.UpdatedAt.String()
	res.DeletedAt = product.DeletedAt.Time.String()

	return res
}
