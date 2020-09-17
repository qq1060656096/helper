package helper

import "errors"

var (
	ErrRowsScanDataAssertTypeNil = errors.New("helper: ErrNilRowsScanDataAssertType")
	ErrRowsScanDataAssertTypeColumnsNil = errors.New("helper: ErrNilRowsScanDataAssertType.parameter.columns.nil")
	ErrRowsScanDataAssertTypeSrcMpNil = errors.New("helper: ErrNilRowsScanDataAssertType.parameter.srcMp.nil")
	ErrRowsStringDataAssertTypeNil = errors.New("helper: ErrRowsStringDataAssertTypeNil")
)