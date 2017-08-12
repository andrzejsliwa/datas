//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Item=BUILTINS"
package datas

import "github.com/cheekybits/genny/generic"

type Item generic.Type
