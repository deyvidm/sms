package types

const StatusFailed = "failed"
const StatusSuccess = "success"

type LoginData struct {
	Username string `json:"username" binding:"required,alphanum,min=3,max=255"` // 3 is a holy number
	Password string `json:"password" binding:"required,alphanum,min=6,max=255"` // min 6 for brcypt hash
}

type NewContactData struct {
	FirstName string `json:"first_name" binding:"required,alpha,min=3,max=50"`
	LastName  string `json:"last_name" binding:"required,alpha,min=3,max=50"`
	Phone     string `json:"phone" binding:"required,e164"` // e164 is the standard +11234567890
}

type NewEventData struct {
	Ttile string `json:"title" binding:"required,max=255"`
}
