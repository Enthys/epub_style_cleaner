package pkg

import (
	"bytes"
	"log"
	"regexp"
)

var (
	cssSelectorRegex 	*regexp.Regexp;
	cssCommentRegex 	*regexp.Regexp;
	spaceRegex 			*regexp.Regexp;
)

func init() {
	cssSelectorRegex = regexp.MustCompile(`(.+?{).+?}`)
	cssCommentRegex = regexp.MustCompile(`\/\*.+?\*\/`)
	spaceRegex = regexp.MustCompile(" +")
}

type Selector string

type CssReader struct {
	content []byte
}

func NewCssReader(content []byte) *CssReader {
	return &CssReader{
		content: content,
	}
}

func (c *CssReader) GetCssSelectors() []Selector {
	content := c.removeNewLines(
		c.content
	);
	content = c.removeCssComments(content)
	content = c.addNewLineAfterSelectorRules(content)
	content = c.removeRepeatedSpaces(content)

	log.Printf("%+v", string(content))
	
	return []Selector{}
}

func (c CssReader) addNewLineAfterSelectorRules(content []byte) []byte {
	return bytes.ReplaceAll(content, []byte("}"), []byte("}\n"))
}

func (c CssReader) removeCssComments(content []byte) []byte {
	return cssCommentRegex.ReplaceAll(content, []byte(""))
}

func (c CssReader) removeNewLines(content []byte) []byte {
	return bytes.ReplaceAll(c.content, []byte("\n"), []byte(""))
}

func (c CssReader) removeRepeatedSpaces(content []byte) []byte {
	return spaceRegex.ReplaceAll(content, []byte(" "))
}

func (c CssReader) removeTabs(content []byte) []byte {
	return bytes.ReplaceAll(content, []byte("\t"), []byte(""))
}
