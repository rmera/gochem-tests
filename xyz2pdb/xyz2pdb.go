package main

import (
	"os"
	"strconv"
	"strings"

	chem "github.com/rmera/gochem"
)

func main() {
	xyz, err := chem.XYZFileRead("../FILES/ovi.xyz")
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < xyz.Len(); i++ {
		a := xyz.Atom(i)
		if len(os.Args) > 1 && os.Args[1] == "-res" {
			a.MolID = i
		} else {
			a.MolID = 1
		}
		a.ID = i
		a.Chain = "A"
		a.MolName = "UNK"
		a.Name = strings.ToUpper(a.Symbol) + strconv.Itoa(i)
	}
	name := "../FILES/results/obi.pdb" //strings.Replace(strings.ToLower("obi.xyz"), ".xyz", ".pdb", 1)
	chem.PDBFileWrite(name, xyz.Coords[0], xyz, nil)
}
