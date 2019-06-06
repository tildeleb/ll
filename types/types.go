package lispylist

import (
	"fmt"
	"math/rand"
	"time"
)

// 4 choices for implmenting a lispy like list datastructure
// Has anyone else run into this issue?

// Option 1 is OK. The code would use "Pair *" as the type for lists. There is no separate type for lists.
// This is what I started with but I think it would be a good idea to have separate types for Pairs and Lists.
type Pair struct {
	Head interface{}
	Tail interface{}
}

// Option 2 is what I switched to but this struct is not really a List, it's a Pair or Cons.
type List struct {
	Head interface{}
	Tail interface{}
}

// Option 3 is OK, but now we can't make metods on List because it's a named pointer type.
// I don't currently use methods on List because it's a functional design but it would be nice to have the option.
type Pair struct {
	Head interface{}
	Tail interface{}
}
type List *Pair

// Option 4 is what I am thinking of switching to. It uses a type alias and now I can make methods on List.
type Pair struct {
	Head interface{}
	Tail interface{}
}
type List = *Pair
