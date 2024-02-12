#include <ncurses.h>
#include "manual.hpp"

void update_screen()
{
    mvprintw(0, 1, "V00.00.01 - manual");
    mvprintw(1, 1, "Dr.Nr 0  M-OP");
    mvprintw(2, 1, "Enc. Status:");
    mvprintw(3, 1, "uS-Mode:");
    mvprintw(4, 1, "Setpoint:");
    mvprintw(5, 1, "Drv-Pos.:");
    mvprintw(6, 1, "");
    mvprintw(7, 1, "Firmware:");
    refresh();
    //getch();

}