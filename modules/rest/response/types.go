package response

type Response struct {
	Err     *Error      `json:"error"`
	Success *Successful `json:"success"`
}

type Successful struct {
	Data any `json:"data"`
}

type Error struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
