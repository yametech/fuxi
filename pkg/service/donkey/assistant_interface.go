package donkey

type IAssistant interface {
	Run() error
	Stop()
}
