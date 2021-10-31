// calculator

package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(500, 280))
	output := ""
	input := widget.NewLabel(output)

	historyStr:= ""
	history := widget.NewLabel(historyStr);
	isHistory := false
	var historyArr []string;
	historyBtn := widget.NewButton("History", func(){
		if isHistory{
			historyStr=""
		}else{
			for i := len(historyArr)-1; i >=0; i--{
				historyStr = historyStr + historyArr[i];
				historyStr += "\n";
			}

		}
		isHistory = !isHistory;
		history.SetText(historyStr)
	})
	backBtn := widget.NewButton("Back", func(){
		if len(output)>0{
			output = output[:len(output)-1];
			input.SetText(output)
		}
	})
	clearBtn := widget.NewButton("Clear", func(){
		output = "";
		input.SetText(output);
	})
	openBtn := widget.NewButton("(", func(){
		output = output + "(";
		input.SetText(output)
	})
	close := widget.NewButton(")", func(){
		output = output + ")";
		input.SetText(output)
	})
	divide := widget.NewButton("/", func(){
		output = output + "/";
		input.SetText(output)
	})
	seven := widget.NewButton("7", func(){
		output = output + "7";
		input.SetText(output)
	})
	eight := widget.NewButton("8", func(){
		output = output + "8";
		input.SetText(output)
	})
	nine := widget.NewButton("9", func(){
		output = output + "9";
		input.SetText(output)
	})
	mul := widget.NewButton("*", func(){
		output = output + "*";
		input.SetText(output)
	})
	four := widget.NewButton("4", func(){
		output = output + "4";
		input.SetText(output)
	})
	five := widget.NewButton("5", func(){
		output = output + "5";
		input.SetText(output)
	})
	six := widget.NewButton("6", func(){
		output = output + "6";
		input.SetText(output)
	})
	minus := widget.NewButton("-", func(){
		output = output + "-";
		input.SetText(output)
	})
	one := widget.NewButton("1", func(){
		output = output + "1";
		input.SetText(output)
	})
	two := widget.NewButton("2", func(){
		output = output + "2";
		input.SetText(output)
	})
	three := widget.NewButton("3", func(){
		output = output + "3";
		input.SetText(output)
	})
	plus := widget.NewButton("+", func(){
		output = output + "+";
		input.SetText(output)
	})
	zero := widget.NewButton("0", func(){
		output = output + "0";
		input.SetText(output)
	})
	dot := widget.NewButton(".", func(){
		output = output + ".";
		input.SetText(output)
	})
	equal := widget.NewButton("=", func(){
		expression, err := govaluate.NewEvaluableExpression(output);
		if err == nil{
			result, err := expression.Evaluate(nil);
			if err == nil{
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64);
				strToAppend := output + " = " + ans;
				historyArr = append(historyArr, strToAppend);
				output = ans;
			}else{
				output = "error";
			}
		}else{
			output = "error";
		}
		input.SetText(output)
	})

	w.SetContent(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
			historyBtn,
			backBtn,
			),
			container.NewGridWithColumns(4,
				clearBtn,
				openBtn,
				close,
				divide,
			),
			container.NewGridWithColumns(4,
			seven,
			eight,
			nine,
			mul,
			),
			container.NewGridWithColumns(4,
			four,
			five,
			six,
			minus,
			),
			container.NewGridWithColumns(4,
			one,
			two,
			three,
			plus,),
		),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zero,
					dot,),
					equal,
			),
	),)

	w.ShowAndRun()
}

