#include <ncurses.h>
#include <iostream>
#include <sstream>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <cstring>
#include <thread>
#include "st7920.h"
//#include "protobuf/drive.pb.h"
#include "screens/manual.hpp"

void curses_main();
void ecal_main(int argc, char **argv);

int main(int argc,char **argv) {
//    std::thread ecal_receiver(ecal_main,argc,argv);
//    std::thread curses_receiver(curses_main);
    initscr();
    resize_term(8, 22);
    update_screen();
    endwin();
    
    int file = open("/dev/glcd", O_RDWR);
    if (file == -1)
    {
        perror("Error opening file");
        return 1;
    }   
    struct ioctl_mesg message;
    ioctl(file, IOCTL_CLEAR_DISPLAY, &message);

    for (int i = 0; i < 8; ++i) {
    std::ostringstream frameContent;
    
        for (int j = 0; j < 22; ++j) {
            char character = mvwinch(stdscr, i, j) & A_CHARTEXT;
            if (isprint(character) || character == ' ' || character == '=' || character == '|') {
                frameContent << character;
            }
        }
    frameContent << "\r\n";
    std::string frameString = frameContent.str();
    memset(&message, 0, sizeof(struct ioctl_mesg));
    message.nthCharacter=i;
    if((i==0) || (i==7))
    {
        message.invert = 1;
    }
    mempcpy(&message.kbuf,frameString.c_str(),frameString.size());
    ioctl(file, IOCTL_PRINT_WITH_POSITION, &message);
    }       

    return 0;
}

void ecal_main(int argc, char **argv)
{
  /*
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
*/
}

void curses_main()
{

}