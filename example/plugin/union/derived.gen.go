// Code generated by goderive DO NOT EDIT.

package union

// deriveUnique returns a list containing only the unique items from the input list.
// It does this by reusing the input list.
func deriveUnique(list []*Person) []*Person {
	if len(list) == 0 {
		return nil
	}
	u := 1
	for i := 1; i < len(list); i++ {
		if !deriveContains(list[:u], list[i]) {
			if i != u {
				list[u] = list[i]
			}
			u++
		}
	}
	return list[:u]
}

// deriveFilter returns a list of all items in the list that matches the predicate.
func deriveFilter(predicate func(*Person) bool, list []*Person) []*Person {
	out := make([]*Person, 0, len(list))
	for i, elem := range list {
		if predicate(elem) {
			out = append(out, list[i])
		}
	}
	return out
}

// deriveUnion returns the union of the items of the two input lists.
// It does this by append items to the first list.
func deriveUnion(this, that []*Person) []*Person {
	for i, v := range that {
		if !deriveContains(this, v) {
			this = append(this, that[i])
		}
	}
	return this
}

// deriveContains returns whether the item is contained in the list.
func deriveContains(list []*Person, item *Person) bool {
	for _, v := range list {
		if deriveEqual(v, item) {
			return true
		}
	}
	return false
}

// deriveEqual returns whether this and that are equal.
func deriveEqual(this, that *Person) bool {
	return (this == nil && that == nil) ||
		this != nil && that != nil &&
			this.Name == that.Name &&
			((this.Vote == nil && that.Vote == nil) || (this.Vote != nil && that.Vote != nil && *(this.Vote) == *(that.Vote)))
}
