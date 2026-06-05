package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	text   *List[*List[rune]] // a lista de linhas
	itLine *Node[*List[rune]] // iterador para a linha corrente
	itChar *Node[rune]        // iterador para o caracter do cursor
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.itChar = e.itLine.Value.Insert(e.itChar, r) // insere antes do elemento apontado pelo cursor
	e.itChar = e.itChar.Next()                    // move o cursor para próxima posição
}

func (e *Editor) KeyLeft() {
	if e.itChar != e.itLine.Value.Front() { // Se o cursor não está no início da linha
		e.itChar = e.itChar.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.itLine != e.text.Front() { // Se não está na primeira linha
		e.itLine = e.itLine.Prev()      // Atualiza iterador de linha para linha anterior
		e.itChar = e.itLine.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() { //
	linhaAtual := e.itLine.Value

	if e.itChar == linhaAtual.End() {
		e.text.Insert(e.itLine.Next(), NewList[rune]()) // cria uma nova linha e insere abaixo da linha corrente
		e.itLine = e.itLine.Next()        // vai pra próxima linha
		e.itChar = e.itLine.Value.Front() // move o cursor para o início da linha
		return
	}

		novaLista := NewList[rune]()
		final := linhaAtual.End()

		for i := e.itChar; i != final; i = i.Next(){
			novaLista.PushBack(i.Value)
		}

		for i := e.itChar; i != final; {
			next := i.Next()
			linhaAtual.Erase(i)
			i = next
		}

		e.text.Insert(e.itLine.Next(), novaLista)	

		e.itLine = e.itLine.Next()    // vai pra próxima linha
		e.itChar = e.itLine.Value.Front()

}

func (e *Editor) KeyRight() { //FEITO
	if e.itChar != e.itLine.Value.End() { // Se o cursor não está no final da linha
		e.itChar = e.itChar.Next() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.itLine != e.text.End() { // Se não está na ultima linha
		e.itLine = e.itLine.Next()      // Atualiza iterador de linha para linha seguinte
		e.itChar = e.itLine.Value.Front() // Move o cursor para o começp da linha
	}
}

func (e *Editor) KeyUp() {
	if e.itLine == e.text.Front() {
		return
	}

	cursor := e.itLine.Value.IndexOf(e.itChar)

	e.itLine = e.itLine.Prev()
	e.itChar = e.itLine.Value.Front()
	for cursor > 0 && e.itChar.Next() != e.itLine.Value.End(){
		e.itChar = e.itChar.Next()
		cursor--
	}
}

func (e *Editor) KeyDown() {
	if e.itLine == e.text.End() {
		return
	}

	cursor := e.itLine.Value.IndexOf(e.itChar)

	e.itLine = e.itLine.Next()
	e.itChar = e.itLine.Value.Front()
	for cursor > 0 && e.itChar.Next() != e.itLine.Value.End(){
		e.itChar = e.itChar.Next()
		cursor--
	}
}

func (e *Editor) KeyBackspace() {
	if e.itChar != e.itLine.Value.Front(){
		anterior := e.itChar.Prev()
		e.itLine.Value.Erase(anterior)
		return
	}
	if e.itLine != e.text.Front() && e.itChar == e.itLine.Value.Front(){
		linhaAnterior := e.itLine.Prev().Value
		noLinhaAnterior := e.itLine.Prev()
		linhaAtual := e.itLine.Value
		cursorNew := linhaAnterior.Back()
		for i := linhaAtual.Front(); i != linhaAtual.End(); i = i.Next(){
				linhaAnterior.PushBack(i.Value)
		}
		e.text.Erase(e.itLine)
		e.itLine = noLinhaAnterior
		e.itChar = cursorNew
		
	}
}

func (e *Editor) KeyDelete() {
	if e.itChar != e.itLine.Value.End(){
		anterior := e.itChar.Next()
		e.itLine.Value.Erase(anterior)
		return
	}
	if e.itLine.Next() != e.text.End() && e.itChar == e.itLine.Value.End(){
		linhaAtual := e.itLine.Value
		linhaAtualaa := e.itLine
		noLinhaPosterior := e.itLine.Next()
		linhaPosterior := e.itLine.Next().Value
		cursorNew := linhaAtual.Back()
		for i := linhaPosterior.Front(); i != linhaPosterior.End(); i = i.Next(){
			linhaAtual.PushBack(i.Value)
		}
		e.text.Erase(noLinhaPosterior)
		e.itLine = linhaAtualaa
		e.itChar = cursorNew
	}
}

func main() { // Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.text = NewList[*List[rune]]()
	e.text.PushBack(NewList[rune]())
	e.itLine = e.text.Front()
	e.itChar = e.itLine.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.text.Front(); line != e.text.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '↲'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.itChar {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
