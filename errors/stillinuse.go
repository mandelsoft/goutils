package errors

type StillInUseError struct {
	errinfo
}

var formatStillInUseError = NewDefaultFormatter("is", "still in use", "for")

func ErrStillInUse(spec ...string) error {
	return &StillInUseError{newErrInfo(formatStillInUseError, spec...)}
}

func ErrStillInUseWrap(err error, spec ...string) error {
	return &StillInUseError{wrapErrInfo(err, formatStillInUseError, spec...)}
}

func IsErrStillInUse(err error) bool {
	return IsA(err, &StillInUseError{})
}

func IsErrStillInUseKind(err error, kind string) bool {
	var uerr *StillInUseError
	if err == nil || !As(err, &uerr) {
		return false
	}
	return uerr.kind == kind
}
