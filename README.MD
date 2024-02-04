# Converter
Conversor de magnitudes de longitud (m to ft) (ft to m)

# Installation
go get -v github.com/Dnnis26/converter

# Usage
 package main

 import (

	 "fmt"
	 "github.com/Dnnis26/converter"

 )

 func main( ) {

	 m := converter.MtoF(14)
	 fmt.Println(m)

	 f := converter.FtoM(44)
	 fmt.Println(f)

 }