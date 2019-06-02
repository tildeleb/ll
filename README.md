# lisplists (lispy lists in Go)

Copyright © 2019 Lawrence E. Bakst. All rights reserved.

Unlike languages such as Python and Ruby, Go does not have a built in list type.

Package lispylist implements a classical lisp list data structure based on a Cons cell or Pair. 

The basic data structure here is a "List". sometimes called a "cons" or "pair" which has two cells, a Head and a Tail. In classic lisp the head of a list is accessed via the "car" fucntion and the Tail is accessed via the "cdr" function. A Pair is often written as a dotted Head "(a . b)" where a is the Head and b is the Tail. Lists are build up from dotted Pairs via cons, with the last element of the list having nil as it's Tail.

    MakeList(1, 2, 3, 4) === (1 . (2 . (3 . (4. nil)))) == Cons(1, Cons(2, Cons(3, Cons(4, nil))))
    
In summary, list.Head contains a list element and the list.Tail contains a pointer to the next list element. Most of the functions here are just straight functions and not methods. Currently the sole exception is Print for which a convenience