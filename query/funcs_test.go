package query

import "testing"
import . "github.com/onsi/gomega"

func Test(t *testing.T) {
	RegisterTestingT(t)
	Expect("").To(BeEquivalentTo(""))
}

func TestRawStatement(t *testing.T) {
	RegisterTestingT(t)
	Expect(Stringize(Raw("name = \"ciao\""))).To(BeEquivalentTo("name = \"ciao\""))
}

func TestNameEqStatement(t *testing.T) {
	RegisterTestingT(t)
	Expect(Stringize(Name().Equal("ciao"))).To(BeEquivalentTo("name = \"ciao\""))
}
func TestNameContainStatement(t *testing.T) {
	RegisterTestingT(t)
	Expect(Stringize(Name().Contains("ciao"))).To(BeEquivalentTo("name contains \"ciao\""))
}
func TestNameNotEqStatement(t *testing.T) {
	RegisterTestingT(t)
	Expect(Stringize(Name().NotEqual("ciao"))).To(BeEquivalentTo("name != \"ciao\""))
}

func TestPropertiesHasStatement(t *testing.T) {
	RegisterTestingT(t)
	Expect(Stringize(PropertiesHas("foo", "ciao"))).To(BeEquivalentTo(`properties has { key="foo" and value="ciao" }`))
}

func TestFullTextStatement(t *testing.T) {
	RegisterTestingT(t)
	Expect(Stringize(FullText().Contains("ciao"))).To(BeEquivalentTo(`fullText contains "ciao"`))
}

func TestMimeTypeEqStatement(t *testing.T) {
	RegisterTestingT(t)
	Expect(Stringize(MimeType().Equal("ciao"))).To(BeEquivalentTo("mimeType = \"ciao\""))
}
func TestMimeTypeNotEqStatement(t *testing.T) {
	RegisterTestingT(t)
	Expect(Stringize(MimeType().NotEqual("ciao"))).To(BeEquivalentTo("mimeType != \"ciao\""))
}

func TestAndQuery(t *testing.T) {
	RegisterTestingT(t)
	var stm1 = MimeType().Equal("text/plain")
	var stm2 = Name().Contains("ciao")
	var query = Query(stm1).And(stm2)
	var expected = `(mimeType = "text/plain") and (name contains "ciao")`
	Expect(Stringize(query)).To(BeEquivalentTo(expected))
}
func TestAndOrQuery1(t *testing.T) {
	RegisterTestingT(t)
	var stm1 = MimeType().Equal("text/plain")
	var stm2 = Name().Contains("ciao")
	var query = Query(stm1).And(stm2).Or(Raw(`name = "pluto"`))
	var expected = `((mimeType = "text/plain") and (name contains "ciao")) or (name = "pluto")`
	Expect(Stringize(query)).To(BeEquivalentTo(expected))
}

func TestAndOrQuery2(t *testing.T) {
	RegisterTestingT(t)
	var stm1 = MimeType().Equal("text/plain")
	var stm2 = Name().Contains("ciao")
	var query = OR(AND(stm1, stm2), Raw(`name = "pluto"`))
	var expected = `((mimeType = "text/plain") and (name contains "ciao")) or (name = "pluto")`
	Expect(Stringize(query)).To(BeEquivalentTo(expected))
}
