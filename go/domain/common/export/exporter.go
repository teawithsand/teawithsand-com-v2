package export

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/teawithsand/webpage/domain/common/export/remtrans"
	"github.com/teawithsand/webpage/domain/common/export/remtypes"
	"github.com/teawithsand/webpage/domain/common/export/remval"
	"github.com/teawithsand/webpage/util/typescript"
)

type Exporter struct {
	Dir string
}

const modeFile = 0660
const modeDir = 0770

func (ex *Exporter) mustExportValidations(to string, val interface{}) {
	res, err := json.MarshalIndent(val, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path.Join(ex.Dir, to), res, modeFile)
	if err != nil {
		panic(err)
	}
}

func (ex *Exporter) mustExportTranslations(to string, val interface{}) {
	res, err := json.MarshalIndent(val, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path.Join(ex.Dir, to), res, modeFile)
	if err != nil {
		panic(err)
	}
}

func (ex *Exporter) mustExportTypescript(to string, ts *typescript.Converter) {
	res, err := ts.Convert(nil)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path.Join(ex.Dir, to), []byte(res), modeFile)
	if err != nil {
		panic(err)
	}
}

func (ex *Exporter) MustClear() {
	if len(ex.Dir) < 5 {
		panic(fmt.Errorf("unsafe remove: %s", ex.Dir))
	}
	os.RemoveAll(ex.Dir)
}

func (ex *Exporter) mustMkdir(dir string) {
	err := os.MkdirAll(path.Join(ex.Dir, dir), modeDir)
	if err != nil {
		panic(err)
	}

}

func (ex *Exporter) MustExportValidations() {
	ex.mustMkdir("livr")

	ex.mustMkdir("livr/user")
	ex.mustExportValidations("livr/user/register.json", remval.UserRegistrationValidationRules)
	ex.mustExportValidations("livr/user/confirm_registration.json", remval.UserConfirmRegistrationRules)
	ex.mustExportValidations("livr/user/change_email.json", remval.UserChangeEmailValidationRules)
	ex.mustExportValidations("livr/user/change_password.json", remval.UserChangePasswordValidationRules)

	ex.mustMkdir("livr/langka")
	ex.mustExportValidations("livr/langka/word_set_create.json", remval.LangkaWordSetCreateDataValidationRules)
}

func (ex *Exporter) MustExportTypescript() {
	ex.mustMkdir("typings")
	ex.mustExportTypescript("typings/api.ts", remtypes.GetConverter())
}

func (ex *Exporter) MustExportTranslations() {
	ex.mustMkdir("trans")
	for k, v := range remtrans.Translations {
		ex.mustExportTranslations(fmt.Sprintf("trans/%s.json", k), v)
	}
}

// TODO(teawithsand): function for exporting not registered translation keys
