package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	chem "github.com/rmera/gochem"
	"github.com/rmera/gochem/align"
	"github.com/rmera/scu"
	// "github.com/rmera/gochem/traj/dcd"
)

func formatLOVO(chain string, data []*align.MolIDandChain) string {
	first := true
	molids := make([][2]int, 1, 5)
	molids[0] = [2]int{0, 0}
	for _, v := range data {
		if chain != v.Chain() {
			continue
		}
		mid := v.MolID()
		if first {
			molids[0][0] = mid
			molids[0][1] = mid
			first = false
			continue
		}
		a := len(molids) - 1
		if mid == molids[a][1]+1 {
			molids[a][1] = mid
		} else {
			molids = append(molids, [2]int{mid, mid})
		}
	}
	ret := ""
	for _, v := range molids {
		if v[0] == v[1] {
			ret = fmt.Sprintf("%s,%d", ret, v[0])
		} else {
			ret = fmt.Sprintf("%s,%d-%d", ret, v[0], v[1])

		}

	}

	return fmt.Sprintf("%s %s BB", strings.TrimLeft(ret, ","), chain)
}

func main() {
	pdbname := "../FILES/test.pdb"
	trjname := "../FILES/test.xtc"
	mol, err := chem.PDBFileRead(pdbname, true)
	scu.QErr(err)
	//	traj, err := xtc.New(trjname)
	//	scu.QErr(err)

	fmt.Printf("#LOVO References:\n# 10.1371/journal.pone.0119264\n# 10.1186/1471-2105-8-306\n")
	opt := align.DefaultOptions()
	name, chain := []string{"CA"}, []string{"A"}
	opt.AtomNames(name)
	opt.Chains(chain)
	//      fmt.Println(opt.AtomNames(), opt.Chains(), opt.Skip(), opt.LessThanRMSD()) /////////////////////
	fmt.Printf("# Starting LOVO calculation. You might as well go for a coffee.\n")
	lovoret, err := align.LOVO(mol, mol.Coords[0], trjname, opt)
	fmt.Printf("# LOVO calculation finished.\n")
	if err != nil {
		log.Fatal("Couldn't obtain LOVO indexes for the superposition: %s", err.Error())
	}
	sort.Sort(lovoret)
	//I'd like to have an option for VMD also, but don't know the selection 'language' there.
	fmt.Println("# LOVO atom indexes:", lovoret)
	fmt.Println("\n# PyMOL selection for LOVO-selected residues: ", lovoret.PyMOLSel())

	fmt.Println("\n# gmx make_ndx selection LOVO-selected residues:\n ")
	fmt.Println(lovoret.GMX())
	fmt.Println("# LOVO CA indexes in goMD selection format:")
	var errcheck error
	for i, v := range chain {
		text := formatLOVO(v, lovoret.Nmols)
		reftext := "2-9,14-22,28-40,42-55,57-76,78-89,92-95,98,100-127,140-152 A BB"
		fmt.Println(text)
		if text != reftext {
			errcheck = fmt.Errorf("Test failed for chain!!! %d %s\n Is %s\n Should be %s ", i, v, text, reftext)

		}
	}
	if errcheck != nil {
		log.Println(err.Error())
	} else {
		fmt.Println("It all looks correct from here!")
	}

}
