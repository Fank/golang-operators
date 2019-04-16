package xmlx

import (
	"encoding/xml"
	"io"
)

func RootTag(r io.Reader) (string, error) {
	decoder := xml.NewDecoder(r)
	decoder.CharsetReader = func(label string, input io.Reader) (io.Reader, error) {
		return input, nil
	}

	for {
		t, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				return "", nil
			}
			return "", err
		}
		if se, ok := t.(xml.StartElement); ok {
			return se.Name.Local, nil
		}
	}
}
