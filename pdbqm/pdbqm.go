package main

import (
	"fmt"
	"strings"

	chem "github.com/rmera/gochem"
	"github.com/rmera/gochem/qm"
	v3 "github.com/rmera/gochem/v3"
	"github.com/rmera/scu"
)

func main() {
	//to actually use it just fill pdbname, charge, otherres and the list of subsequences from os.Args
	pdbname := "../FILES/test.pdb"
	charge := 0
	otherres := []int{111}
	subseqs := make([][]int, 0, 5)
	chains := make([]string, 0, 5)
	for _, v := range []string{"5-8", "16-20"} {
		ss := strings.Split(v, "-")
		subse := make([]int, 0, 10)
		ini := scu.MustAtoi(ss[0])
		end := scu.MustAtoi(ss[1])
		for i := ini; i <= end; i++ {
			subse = append(subse, i)
		}
		chains = append(chains, "A") //This is pretty bad, only works for the example
		subseqs = append(subseqs, subse)
	}
	mol, err := chem.PDBFileRead(pdbname)
	scu.QErr(err)
	fmt.Println(mol.Len()) ///////////////////////////////////////
	//bkats:=make([]*chem.Atom,mol.Len())
	//bktop:=chem.NewTopology(charge,1,bkats)
	//bktop.CopyAtoms(mol)
	list, err := chem.CutBackRef(mol, chains, subseqs)
	fmt.Println(mol.Len()) /////////////////////////wAnt to see if CutBackRef removed atoms as it says in its comments.
	scu.QErr(err)
	list2, err := chem.CutBackRef(mol, []string{"A"}, [][]int{otherres})
	scu.QErr(err)
	list = append(list, list2...)
	qmmol := chem.NewTopology(charge, 1, nil)
	qmmol.SomeAtoms(mol, list)
	qmcoords := v3.Zeros(len(list))
	qmcoords.SomeVecs(mol.Coords[0], list)
	xtb := qm.NewXTBHandle()
	Q := new(qm.Calc)
	Q.Job = &qm.Job{Opti: true}
	Q.Method = "gfn2"
	fixed := make([]int, 0, 2*len(subseqs))
	for i := 0; i < qmmol.Len(); i++ {
		at := qmmol.Atom(i)
		if at.Name == "CTZ" || at.Name == "NTZ" {
			fixed = append(fixed, i)
		}
	}
	fmt.Println(fixed) ///////////
	Q.CConstraints = fixed
	xtb.BuildInput(qmcoords, qmmol, Q)
	chem.PDBFileWrite("4QM.pdb", qmcoords, qmmol, nil)

}
