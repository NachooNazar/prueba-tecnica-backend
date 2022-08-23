package models

// Event events
type Event struct {
	Id               string `json:"uid"`
	Title            string `json:"title"`
	ShortDescription string `json:"shortdescription" `
	LargeDescription string `json:"largedescription"`
	Organizer        string `json:"organizer"`
	Date             string `json:"date"`
	Place            string `json:"place"`
	State            bool   `json:"state"`
}

type Events []Event
