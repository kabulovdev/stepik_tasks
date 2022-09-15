package main

import "fmt"

type futballteam interface {
	name()
	position()
	bandscore()
	playtime()
	goals()
}

func plname(p futballteam) {
	p.name()
}
func playerpos(p futballteam) {
	p.position()
}
func bandscorepl(p futballteam) {
	p.bandscore()
}
func plplaytime(p futballteam) {
	p.playtime()
}
func plGoals(p futballteam) {
	p.goals()
}

type hujumchi struct {
	plname    string
	pazitsiya string
	pltime    int
	goalspl   int
}
type himoychi struct {
	plname    string
	pazitsiya string
	pltime    int
	goalspl   int
}

func (d *hujumchi) name() {
	fmt.Println(d.plname)
}
func (d *hujumchi) position() {
	fmt.Println(d.pazitsiya)
}
func (d *hujumchi) playtime() {
	fmt.Println(d.pltime)
}
func (d *hujumchi) goals() {
	fmt.Println(d.goalspl)
}
func (d *hujumchi) bandscore() {
	if d.pltime > 10 && d.goalspl > 10 {
		fmt.Println("8.50 - 10.00")
	} else if d.goalspl < 10 && d.goalspl > 5 {
		fmt.Println("6.00-8.50")
	} else {
		fmt.Println("1.00 - 6.00")
	}
}

func main() {
	neymar := &hujumchi{
		plname:    "neymar jr",
		pazitsiya: "hujumji",
		pltime:    20,
		goalspl:   12,
	}

	neymar.goals()
	neymar.bandscore()
	neymar.position()
	neymar.playtime()
	plGoals(neymar)
	plname(neymar)

}
