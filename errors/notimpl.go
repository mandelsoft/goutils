package errors

type NotImplementedError struct {
	errinfo
}

var formatNotImplemented = NewDefaultFormatter("", "not implemented", "by")

func ErrNotImplemented(spec ...string) error {
	return &NotImplementedError{newErrInfo(formatNotImplemented, spec...)}
}

func IsErrNotImplemented(err error) bool {
	return IsA(err, &NotImplementedError{})
}

func IsErrNotImplementedKind(err error, kind string) bool {
	var uerr *NotImplementedError
	if err == nil || !As(err, &uerr) {
		return false
	}
	return uerr.kind == kind
}
