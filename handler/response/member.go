package response

import "kstyle-test/entity"

type Member struct {
	MemberID  uint   `json:"id_member"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	Skintype  string `json:"skintype"`
	Skincolor string `json:"skincolor"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func BuildMember(member entity.Member) Member {
	return Member{
		MemberID:  member.ID,
		Username:  member.Username,
		Gender:    member.Gender,
		Skintype:  member.Skintype,
		Skincolor: member.Skincolor,
		CreatedAt: member.CreatedAt.String(),
		UpdatedAt: member.UpdatedAt.String(),
		DeletedAt: member.DeletedAt.Time.String(),
	}
}

func BuildMembers(members []entity.Member) (res []Member) {
	for _, v := range members {
		member := BuildMember(v)
		res = append(res, member)
	}
	return res
}
