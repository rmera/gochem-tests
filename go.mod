module github.com/rmera/gochem_tests

go 1.22

//replace github.com/rmera/gochem => github.com/rmera/gochem v0.6.4-0.20231011225313-3a879af9b24f

//replace github.com/rmera/gochem => /wrk/programs/github.com/rmera/gochem

require (
	github.com/rmera/gochem v0.7.1
	github.com/rmera/scu v0.2.0
)

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/skelterjohn/go.matrix v0.0.0-20130517144113-daa59528eefd // indirect
	gonum.org/v1/gonum v0.7.0 // indirect
)
