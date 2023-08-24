package validation

import "github.com/liquiddev99/komikan-backend/pb"

func ValidateCreateUserRequest(req *pb.CreateUserRequest) (err error) {
	if err = validateFullName(req.GetFullName()); err != nil {
		return
	}
	if err = validateUsername(req.GetUsername()); err != nil {
		return
	}
	if err = validateEmail(req.GetEmail()); err != nil {
		return
	}
	if err = validatePassword(req.GetPassword()); err != nil {
		return
	}

	return nil
}
