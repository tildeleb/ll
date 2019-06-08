// Copyright © 2019 Lawrence E. Bakst. All rights reserved.

package ll

// Nconc is horrible old lispy name so I cloned it as "splice" for now.
func Splice(x, y List) List {
	return Nconc(x, y)
}
