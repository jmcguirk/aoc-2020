package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem3B struct {

}


func (this *Problem3B) Solve() {
	Log.Info("Problem 3B solver beginning!")


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

	slopes := make([]*IntVec2, 0);

	slope := &IntVec2{};
	slope.X = 1;
	slope.Y = 1;
	slopes = append(slopes, slope);

	slope = &IntVec2{};
	slope.X = 3;
	slope.Y = 1;
	slopes = append(slopes, slope);

	slope = &IntVec2{};
	slope.X = 5;
	slope.Y = 1;
	slopes = append(slopes, slope);


	slope = &IntVec2{};
	slope.X = 7;
	slope.Y = 1;
	slopes = append(slopes, slope);

	slope = &IntVec2{};
	slope.X = 1;
	slope.Y = 2;
	slopes = append(slopes, slope);

	product := 1;

	for _, slope := range slopes{
		posX := 0;
		posY := 0;

		deltaX := slope.X;
		deltaY := slope.Y;

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
			}

			steps++;
		}

		Log.Info("Completed toboggan simulation (%d, %d) after %d steps. %d trees encountered", deltaX, deltaY, steps, treesEncountered);
		product = product * treesEncountered;
	}

	Log.Info("Completed slopes simulation - product is %d", product);

}

