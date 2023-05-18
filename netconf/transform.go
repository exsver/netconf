package netconf

import (
	"bytes"
	"regexp"
)

func CorrectBackspaces(in []byte) (out []byte) {
	out = []byte{}
	temp := bytes.SplitAfter(in, []byte{0x8})

	for _, v := range temp {
		if len(v) == 0 { // empty slice in the end
			continue
		}

		if len(v) == 1 {
			if v[0] == 0x8 {
				if len(out) == 0 { // nothing to delete
					continue
				}

				out = out[0 : len(out)-1] // delete 1 byte
			} else {
				out = append(out, v...)
			}
		} else {
			if v[len(v)-1] == 0x8 {
				out = append(out, v[0:len(v)-2]...)
			} else {
				out = append(out, v...)
			}
		}
	}

	return out
}

func Normalize(in []byte) (out []byte) {
	reSpaces := regexp.MustCompile(`[\s]{1,}`)
	out = reSpaces.ReplaceAll(in, []byte(" "))
	out = bytes.ReplaceAll(out, []byte(" <"), []byte("<"))
	out = bytes.ReplaceAll(out, []byte("> "), []byte(">"))

	return out
}

func ConvertToPairedTags(in []byte) []byte {
	tags := bytes.SplitAfter(in, []byte(">"))

	for i, tag := range tags {
		if bytes.Contains(tag, []byte("/>")) {
			begin := bytes.Index(tag, []byte("<"))
			end := bytes.Index(tag, []byte("/"))
			openTag := append(tag[begin:end], []byte(">true")...)
			closeTag := append([]byte("</"), append(tag[begin+1:end], []byte(">")...)...)
			allTag := append(openTag, closeTag...)
			tags[i] = append(tag[:begin], allTag...)
		}
	}

	return bytes.Join(tags, []byte{})
}

func ConvertToSelfClosingTag(in []byte) []byte {
	re := regexp.MustCompile(`>true</.*?>`)
	out := re.ReplaceAll(in, []byte(`/>`))

	return out
}

// Convert:
//
//	level1/level2/level3
//
// to (without spaces):
//
//	<level1>
//	  <level2>
//	    <level3/>
//	  </level2>
//	</level1>
func ConvertToXML(in []byte) (out []byte) {
	ins := bytes.Split(in, []byte(`/`))
	out = append([]byte(`<`), append(ins[len(ins)-1], []byte(`/>`)...)...)

	for i := len(ins) - 2; i >= 0; i-- {
		out = append([]byte(`<`), append(ins[i], append([]byte(`>`), append(out, append([]byte(`</`), append(ins[i], []byte(`>`)...)...)...)...)...)...)
	}

	return out
}
