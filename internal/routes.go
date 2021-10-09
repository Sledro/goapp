package internal

func (s *Server) routes() {
	s.router.HandleFunc("/accounts/{publickey}", accountViewHandler).Methods("GET")
	r.HandleFunc("/accounts/{publickey}", accountEditHandler).Methods("POST")

	r.HandleFunc("/transactions/sign", transactionSignHandler).Methods("POST")
	r.HandleFunc("/transactions/send", transactionSendForwardHandler).Methods("POST")
	r.HandleFunc("/transactions/delete", transactionDeleteHandler).Methods("POST")
	r.HandleFunc("/transactions", transactionSendHandler).Methods("POST")
	r.HandleFunc("/transactions/{hash}", transactionEditHandler).Methods("POST")
	r.HandleFunc("/transactions/{hash}", transactionViewHandler).Methods("GET")
	r.HandleFunc("/transactions/account/{publickey}", transactionListAllWithAccount).Methods("GET")
	r.HandleFunc("/transactions/block/{blockid}", transactionListAllWithBlock).Methods("GET")

	r.HandleFunc("/blocks/last", getLastBlockHandler).Methods("GET")
}
