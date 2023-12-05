// 05-12-2023 i4k
// ./a.out | aplay

#include <stdio.h>

int main(void) {
    unsigned int t = 0;
    for (;;) {
        sample = ( /* paste your formula here */ ); // for example: t&t>>8
        output = (sample << 16) | (sample & 0xffffffff);
        putchar(output);
        t++;
    }
    return 0;
}
