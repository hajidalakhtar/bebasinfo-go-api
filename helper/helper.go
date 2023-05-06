package helper

import (
	"bebasinfo/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func PasswordEncrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func PasswordCompare(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

//func ToPaginatedResponse(TotalItems int64, TotalPages int, CurrentPage int, NextPage int, PrevPage int) domain.PaginatedResponse {
//	paginate := domain.PaginatedResponse{
//		TotalItems:  TotalItems,
//		TotalPages:  TotalPages,
//		CurrentPage: CurrentPage,
//		NextPage:    NextPage,
//		PrevPage:    PrevPage,
//	}
//	return paginate
//}

func ToUserResponses(users []domain.User) []domain.UserResp {
	var userResponses []domain.UserResp
	for _, user := range users {
		userResp := ToUserResponse(user)
		userResponses = append(userResponses, userResp)

	}
	return userResponses

}

func ToUserResponse(user domain.User) domain.UserResp {
	userResp := domain.UserResp{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
	return userResp
}

func StringToUUIDS(bulkID []string) []uuid.UUID {
	var uuids []uuid.UUID
	for _, id := range bulkID {
		uuid, _ := uuid.Parse(id)
		uuids = append(uuids, uuid)
	}
	return uuids
}

func ToEmptyArray(input interface{}) interface{} {
	if input == nil {
		return []string{}
	}
	return input
}
