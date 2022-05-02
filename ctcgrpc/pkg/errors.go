package pkg

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

var (
	ErrNotFound = errors.New("not found")
)

func ErrorHttpStatusCode(err error) int {
	for err != nil {
		switch err {
		case ErrNotFound:
			return http.StatusNotFound
		}

		err = errors.Unwrap(err)
	}

	return http.StatusInternalServerError
}

func ToGrpcError(err error) error {
	if err == nil {
		return nil
	}

	// do not modify if already grpc status error
	if _, ok := status.FromError(err); ok {
		return err
	}

	switch ErrorHttpStatusCode(err) {
	case http.StatusConflict, http.StatusBadRequest:
		return status.Errorf(codes.FailedPrecondition, err.Error())
	case http.StatusNotFound:
		return status.Errorf(codes.NotFound, err.Error())
	case http.StatusTooManyRequests:
		return status.Errorf(codes.ResourceExhausted, err.Error())
	case http.StatusForbidden:
		return status.Errorf(codes.PermissionDenied, err.Error())
	case http.StatusUnauthorized:
		return status.Errorf(codes.Unauthenticated, err.Error())
	default:
		return status.Errorf(codes.Internal, err.Error())
	}

}

func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrNotFound)
}
