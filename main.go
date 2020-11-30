package main

func main() {
	Log.Init();
	Log.Info("Starting up AOC 2020");

	solver := Problem1A{};

	solver.Solve();
	Log.Info("Solver complete - exiting");
}
