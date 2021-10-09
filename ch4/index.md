# Composite Types

composite type: created by combining the basic types in various way.
    arrays, slices, maps, and structs

aggregate type: values are concatenations of otther values in memory.
    array: homogeneous -- elements all have the same type.
    struct: heterogeneous
    both array and struct are fixed size
    slices and maps are dynamic data structure that grow as values are added

array: fixed-length sequence of zero or more element of a particular type.
    rarely used directly in Go. But to understand slice we must understand arrays first.

