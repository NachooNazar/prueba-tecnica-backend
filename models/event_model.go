package models

// Event events
type Event struct {
	Id               string   `json:"id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"shortdescription" `
	LargeDescription string   `json:"largedescription"`
	Organizer        string   `json:"organizer"`
	Date             string   `json:"date"`
	Hour             string   `json:"hour"`
	Place            string   `json:"place"`
	State            bool     `json:"state"`
	Inscriptos       []string `json:"inscriptos"`
}

type Events []Event
