package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem3A struct {

}

const TreeCode = int('#');
const EmptyCode = int('.');

func (this *Problem3A) Solve() {
	Log.Info("Problem 3A solver beginning!")


	file, err := os.Open("source-data/input-day-03a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	grid := &IntegerGrid2D{};
	grid.Init();


	scanner := bufio.NewScanner(file)

	y := 0;
	x := 0;
	for scanner.Scan() {
		lineRaw := strings.TrimSpace(scanner.Text());
		if(lineRaw != ""){
			for _, c := range lineRaw{
				val := int(c);
				grid.SetValue(x, y, val);
				x++;
			}
		}
		x = 0;
		y ++;
	}

	Log.Info("Loaded Grid\n%s", grid.PrintAscii());

	posX := 0;
	posY := 0;

	deltaX := 3;
	deltaY := 1;

	treesEncountered := 0;

	maxX := grid.MaxX();
	maxY := grid.MaxY();
	steps := 0;
	for {
		posX = posX + deltaX;
		posX = posX % (maxX+1);
		posY = posY + deltaY;

		if(posY > maxY){
			break;
		}


		val := grid.GetValue(posX, posY);
		if(val == TreeCode){
			treesEncountered++;
			Log.Info("%d.) Found tree at %d,%d", steps, posX, posY);
		} else{
			Log.Info("%d.) Empty at %d,%d", steps, posX, posY);
		}

		steps++;
	}

	Log.Info("Completed toboggan simulation after %d steps. %d trees encountered", steps, treesEncountered);
}
