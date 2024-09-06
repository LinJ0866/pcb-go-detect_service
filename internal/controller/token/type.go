package token

type tokenResDataStruct struct {
	Token string `json:"token"`
}

type tokenResStruct struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data tokenResDataStruct `json:"data"`
}
