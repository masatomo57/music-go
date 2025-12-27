package conf

type Chord struct {
	Notes []Note
}

var (
	ChordC  = Chord{Notes: []Note{C3, E3, G3}}
	ChordD7 = Chord{Notes: []Note{C3, D3, Fs3, A3}}
	ChordF  = Chord{Notes: []Note{F3, A3, C4}}
	ChordG  = Chord{Notes: []Note{B2, D3, G3}}
	ChordG7 = Chord{Notes: []Note{B2, F3, G3}}

	ChordD   = Chord{Notes: []Note{D3, A3, D4}}
	ChordEm  = Chord{Notes: []Note{E3, B3, E4}}
	ChordGD  = Chord{Notes: []Note{G3, D4, G4}}
	ChordGBD = Chord{Notes: []Note{G3, B3, D4}}
	ChordEGB = Chord{Notes: []Note{E3, G3, B3}}
	ChordG4  = Chord{Notes: []Note{G3, B3, D4, G4}}

	ChordNone = Chord{Notes: []Note{None}}
)
