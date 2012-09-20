package main
import (
  "fmt"
  "math/rand"
  //"math"
  "time"
) 


type Juego struct {
  fila1Objetivo [4]int
  fila2Objetivo [8]int
  
  fila1 [4]int
  fila2 [8]int
  
  tipoHeuristica int
  
  heuristica int
}

func (this *Juego)Init(){
  for t:=1;t<9;t++{
    this.fila2Objetivo[t-1]=t;
  }     
  l:=0
  for t:=9;t<12;t++{
    this.fila1Objetivo[l]=t;
    l++;
  }    
  this.fila1Objetivo[l]=-1
  rand.Seed(time.Now().UTC().UnixNano())
  vector:=make([]int,12,12)    
  for i:=1;i<12;i++{
    vector[i-1]=i
  }
  vector[11]=-1
  i:=0 
  j:=0
  for v:=0; v<12;v++{
    rn:=rand.Intn(len(vector))
    if(i<8){
      this.fila2[i]=vector[rn]
      i++      
    }else{
      this.fila1[j]=vector[rn]
      j++
    }
    vector=append(vector[:rn],vector[rn+1:]...)
  
  }
  fmt.Println("Circulo 1",this.fila1)
  fmt.Println("Circulo 2",this.fila2)
  
}

func (this *Juego)Pos(Valor int,objetivo bool)(x,y int){
  var c1 [4]int
  var c2 [8]int
  if(objetivo){
    c1=this.fila1Objetivo
    c2=this.fila2Objetivo
  }else{
    c1=this.fila1
    c2=this.fila2
  }
  r:=0; p:=0;
  //fmt.Println("->c1",c1)
  //fmt.Println("->c2",c2)
  for i:=0;i<len(c1);i++{
    if(c1[i]==Valor){ 
      r=1; p=i      
    }
  }
  if(r==0){
    for i:=0;i<len(c2);i++{
    if(c2[i]==Valor){ 
      r=2; p=i      
    }
  }
  }  
  return r,p;
}
  

func (this *Juego) HeuristicaJuego()int{
  this.heuristica=0;
  for i:=0;i<len(this.fila2Objetivo);i++ {
    this.heuristica+=this.HeuristicaPieza(this.fila2Objetivo[i])
  }
  
  for i:=0;i<len(this.fila1Objetivo);i++{
    this.heuristica+=this.HeuristicaPieza(this.fila1Objetivo[i])
  } 
  
  return this.heuristica;
}


func (this *Juego) HeuristicaPieza(pieza int) int{
  fx:=this.Heuristicas()  
  return fx(pieza);
}

func max(x,y int)int {
  if(x>=y){
    return x
  }
  return y
}

func min(x,y int)int {
  if(x<y){
    return x
  }
  return y
}

func abs(x int)int{
  if(x<0){
    return -x
  }
  return x  
}



func (this *Juego) Heuristicas()func(Pieza int)int {
  if(this.tipoHeuristica==1){
    return func (Pieza int)int { 
		ro,po:=this.Pos(Pieza,true); r,p:=this.Pos(Pieza,false); 		
		var d int
		if(ro==r){
		  Lm:=r*4-1;
		  pm:=max(po,p)
		  pmi:=min(po,p)
		  d=min(abs(p-po),abs(Lm-pm)+abs(pmi)+1)		  
		}else{
		  if(ro>r){
		    Lm:=r*4-1
		    pt:=po/2
		    
		    //pt:=-(po/2)*(po%2-1)+(po/2+1)*(po%2)
		    
		    pm:=max(p,pt)
		    pmi:=min(p,pt)
		    
		    d=min(abs(p-pt),abs(Lm-pm)+abs(pmi)+1)+1
		    
		  }else{
		    Lm:=ro*4-1		    
		    pt:=p/2
		    //pt:=-(p/2)*(p%2-1)+(p/2+1)*(p%2)
		    
		    pm:=max(po,pt)
		    pmi:=min(po,pt)
		    
		    d=min(abs(po-pt),abs(Lm-pm)+abs(pmi)+1)+1		    
		  }		  
		}
		//fmt.Println(Pieza,d,"->>>",ro,po,"-->>",r,p)
		//fmt.Println(Pieza,d)
		return d;      
		}
  }
    return func (Pieza int)int {  
		ro,po:=this.Pos(Pieza,true); r,p:=this.Pos(Pieza,false); 
		if(ro==r&&po==p){return 0}		
		return 1;      
		} 
}

func (this *Juego) posiblesEstados()[]Juego{
  res:=make([]Juego,1,1)  
  r,p:=this.Pos(-1,false)
  longitud:=r*4-1 
  var valores []int;
    if(r==1){
      valores=[]int{-1,+1,2*p,2*p+1}
    }else{
      valores=[]int{-1,+1,p/2}
    }    
      var newState Juego;      
      var pf,rf int;
      for i:=0;i<len(valores);i++{	
      newState=Juego{this.fila1Objetivo,this.fila2Objetivo,this.fila1,this.fila2,this.tipoHeuristica,this.heuristica};      
      
      if(i>=2){
	rf=-r+3
	pf=valores[i]
	//Intercambio	
	newState.Intercambio(r,p,rf,pf)
      }else{
	if((p+valores[i]<0)||(p+valores[i]>longitud)){
	  rf=r
	  pf=abs(p-longitud)
	  //Intercambio
	  newState.Intercambio(r,p,rf,pf)
	}else{
	 rf=r
	 pf=p+valores[i]	 
	 //Intercambio
	 newState.Intercambio(r,p,rf,pf)
	}
      }
      heut:=newState.heuristica-this.HeuristicaPieza(-1)-this.HeuristicaPieza(this.GetValor(rf,pf))+newState.HeuristicaPieza(-1)+newState.HeuristicaPieza(newState.GetValor(r,p))
      fmt.Println("H(t+1) ",heut)
      //fmt.Println("hola :)",newState.fila1,newState.fila2)
      fmt.Println("recalc ",newState.HeuristicaJuego())
      
      }
  fmt.Println("",r,p)
  return res
}

func (this *Juego) GetValor(r,p int)int{
  longitud:=4*r-1
  if(r==1){
    if(p>=0&&p<=longitud){
    return this.fila1[p]  
    }
  }else{
    if(p>=0&&p<=longitud){
    return this.fila2[p]
    }
  }
   return 0;
}

func (this *Juego) SetValor(r,p int,valor int){
  longitud:=4*r-1
  if(r==1){
    if(p>=0&&p<=longitud){
    this.fila1[p]=valor
    }
  }else{
    if(p>=0&&p<=longitud){
    this.fila2[p]=valor
    }
  }
}

func (this *Juego) Intercambio(r,p int, rf,pf int){
  valorTemp1:=this.GetValor(r,p)
  valorTemp2:=this.GetValor(rf,pf)
  this.SetValor(r,p,valorTemp2)
  this.SetValor(rf,pf,valorTemp1)
  
}


type Nodo struct {  
  Estado Juego
  Hijos []*Nodo
  Nivel int
}
func (this *Nodo) addChild(estado Juego)bool{
  if(this.Hijos!=nil){
    this.Hijos=make([]*Nodo,1,1)
    nodo:=&Nodo{estado,nil,this.Nivel+1}
    this.Hijos[0]=nodo
  }else{
    HijoExtra:=make([]*Nodo,1,1)
    nodo:=&Nodo{estado,nil,this.Nivel+1}
    HijoExtra[0]=nodo
    this.Hijos=append(this.Hijos,HijoExtra...)
  }
  return true
}

/*func AStar(estadoI Juego)(Raiz *Nodo){
   raiz:=&Nodo{estadoI,nil,0}
  
  return raiz
}*/



func main(){

  var prueba Juego
 
  prueba.Init()
  prueba.tipoHeuristica=1;
  
  fmt.Println(" ",prueba.HeuristicaJuego())
  prueba.posiblesEstados()
  
  
}

