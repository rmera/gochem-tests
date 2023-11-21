package main

import (
	chem "github.com/rmera/gochem"
)

func main() {
	mol, err := chem.PDBFileRead("../FILES/renumber_test.pdb", true)
	if err != nil {
		panic(err.Error())
	}

	coord := mol.Coords[0]
	//	mol.FillIndexes()
	//	mol.ResetIDs()
	oldmolid := -1
	currmolid := 1
	for i := 0; i < mol.Len(); i++ {
		at := mol.Atom(i)
		if i == 0 {
			oldmolid = at.MolID
			currmolid = 1
		}
		if at.MolID != oldmolid {
			currmolid++
			oldmolid = at.MolID
		}
		at.MolID = currmolid
		at.ID = i + 1

	}

	newname := "../FILES/results/renumber_test_fixed.pdb" //strings.Replace(os.Args[1], ".pdb", "_fixed.pdb", 1)
	chem.PDBFileWrite(newname, coord, mol, nil)

}
