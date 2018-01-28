package nlp

//I18n used to define different languages
type I18n string

const (
	//ES spanish identifier
	ES I18n = "es"
	//EN englush identifier
	EN = "en"
)

//DataInput structure used as an input for any nlp technique
type DataInput struct {
	SrcCntent string
}

//DataOutput interface defined to be used as an output for Nlp techniques
type DataOutput interface{}

//DefNlp interface
type DefNlp interface {
	SetLanguage(language I18n)
	SetDataInput(dataInput *DataInput)
	Process() DataOutput
}

//BaseNlp structure for common data
type BaseNlp struct {
	language  I18n
	dataInput *DataInput
}

//SetLanguage implements interface funcion
func (b *BaseNlp) SetLanguage(language I18n) {
	b.language = language
}

//SetDataInput implements interface funcion
func (b *BaseNlp) SetDataInput(dataInput *DataInput) {
	b.dataInput = dataInput
}

//Run generic function to be executed when running a nlp technique
func Run(nlp DefNlp) DataOutput {
	return nlp.Process()
}
