// Code generated by "stringer -type Membership"; DO NOT EDIT.

package checkgroup

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MembershipUnknown-0]
	_ = x[IsMember-1]
	_ = x[NotMember-2]
}

const _Membership_name = "MembershipUnknownIsMemberNotMember"

var _Membership_index = [...]uint8{0, 17, 25, 34}

func (i Membership) String() string {
	if i < 0 || i >= Membership(len(_Membership_index)-1) {
		return "Membership(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Membership_name[_Membership_index[i]:_Membership_index[i+1]]
}