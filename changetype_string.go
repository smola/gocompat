// Code generated by "stringer -type=ChangeType"; DO NOT EDIT.

package compat

import "strconv"

const _ChangeType_name = "TypeChangedFieldAddedFieldDeletedFieldChangedTypeSignatureChangedMethodAddedMethodDeletedMethodSignatureChangedInterfaceChanged"

var _ChangeType_index = [...]uint8{0, 11, 21, 33, 49, 65, 76, 89, 111, 127}

func (i ChangeType) String() string {
	i -= 1
	if i < 0 || i >= ChangeType(len(_ChangeType_index)-1) {
		return "ChangeType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ChangeType_name[_ChangeType_index[i]:_ChangeType_index[i+1]]
}
