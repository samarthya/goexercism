package erratum

// Use Open the resource
func Use(o ResourceOpener, input string) (err error) {
	// Iniatilize the value and error if any.
	v, err := o()

	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		v, err = o()
	}

	defer v.Close()

	defer func() {
		if er1 := recover(); er1 != nil {
			if frob, ok := er1.(FrobError); ok {
				v.Defrob(frob.defrobTag)
			}
			err = er1.(error)
		}
	}()

	v.Frob(input)
	return err
}
