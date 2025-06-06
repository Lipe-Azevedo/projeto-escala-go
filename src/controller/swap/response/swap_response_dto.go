package response // O nome do pacote permanece 'response' dentro deste subdiretório

// SwapResponse é o DTO para retornar informações de uma troca.
type SwapResponse struct {
	ID            string  `json:"id"`
	RequesterID   string  `json:"requester_id"`
	RequestedID   string  `json:"requested_id"`
	CurrentShift  string  `json:"current_shift"`
	NewShift      string  `json:"new_shift"`
	CurrentDayOff string  `json:"current_day_off"`
	NewDayOff     string  `json:"new_day_off"`
	Status        string  `json:"status"`
	Reason        string  `json:"reason,omitempty"` // Adicionado omitempty se a razão puder ser vazia
	CreatedAt     string  `json:"created_at"`
	ApprovedAt    *string `json:"approved_at,omitempty"`
	ApprovedBy    *string `json:"approved_by,omitempty"`
}
