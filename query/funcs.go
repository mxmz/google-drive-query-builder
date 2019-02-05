package query

func PropertiesHas(name string, value string) Statement { return &propertiesHasStatement{name, value} }

func Query(e Statement) QueryStatement { return &query{operatorDropSecond, e, nil} }

func FullText() FreeTextAttribute { return &fullTextAttribute{} }

func Name() TextAttribute { return &textAttribute{"name"} }

func MimeType() TextAttribute { return &textAttribute{"mimeType"} }

func ModifiedTime() TimeAttribute { return &timeAttribute{"modifiedTime"} }

func CreatedTime() TimeAttribute { return &timeAttribute{"createdTime"} }

func Stringize(s Statement) string { return s.toString() }

func Raw(v string) Statement { return &rawStatement{v} }

func Parents() CollectionAttribute { return &collectionAttributeAttribute{"parents"} }

func Owners() CollectionAttribute { return &collectionAttributeAttribute{"owners"} }

func Readers() CollectionAttribute { return &collectionAttributeAttribute{"readers"} }

func Writers() CollectionAttribute { return &collectionAttributeAttribute{"writers"} }

func AND(l Statement, r Statement) Statement { return Query(l).And(r) }
func OR(l Statement, r Statement) Statement  { return Query(l).Or(r) }
