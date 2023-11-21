package main

import (
	"fmt"

	chem "github.com/rmera/gochem"
)

func main() {
	mol, err := chem.PDBFileRead("../FILES/test.pdb", true)
	if err != nil {
		panic(err.Error())
	}
	coord := mol.Coords[0]
	c, top, err := chem.BackboneCGize(coord, mol, true)
	fmt.Println(err)
	fmt.Println(c, c.NVecs(), top.Len())
	chem.PDBFileWrite("../FILES/results/CGized.pdb", c, top, nil)

}
