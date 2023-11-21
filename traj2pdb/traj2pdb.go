package main

import (
	"flag"
	"fmt"

	"os"

	chem "github.com/rmera/gochem"
	"github.com/rmera/gochem/traj/dcd"
	v3 "github.com/rmera/gochem/v3"
	//	"github.com/rmera/gochem/xtc"
	//	"github.com/rmera/scu"
	//	"gonum.org/v1/gonum/mat"
	//	"math"
	//	"sort"
	//	"strconv"
)

func main() {
	//The skip options
	skip := flag.Int("skip", 10, "How many frames to skip between reads.")
	begin := flag.Int("begin", 1, "The frame from where to start reading.")
	//	format := flag.Int("format", 0, "0 for OldAmber (crd, default), 2 for dcd (NAMD)")
	outformat := flag.String("outformat", "pdb", "dcd xyz or pdb")
	end := flag.Int("end", 100000, "The last frame")
	savelast := flag.Bool("savelast", false, "Save a pdb file with the last frame")
	flag.Parse()
	fmt.Println("program [-skip=number -begin=number2] pdbfile trajname outname")
	//	println("SKIP", *skip, *begin, args) ///////////////////////////
	mol, err := chem.PDBFileRead("../FILES/test.pdb", false)
	if err != nil {
		panic(err.Error())
	}
	var traj chem.Traj
	traj, err = dcd.New("../FILES/test.dcd")
	if err != nil {
		panic(err.Error())
	}
	Coords := make([]*v3.Matrix, 0, 0)
	var coords *v3.Matrix
	lastread := -1
	var trajW *dcd.DCDWObj
	if *outformat == "dcd" {
		trajW, err = dcd.NewWriter("../FILES/results/testOut.dcd", mol.Len())
		if err != nil {
			panic(err.Error())
		}
	}
	prevcoords := v3.Zeros(traj.Len())
	for i := 0; i < *end; i++ { //infinite loop, we only break out of it by using "break"  //modified for profiling
		if lastread < 0 || (i >= lastread+(*skip) && i >= (*begin)-1) {
			coords = v3.Zeros(traj.Len())
		}
		err := traj.Next(coords) //Obtain the next frame of the trajectory.
		if err != nil {
			_, ok := err.(chem.LastFrameError)
			if ok {
				if *savelast {
					chem.PDBFileWrite("lastframe.pdb", prevcoords, mol, nil)
				}
				break //We processed all frames and are ready, not a real error.

			} else {
				panic(err.Error())
			}
		}
		if (lastread >= 0 && i < lastread+(*skip)) || i < (*begin)-1 { //not so nice check for this twice
			continue
		}
		lastread = i
		if *outformat == "dcd" {
			err = trajW.WNext(coords)
			if err != nil {
				panic(err.Error())
			}
		} else {
			Coords = append(Coords, coords)
		}
		if *savelast {
			prevcoords.Copy(coords)
		}
		coords = nil // Not sure this works
	}
	var fout *os.File
	if *outformat != "dcd" {
		fout, err = os.Create("../FILES/results/test2.pdb")
		if err != nil {
			panic(err.Error())
		}
		defer fout.Close()
	}
	err = chem.MultiPDBWrite(fout, Coords, mol, nil)
	if err != nil {
		panic("Couldn't write output file: " + err.Error())
	}
}
