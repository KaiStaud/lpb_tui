#define MAX_BUF_LENGTH 1042
#define IOCTL_CLEAR_DISPLAY '0'
#define IOCTL_PRINT '1'
#define IOCTL_PRINT_WITH_POSITION '3'
#define IOCTL_PRINT_BMP '4'
#define IOCTL_RESET '5'
#define IOCTL_BACKLIGHT_ON '6'
#define IOCTL_BACKLIGHT_OFF '7'

struct ioctl_mesg
{
    char kbuf[MAX_BUF_LENGTH];
    unsigned int nthCharacter;
    unsigned int lineNumber;
    unsigned int nbytes;
    unsigned int invert;
};