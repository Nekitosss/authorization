package structures


type RegisterVerificationType int

const (
	RegisterTypeEmail RegisterVerificationType = iota
	RegisterTypeTelegram
)


type RegisterInfo struct {
	Login string `json:"login"`
	
	Email string `json:"email"`
	
	Password string `json:"password"`
	
	NameAlias string `json:"nameAlias"`
	
	RegisterType RegisterVerificationType `json:"registerType"`
}