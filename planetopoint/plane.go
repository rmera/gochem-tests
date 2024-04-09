package main

import (
	"fmt"
	"os"

	chem "github.com/rmera/gochem"
	v3 "github.com/rmera/gochem/v3"
	"github.com/rmera/scu"
)

//This program will align the best plane passing through a set of atoms in a molecule with the XY-plane.
//Usage:
func main() {
	mol, err := chem.XYZFileRead("../FILES/test_plane.xyz")
	if err != nil {
		panic(err.Error())
	}
	var indexes []int
	indexes, err = scu.IndexFileParse("../FILES/indexes") //we take the indexes from a file
	if err != nil {
		panic(err.Error())
	}
	//	}
	some := v3.Zeros(len(indexes)) //will contain the atoms selected to define the plane.
	some.SomeVecs(mol.Coords[0], indexes)
	//for most rotation things it is good to have the molecule centered on its mean.
	mol.Coords[0], _, _ = chem.MassCenter(mol.Coords[0], some, nil)
	//As we changed the atomic positions, must extract the plane-defining atoms again.
	some.SomeVecs(mol.Coords[0], indexes)
	//The strategy is: Take the normal to the plane of the molecule (os molecular subset), and rotate it until it matches the Z-axis
	//This will mean that the plane of the molecule will now match the XY-plane.
	best, err := chem.BestPlane(some, nil)
	if err != nil {
		panic(err.Error())
	}
	z, _ := v3.NewMatrix([]float64{0, 0, 1})
	zero, _ := v3.NewMatrix([]float64{0, 0, 0})
	fmt.Fprintln(os.Stderr, "Best  Plane", best, z, indexes)
	axis := v3.Zeros(1)
	axis.Cross(best, z)
	fmt.Fprintln(os.Stderr, "axis", axis)
	//The main part of the program, where the rotation actually happens. Note that we rotate the whole
	//molecule, not just the planar subset, this is only used to calculate the rotation angle.
	mol.Coords[0], err = chem.RotateAbout(mol.Coords[0], zero, axis, chem.Angle(best, z))
	if err != nil {
		panic(err.Error())
	}
	//Now we write the rotated result.
	final, err := chem.XYZStringWrite(mol.Coords[0], mol)
	fmt.Print(final)
	fmt.Fprintln(os.Stderr, err)
}
