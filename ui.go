package main
import( 
  //"os"
  "fmt"
  "strconv"
  "github.com/mattn/go-gtk/glib"
  "github.com/mattn/go-gtk/gtk"
  "github.com/mattn/go-gtk/gdk"
   "strings"
   //"io"
   "os"
  //"unsafe"
  //"path"
  //"github.com/mattn/go-gtk/gdkpixbuf"
)
func main(){
  var prueba int
  
  var Actual Juego
  
  Actual.Init()
  
  var  heuristica int
  
  heuristica=1
  
  var r1 [4]int
  var r2 [8]int
  
  Actual.fila1=Actual.fila1Objetivo
  Actual.fila2=Actual.fila2Objetivo
  
  r1=Actual.fila1Objetivo
  r2=Actual.fila2Objetivo
  
  //fmt.Println(r1,r2)
  
  prueba=100
  //var menuitem *gtk.GtkMenuItem
  gtk.Init(nil)
  window:=gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
  window.SetPosition(gtk.GTK_WIN_POS_CENTER);
  window.SetTitle(" A star");
  window.Connect("destroy",func(ctx *glib.CallbackContext){
      println("got destory!\n",ctx.Data().(string))
      gtk.MainQuit()
  },"foo")
  
  vbox:=gtk.HBox(false,14)
  
  //menubar:=gtk.MenuBar()
  //vbox.PackStart(menubar,false,false,0)
 
  vpaned:=gtk.VPaned()
  vbox.Add(vpaned)
  
  frame:=gtk.Frame("Profundiad")      
  framebox:=gtk.HBox(false,14)  

  framebox1:=gtk.VBox(true,0)
  frame1:=gtk.Frame("Rompecabezas ")      
  
  frame.Add(framebox)    
  frame1.Add(framebox1)
  
  drawingarea := gtk.DrawingArea()
  
 
 
  vpaned.Add(frame)
  vpaned.Add(frame1)
  
  
  button := gtk.ButtonWithLabel("Resolver")
  button2 := gtk.ButtonWithLabel("Desordenar")
  button3:=gtk.ButtonWithLabel(" Nivel de Desorden")
  
  //TxtPasos:=gtk.Entry()
  
  frameHu:=gtk.VBox(true,2)  
  
  lheuristica:=gtk.Label("Heuristica")
  comboBox:=gtk.ComboBoxEntryNewText()
  comboBox.AppendText("Piezas Desordenadas")
  comboBox.AppendText("Distancia entre Piezas")
  
  comboBox.Connect("changed",func(){
    if(strings.EqualFold(comboBox.GetActiveText(),"Piezas Desordenadas")){
      heuristica=2
    }else{
      heuristica=1
    }
  })
  
  
  frameHu.Add(lheuristica)
  frameHu.Add(comboBox)

  entry := gtk.Entry()
  
  framebox.Add(entry)
  
  framebox.Add(button)
  framebox.Add(frameHu)
  
  framebox.Add(button2)
  framebox.Add(button3)
  
   
   //framebox1.Add(TxtPasos)
  //var gdkwin *gdk.GdkWindow
  var pixmap *gdk.GdkPixmap
  var gc *gdk.GdkGC
  
  
  drawingarea.Connect("configure-event", func() {
		if pixmap != nil {
			pixmap.Unref()
		}
		var allocation gtk.GtkAllocation
		drawingarea.GetAllocation(&allocation)
		pixmap = gdk.Pixmap(drawingarea.GetWindow().GetDrawable(), allocation.Width, allocation.Height, 24)
		gc = gdk.GC(pixmap.GetDrawable())
		gc.SetRgbFgColor(gdk.Color("white"))
		pixmap.GetDrawable().DrawRectangle(gc, true, 0, 0, -1, -1)
		gc.SetRgbFgColor(gdk.Color("blue"))
		gc.SetRgbBgColor(gdk.Color("white"))
		pixmap.GetDrawable().DrawArc(gc,false,100,100,200,200,0,30000)
		pixmap.GetDrawable().DrawArc(gc,false,150,150,100,100,0,30000)
		
		pixmap.GetDrawable().DrawLine(gc,200,100,200,300)
		pixmap.GetDrawable().DrawLine(gc,100,200,300,200)
		
		pixmap.GetDrawable().DrawLine(gc,130,130,165,165)
		pixmap.GetDrawable().DrawLine(gc,235,235,270,270)
		
		
		pixmap.GetDrawable().DrawLine(gc,271,129,235,165)
		pixmap.GetDrawable().DrawLine(gc,165,235,129,271)
		
		pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,200,200,strconv.Itoa(prueba))
		
		pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,100,200,strconv.Itoa(prueba))
		
		pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,100,300,strconv.Itoa(prueba))
		pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,200,300,strconv.Itoa(prueba))
		
	})

  
  drawingarea.Connect("expose-event", func() {
		if pixmap != nil {
			drawingarea.GetWindow().GetDrawable().DrawDrawable(gc, pixmap.GetDrawable(), 0, 0, 0, 0, -1, -1)
			gc.SetRgbFgColor(gdk.Color("white"))
			pixmap.GetDrawable().DrawRectangle(gc, true, 0, 0, -1, -1)
			gc.SetRgbFgColor(gdk.Color("blue"))
			gc.SetRgbBgColor(gdk.Color("white"))
			pixmap.GetDrawable().DrawArc(gc,false,100,100,200,200,0,30000)
			pixmap.GetDrawable().DrawArc(gc,false,150,150,100,100,0,30000)
		
			pixmap.GetDrawable().DrawLine(gc,200,100,200,300)
			pixmap.GetDrawable().DrawLine(gc,100,200,300,200)
		    
			pixmap.GetDrawable().DrawLine(gc,130,130,165,165)
			pixmap.GetDrawable().DrawLine(gc,235,235,270,270)
		
		
			pixmap.GetDrawable().DrawLine(gc,271,129,235,165)
			pixmap.GetDrawable().DrawLine(gc,165,235,129,271)
		
		        pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,210,190,strconv.Itoa(r1[0]))			
			pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,210,140,strconv.Itoa(r2[0]))
			pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,260,190,strconv.Itoa(r2[1]))
			
			
			pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,210,230,strconv.Itoa(r1[1]))
			pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,260,230,strconv.Itoa(r2[2]))
			pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,210,280,strconv.Itoa(r2[3]))
			
			
				
			pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,170,230,strconv.Itoa(r1[2]))
			pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,170,280,strconv.Itoa(r2[4]))
			pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,120,230,strconv.Itoa(r2[5]))
			
			
			pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,170,190,strconv.Itoa(r1[3]))
			pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,120,190,strconv.Itoa(r2[6]))
			pixmap.GetDrawable().DrawString(gdk.FontsetLoad("-adobe-helvetica-bold-r-normal--12-120-75-75-p-70-iso8859-1"),gc,170,140,strconv.Itoa(r2[7]))
			
			
		        drawingarea.GetWindow().Invalidate(nil, false)
		}
	})
  
  
    
  button.Clicked(func(){
    //prueba=prueba+1    
    vdeep:=entry.GetText()
  
    
    fo,err:=os.Create("Pasos")
     if(err!=nil){panic(err)}
     defer fo.Close()
     
    Actual.tipoHeuristica=heuristica
    Actual.HeuristicaJuego()
    
    deep,err:=strconv.Atoi(vdeep)
    if(err==nil){      
      fmt.Println("",deep)
    b,camino:=AStar(Actual,deep)
    if(b){
     
    for i:=0;i<len(camino);i++{
     
     r1=camino[i].fila1
     r2=camino[i].fila2
     
     var Valores string
     Valores="["
     
     for j:=0;j<len(r1);j++{
	Valores=Valores+" "+strconv.Itoa(r1[j])
     }
     Valores=Valores+"]["
     
     for j:=0;j<len(r2);j++ {
	Valores=Valores+" "+strconv.Itoa(r2[j])
     }
          Valores=Valores+"] "+strconv.Itoa(camino[i].heuristica)+"\n"

     
     
     if _,err:=fo.WriteString(Valores); err!=nil {
       panic(err)
      }
      
    }
    
    }else{
      fmt.Println("No encontro la solución")
    }
    }
  })

  button2.Clicked(func(){
    
    Actual.Init()
    Actual.tipoHeuristica=heuristica
    Actual.HeuristicaJuego()
    
    r1=Actual.fila1
    r2=Actual.fila2
    
    //prueba=prueba+1    
    //fmt.Println("->",prueba)
  })
  
  button3.Clicked(func(){
    
    Actual.Init2()
    Actual.tipoHeuristica=heuristica
    Actual.HeuristicaJuego()
    
    r1=Actual.fila1
    r2=Actual.fila2
  })
  
  drawingarea.SetEvents(int(gdk.GDK_POINTER_MOTION_MASK | gdk.GDK_POINTER_MOTION_HINT_MASK | gdk.GDK_BUTTON_PRESS_MASK))
  framebox1.Add(drawingarea)
    
  
  window.Add(vbox)
  window.SetSizeRequest(600,600)
  window.ShowAll()
  gtk.Main()
}