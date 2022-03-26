package main

import (
    
    "fmt"
  
    
)
//07:05:45PM
func main(){
    var h,m,s int
    var suffix string
    
    fmt.Scanf("%d:%d:%d%s",&h,&m,&s,&suffix)
   
    var hf24 int  
    var mf24 int
    var sf24 int
    
    
    if suffix == "AM" && h == 12 {
        hf24 = 0
        mf24 = m
        sf24 = s
        fmt.Printf("%02d:%02d:%02d",hf24,mf24,sf24)
    } else if suffix == "AM" {
        hf24 = h
        mf24 = m
        sf24 = s
         fmt.Printf("%02d:%02d:%02d",hf24,mf24,sf24)
    } else if suffix == "PM" && h != 12 {
        hf24 = h +12
        mf24 = m
        sf24 = s
         fmt.Printf("%02d:%02d:%02d",hf24,mf24,sf24)
    } else if suffix == "PM" && h == 12 {
        hf24 = h 
        mf24 = m
        sf24 = s
         fmt.Printf("%02d:%02d:%02d",hf24,mf24,sf24)
    }
    
    
  
    
    
}
