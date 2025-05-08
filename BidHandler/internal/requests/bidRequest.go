package requests

import "github.com/google/uuid"

type BidRequest struct {
	ID     string  `json:"id"`               // Унікальний ID запиту
	Imp    []Imp   `json:"imp"`              // Массив імпресій
	Site   *Site   `json:"site,omitempty"`   // Інфо про сайт (або App)
	Device *Device `json:"device,omitempty"` // Інфо про пристрій
	User   *User   `json:"user,omitempty"`   // Інфо про користувача
	Test   int     `json:"test,omitempty"`   // 1 = тестовий запит
	TMax   int     `json:"tmax,omitempty"`   // Максимальний час очікування, мс
	AT     int     `json:"at,omitempty"`     // Auction Type: 1 = first-price, 2 = second-price
}

type Imp struct {
	ID       string    `json:"id"`                 // ID імпресії
	Banner   *Banner   `json:"banner,omitempty"`   // Якщо банер
	Video    *Video    `json:"video,omitempty"`    // Якщо відео
	Instl    int       `json:"instl,omitempty"`    // Interstitial
	TagID    uuid.UUID `json:"tagid,omitempty"`    // ID placement-а
	BidFloor float32   `json:"bidfloor,omitempty"` // Мінімальна ціна
	Ext      *ImpExt   `json:"ext,omitempty"`
}

// Розширення з кастомним таргетингом
type ImpExt struct {
	Targeting *ImpTargeting `json:"targeting,omitempty"`
}

type ImpTargeting struct {
	Keyword string `json:"keyword,omitempty"` // Наприклад: "sports", "finance", "gaming"
}

// Банерна імпресія
type Banner struct {
	W   int `json:"w,omitempty"`
	H   int `json:"h,omitempty"`
	Pos int `json:"pos,omitempty"` // 1 = above the fold, 3 = below
}

// Відео-імпресія
type Video struct {
	MIMEs       []string `json:"mimes"`       // Підтримувані формати (video/mp4 тощо)
	Minduration int      `json:"minduration"` // Мін. тривалість
	Maxduration int      `json:"maxduration"` // Макс. тривалість
	StartDelay  int      `json:"startdelay"`  // Коли відео починається
	Protocols   []int    `json:"protocols"`   // VAST-протоколи
	W           int      `json:"w,omitempty"`
	H           int      `json:"h,omitempty"`
}

// Інформація про сайт
type Site struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain,omitempty"`
	Page   string `json:"page,omitempty"`
}

// Інформація про пристрій
type Device struct {
	UA         string `json:"ua,omitempty"`
	IP         string `json:"ip,omitempty"`
	OS         string `json:"os,omitempty"`
	DeviceType int    `json:"devicetype,omitempty"`
	IFA        string `json:"ifa,omitempty"` // ID for advertisers
}

// Інформація про користувача
type User struct {
	ID       string `json:"id,omitempty"`
	BuyerUID string `json:"buyeruid,omitempty"`
	YOB      int    `json:"yob,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Geo      *Geo   `json:"geo,omitempty"`
}

// Геолокація
type Geo struct {
	Country string  `json:"country,omitempty"`
	Lat     float64 `json:"lat,omitempty"`
	Lon     float64 `json:"lon,omitempty"`
}
