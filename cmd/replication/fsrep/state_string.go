// Code generated by "stringer -type=State"; DO NOT EDIT.

package fsrep

import "strconv"

const (
	_State_name_0 = "ReadyRetryWait"
	_State_name_1 = "PermanentError"
	_State_name_2 = "Completed"
)

var (
	_State_index_0 = [...]uint8{0, 5, 14}
)

func (i State) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _State_name_0[_State_index_0[i]:_State_index_0[i+1]]
	case i == 4:
		return _State_name_1
	case i == 8:
		return _State_name_2
	default:
		return "State(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
