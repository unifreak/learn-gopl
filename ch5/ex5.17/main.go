// @todo

// Write a variadic function ElementsByTagName that, given an HTML node tree and
// zero or more names, returns all the elments that match one of those names. Here
// are two example calls:
//
// 		func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
//
// 		images := ElementsByTagName(doc, "img")
// 		headings := ElmentsByTagName(doc, "h1", "h2", "h3", "h4")