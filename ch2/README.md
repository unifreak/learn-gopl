# Names

Case matters.

25 `keywords`. Can't be used as names:

    break       default     func    interface   select
    case        defer       go      map         struct
    chan        else        goto    package     switch
    const       fallthrough if      range       type
    continue    for     import  return      var

`predeclared` names. Can be redeclared when make sense:
- Constants:    true false iota nil
- Types:        int int8 int16 int32 int64
                uint uint8 uint16 uint32 uint64 uintptr
                float32 float64 complex128 complex64
                bool type rune string error
- Functions:    make len cap new append copy close delete
                complex real imag
                panic recover

Go prefer short names, camel case, and letters of acronyms and initialisms are always in the same case. Hence `htmlEscape` or `escapeHTML` is okay, but NOT `escapeHtml`.

# Declarations

`declaration`: names a program entity and specifies some or all of its properties. like `var`, `const`, `type` and `func`.

