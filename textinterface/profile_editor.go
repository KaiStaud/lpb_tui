package tui

/*
* The profile editor allowes the user to change database entries during runtime.
* Commands are retrieved from input and parsed into seperate units:
* The command itself: print (p),create (c), overwrite (o), delete (d), send (s)
* Argument 1: ID of choosen database entry
* Argument 2: String with coordinates e.g.: x<100>,y<200>,z<300>
*
* If the command doesn't requires any input arguments, the passed ones are ignored.
* The Editor returns an error message on incorrect input
 */
