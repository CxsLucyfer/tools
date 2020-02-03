package packages

import (
	"github.com/blang/semver"
)

// ByVersion implements sort.Interface for []Asset based on
// the Version field.
type ByVersion []Asset

func (a ByVersion) Len() int      { return len(a) }
func (a ByVersion) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByVersion) Less(i, j int) bool {
	left, leftErr := semver.Make(a[i].Version)
	if leftErr != nil {
		return false
	}
	right, rightErr := semver.Make(a[j].Version)
	if rightErr != nil {
		return true
	}
	return left.Compare(right) == 1
}