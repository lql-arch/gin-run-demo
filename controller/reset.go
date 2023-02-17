package controller

func LoginReset() {
	if !ChatEmpty() { // 进行判断是否重置tempChat
		ChatReset()
	}
}
