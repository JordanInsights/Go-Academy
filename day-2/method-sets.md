# Method Sets
Methods, attached to a type, are the type's 'Method Set'. Whether you use pointers, or a normal type (i.e. not a pointer to a type) determine which methods attach to the given type. 

- The method set of a pointer to a defined type T (where T is neither a pointer nor an interface) is the set of all methods declared with receiver *T or T.

- The method set of an interface type is the intersection of the method sets of each type in the interface's type set (the resulting method set is usually just the set of declared methods in the interface). 

Further rules apply to structs containing anonymous fields. Any other type has an empty method set. In a method set, each method must have a unique non-blank method name. 


## THE ONLY SET OF VALUES THAT WON'T WORK TOGETHER ARE POINTER RECEIVER AND NON-POINTER VALUE
