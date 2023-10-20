package fileinput

import "github.com/cdvelop/model"

func (f *File) Input() *model.Input {

	return &model.Input{
		InputName: f.Name(),
		Tag:       f,
		Validate:  f,
		TestData:  f,
	}
}

func (File) Name() string {
	return "File"
}

func (File) HtmlName() string {
	return "file"
}

func (File) ValidateField(data_in string, skip_validation bool, options ...string) error {
	return nil
}

func (File) GoodTestData() (out []string) {

	return
}

func (File) WrongTestData() (out []string) {
	return
}
