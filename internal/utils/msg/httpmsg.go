package msg

const (
	LoginSuccess        string = "login success"
	RegistrationSuccess string = "registration success"
	DeleteSuccess       string = "delete success"
	OrderCreateSuccess  string = "order create success"
	UpdateSuccess       string = "update success"
	UserIsAuthorized    string = "user is authorized"
)

type HTTPMessage map[string]string

func New(message string) HTTPMessage {
	return HTTPMessage{"message": message}
}

func (e HTTPMessage) Err() map[string]string {
	return HTTPMessage{"error": e["message"]}
}
