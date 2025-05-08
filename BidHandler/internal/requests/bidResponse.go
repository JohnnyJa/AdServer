package requests

type BidResponse struct {
	ID      string    `json:"id"`              // ID відповіді (такий самий, як у BidRequest)
	SeatBid []SeatBid `json:"seatbid"`         // Массив seat'ів з відповідними ставками
	BidID   string    `json:"bidid,omitempty"` // Унікальний ID ставки DSP
	Cur     string    `json:"cur,omitempty"`   // Валюта (наприклад "USD")
}

// Один seat (DSP або логічна група ставок)
type SeatBid struct {
	Bid   []Bid  `json:"bid"`             // Массив ставок
	Seat  string `json:"seat,omitempty"`  // Ідентифікатор seat'а
	Group int    `json:"group,omitempty"` // 0 = не згруповано, 1 = згруповано
}

// Ставка на конкретну імпресію
type Bid struct {
	ID     string  `json:"id"`               // ID цієї ставки
	ImpID  string  `json:"impid"`            // ID відповідної імпресії
	Price  float32 `json:"price"`            // Ставка в поточній валюті
	AdID   string  `json:"adid,omitempty"`   // Ідентифікатор креативу
	NURL   string  `json:"nurl,omitempty"`   // Win notice URL (трекає виграш)
	Adm    string  `json:"adm,omitempty"`    // Адмова — креатив (HTML/JS/VAST)
	CrID   string  `json:"crid,omitempty"`   // Creative ID
	W      int     `json:"w,omitempty"`      // Ширина креативу
	H      int     `json:"h,omitempty"`      // Висота креативу
	DealID string  `json:"dealid,omitempty"` // Deal ID, якщо private deal
	Attr   []int   `json:"attr,omitempty"`   // Атрибути (наприклад, 1 = аудіо, 6 = autoplay)
}
