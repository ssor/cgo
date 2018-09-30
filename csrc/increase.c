#include <stdio.h>
#include <stdlib.h>
#include <string.h>
//#include "increase.h"

void increaseSingle(unsigned char *s){
    (*s) = (*s)+1;
    return;
}

int increaseu(unsigned char *s){
    int count;
    size_t len = sizeof(s) / sizeof(s[0]);
    for(count = 0; count < len; count++){
        s[count] = s[count] + 1;
    }
    return len;
}

int increase(char *s){
    int count;
    int len = strlen(s);
    for(count = 0; count < len; count++){
        s[count] = s[count] + 1;
    }
    return len;
}