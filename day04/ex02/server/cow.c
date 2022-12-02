
#include "cow.h"

unsigned int i;
unsigned int argscharcount = 0;

char *ask_cow(char phrase[]) {
  int phrase_len = strlen(phrase);
  char *buf = (char *)malloc(sizeof(char) * (160 + (phrase_len + 2) * 3));
  strcpy(buf, " ");

  for (i = 0; i < phrase_len + 2; ++i) {
    strcat(buf, "_");
  }

  strcat(buf, "\n< ");
  strcat(buf, phrase);
  strcat(buf, " ");
  strcat(buf, ">\n ");

  for (i = 0; i < phrase_len + 2; ++i) {
    strcat(buf, "-");
  }
  strcat(buf, "\n");
  strcat(buf, "        \\   ^__^\n");
  strcat(buf, "         \\  (oo)\\_______\n");
  strcat(buf, "            (__)\\       )\\/\\\n");
  strcat(buf, "                ||----w |\n");
  strcat(buf, "                ||     ||\n");
  return buf;
}

//int main(int argc, char *argv[]) {
//  for (i = 1; i < argc; ++i) {
//    argscharcount += (strlen(argv[i]) + 1);
//  }
//  argscharcount = argscharcount + 1;
//
//  char *phrase = (char *)malloc(sizeof(char) * argscharcount);
//  strcpy(phrase, argv[1]);
//
//  for (i = 2; i < argc; ++i) {
//    strcat(phrase, " ");
//    strcat(phrase, argv[i]);
//  }
//  char *cow = ask_cow(phrase);
//  printf("%s", cow);
//  free(phrase);
//  free(cow);
//  return 0;
//}