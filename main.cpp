#include <ecal/ecal.h>
#include <ecal/msg/string/subscriber.h>
#include <CLI/CLI.hpp>
#include <panel.h>
#include <string.h>
#include <thread>

#include "protobuf/drive.pb.h"

typedef struct _PANEL_DATA {
	int hide;	/* TRUE if panel is hidden */
}PANEL_DATA;

#define NLINES 12
#define NCOLS 40
#define BUILDVERSION "V00.00.00"

void init_wins(WINDOW **wins, int n);
void win_show(WINDOW *win, char *label, int label_color);
void print_in_middle(WINDOW *win, int starty, int startx, int width, char *string, chtype color);
void curses_main();
void ecal_main(int argc, char **argv);
void print_win1(WINDOW *win);

std::string set_position = "100" ;

enum class ActivePanel{
	none,
	manual,
	automatic,
	systems,
};

WINDOW *my_wins[3];
PANEL  *my_panels[3];
ActivePanel active_panel;
std::vector<std::string>subscriptions;

int main(int argc,char **argv)
{		
	PANEL_DATA panel_datas[3];
	PANEL_DATA *temp;
	int ch;
    CLI::App app{"Text Interface for CANOpen Network"};
	app.set_config("--config")->required();
	app.add_option("--servers",subscriptions,"Listen to Servers");

    CLI11_PARSE(app, argc, argv);
	
	/* Initialize curses */
	initscr();
	start_color();
	cbreak();
	noecho();
	keypad(stdscr, TRUE);

	/* Initialize all the colors */
	init_pair(1, COLOR_RED, COLOR_BLACK);
	init_pair(2, COLOR_GREEN, COLOR_BLACK);
	init_pair(3, COLOR_BLUE, COLOR_BLACK);
	init_pair(4, COLOR_CYAN, COLOR_BLACK);

	init_wins(my_wins, 3);
	
	/* Attach a panel to each window */ 	/* Order is bottom up */
	my_panels[0] = new_panel(my_wins[0]); 	/* Push 0, order: stdscr-0 */
	my_panels[1] = new_panel(my_wins[1]); 	/* Push 1, order: stdscr-0-1 */
	my_panels[2] = new_panel(my_wins[2]); 	/* Push 2, order: stdscr-0-1-2 */

	/* Initialize panel datas saying that nothing is hidden */
	panel_datas[0].hide = FALSE;
	panel_datas[1].hide = FALSE;
	panel_datas[2].hide = FALSE;

	set_panel_userptr(my_panels[0], &panel_datas[0]);
	set_panel_userptr(my_panels[1], &panel_datas[1]);
	set_panel_userptr(my_panels[2], &panel_datas[2]);

	/* Update the stacking order. 2nd panel will be on top */
	update_panels();

	/* Show it on the screen */
	attron(COLOR_PAIR(4));
	mvprintw(LINES - 3, 0, "Show or Hide a window with 'a'(first window)  'b'(Second Window)  'c'(Third Window)");
	mvprintw(LINES - 2, 0, "F1 to Exit");

	attroff(COLOR_PAIR(4));
	doupdate();

  std::thread ecal_receiver(ecal_main,argc,argv);
  std::thread curses_receiver(curses_main);

	while((ch = getch()) != KEY_F(1))
	{	switch(ch)
		{	case 'a':			
				temp = (PANEL_DATA *)panel_userptr(my_panels[0]);
					print_win1(my_wins[0]);
					show_panel(my_panels[0]);
					hide_panel(my_panels[1]);
					hide_panel(my_panels[2]);
					temp->hide = FALSE;
					active_panel = ActivePanel::manual;
				break;
			case 'b':
				temp = (PANEL_DATA *)panel_userptr(my_panels[1]);
					show_panel(my_panels[1]);
					hide_panel(my_panels[0]);
					hide_panel(my_panels[2]);
					temp->hide = FALSE;
					active_panel = ActivePanel::automatic;
				break;
			case 'c':
				temp = (PANEL_DATA *)panel_userptr(my_panels[2]);
					show_panel(my_panels[2]);
					hide_panel(my_panels[1]);
					hide_panel(my_panels[0]);
					temp->hide = FALSE;
					active_panel = ActivePanel::systems;
				break;
		}
		update_panels();
		doupdate();
	}
	endwin();
	return 0;
}

/* Put all the windows */
void init_wins(WINDOW **wins, int n)
{	int x, y, i;
	char label[80];
	y = 2;
	x = 1;
	char *headings[3]={"manual","automatic","systems"};
	for(i = 0; i < n; ++i)
	{	wins[i] = newwin(NLINES, NCOLS, y, x);
//		wresize(wins[i],14,50);
		sprintf(label, ""BUILDVERSION" - %s",headings[i]);
		win_show(wins[i], label, i + 1);

	}
}

/* Show the window with a border and a label */
void win_show(WINDOW *win, char *label, int label_color)
{	int startx, starty, height, width;

	getbegyx(win, starty, startx);
	getmaxyx(win, height, width);

	box(win, 0, 0);
	mvwaddch(win, 2, 0, ACS_LTEE); 
	mvwhline(win, 2, 1, ACS_HLINE, width - 2); 
	mvwaddch(win, 2, width - 1, ACS_RTEE); 
	
	print_in_middle(win, 1, 0, width, label, COLOR_PAIR(label_color));
}

void print_in_middle(WINDOW *win, int starty, int startx, int width, char *string, chtype color)
{	int length, x, y;
	float temp;

	if(win == NULL)
		win = stdscr;
	getyx(win, y, x);
	if(startx != 0)
		x = startx;
	if(starty != 0)
		y = starty;
	if(width == 0)
		width = 80;

	length = strlen(string);
	temp = (width - length)/ 2;
	x = startx + (int)temp;
	wattron(win, color);
	mvwprintw(win, y, x, "%s", string);
	wattroff(win, color);
	refresh();
}

void print_win1(WINDOW *win)
{
	#define Y_OFFSET 2
	#define X_OFFSET 1
    mvwprintw(win,Y_OFFSET+1,X_OFFSET,"Drive-Nummer + M-OP");
    mvwprintw(win,Y_OFFSET+2,X_OFFSET,"Encoder-Status:");
    mvwprintw(win,Y_OFFSET+3,X_OFFSET,"Microstepping");
    mvwprintw(win,Y_OFFSET+4,X_OFFSET,"Set-Position: %s",set_position.c_str());
    mvwprintw(win,Y_OFFSET+5,X_OFFSET,"Drive-Position:");
    attron(A_REVERSE);
    mvwprintw(win,Y_OFFSET+7,X_OFFSET, "Zephyr Build:");
    attroff(A_REVERSE);
}

void ecal_main(int argc, char **argv)
{
  // initialize eCAL API
  eCAL::Initialize(argc, argv, "minimal_rec");

  eCAL::string::CSubscriber<std::string> sub("Hello");
  while (eCAL::Ok())
  {
    if (sub.Receive(set_position, nullptr, 100))
    {
		if(active_panel == ActivePanel::manual)
		{
		print_win1(my_wins[0]);
		show_panel(my_panels[0]);
		update_panels();
		doupdate();
		}
	}
  }
}

void curses_main()
{

}