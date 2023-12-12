package fileinput

import (
	"github.com/cdvelop/model"
)

func (f FileInput) Input(default_enable bool) *model.Input {

	return &model.Input{
		InputName:       "file",
		Tag:             f,
		ItemViewAdapter: f,
		ResetParameters: &model.ResetParameters{
			CallJsFunWithParameters: model.CallJsFunWithParameters{
				FuncNameCall: "resetInputFile",
				Enable:       default_enable,
				AddParams:    map[string]any{},
			},
		},
		Validate: nil,
		TestData: nil,
	}
}

func (FileInput) HtmlName() string {
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
