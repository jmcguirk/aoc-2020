package main

func main() {
	Log.Init();
	Log.Info("Starting up AOC 2020");

	solver := Problem3B{};

	solver.Solve();
	Log.Info("Solver complete - exiting");
}
