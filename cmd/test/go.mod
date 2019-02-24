module test

// automatically added
// require:
// github.com/johndelavega/gomod v0.0.4
// github.com/johndelavega/gomod/mod2 v0.0.0-20190224164920-adba6b188599

// replace github.com/johndelavega/gomod/mod2 v0.0.0 => ./local
// ../../ points to local primary package gomod
replace github.com/johndelavega/gomod/mod2 v0.0.0 => ../../

// OK
//replace github.com/johndelavega/gomod/mod2 v0.0.0 => ../../mod2

require github.com/johndelavega/gomod/mod2 v0.0.0

// the module system will update to minimum version automatically
// changing to v0.0.6 or lower will revert back update automatically to v0.0.7
require github.com/johndelavega/gomod v0.0.11
