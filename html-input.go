package fileinput

import (
	"github.com/cdvelop/model"
)

func (f *File) Input() *model.Input {

	return &model.Input{
		InputName:  "file",
		Tag:        f,
		InputView:  f,
		InputReset: f,
		Validate:   nil,
		TestData:   nil,
	}
}

func (File) HtmlName() string {
	return "file"
}

// func (f File) ValidateField(data_in string, skip_validation bool, options ...string) error {
// 	return f.input_id.ValidateField(data_in, skip_validation)
// }

// func (File) GoodTestData() (out []string) {

// 	return
// }

// func (File) WrongTestData() (out []string) {
// 	return
// }
