package arxiv

import "time"

type Entry struct {
	ID        string    `xml:"id"`
	Updated   time.Time `xml:"updated"`
	Published time.Time `xml:"published"`
	Title     string    `xml:"title"`
	Summary   string    `xml:"summary"`
}

type Feed struct {
	Entries []Entry `xml:"entry"`
}
