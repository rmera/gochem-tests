package main

import (
	"fmt"

	chem "github.com/rmera/gochem"
	v3 "github.com/rmera/gochem/v3"
	"github.com/rmera/scu"
)

//Takes an XYZ file and the zero-based indexes of 4 atoms (a,b,c and d) It calculates the angle
// between a plane defined by the vectors ab and ac and the vector ad. The atoms b and c must be given
//in an order such that the cross produc ab x ac points in the same general direction as the vector ad. The
//indexes a,b,c and d have to be given after the XYZ file name as a quoted string separated by spaces.
//This is not really the improper angle!!

func main() {
	mol, err := chem.XYZFileRead("../FILES/sample.xyz")
	if err != nil {
		panic(err.Error())
	}
	//	ndx, _ := scu.IndexStringParse("0 1 3 4") //""1 2 3 6")
	ndx, _ := scu.IndexStringParse("21 19 23 73")
	fmt.Println(ndx)
	coord := v3.Zeros(mol.Len())
	coord.Copy(mol.Coords[0])
	piv := coord.VecView(ndx[0])
	v1 := coord.VecView(ndx[1])
	v2 := coord.VecView(ndx[2])

	v1.Sub(v1, piv)
	v2.Sub(v2, piv)
	plane := v3.Zeros(1)
	plane.Cross(v1, v2)
	vfin := coord.VecView(ndx[3])
	vfin.Sub(vfin, piv)
	r2d := chem.Rad2Deg
	angle := chem.Angle(vfin, plane) * r2d //this
	fmt.Println(90-angle, 90+angle)
	coord.Copy(mol.Coords[0])                   //reset
	ndx2, _ := scu.IndexStringParse("1 2 3 55") //0")
	piv = coord.VecView(ndx2[0])
	v1 = coord.VecView(ndx2[1])
	v2 = coord.VecView(ndx2[2])
	vfin = coord.VecView(ndx2[3])
	fmt.Println(chem.Improper(piv, v1, v2, vfin)*r2d, chem.ImproperAlt(piv, v2, v1, vfin)*r2d, chem.Dihedral(piv, v1, v2, vfin)*r2d, "Improper should be 21.3deg") //21.3
	v1.Sub(v1, piv)
	v2.Sub(v2, piv)
	fmt.Println(chem.Angle(v1, v2)*r2d, "Should be ~108.2deg")
	methyl := coord.VecView(0)
	methyl.Sub(methyl, piv)
	fmt.Println("Distance", methyl.Norm(2), "should be ~1.5A")
}
