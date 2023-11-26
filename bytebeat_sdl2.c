/*
 * bytebeat_sdl2.c • 26-11-2023 • i4k • https://github.com/i4k1/bytebeat
 *
 * Bytebeat is a kind of music without a score or instruments, which
 * was invented in September 2011. It is simply a one-line formula that
 * determines the waveform as a function of time. If you put this
 * formula in a program with a loop that increments the time variable
 * each iteration, the output is 8-bit unsigned monophonic audio at
 * 8 kHz. Bytebeat music is quite optimized and can often work even on
 * the weakest embedded devices.
 *
 * This code is a simple example of such music.
 *
 */

#include <stdio.h>
#include <SDL2/SDL.h>

int main(int argc, char **argv) {
    SDL_Init(SDL_INIT_AUDIO);
    SDL_AudioSpec want, have;
    want.freq = 32000;
    want.format = AUDIO_U8;
    want.channels = 1;
    want.samples = 2048;
    want.callback = NULL;

    SDL_AudioDeviceID device = SDL_OpenAudioDevice(NULL, 0, &want, &have, 0);
    SDL_PauseAudioDevice(device, 0);

    unsigned int t = 0;
    while (1) {
        unsigned int sample = ( /* paste your formula here */ );
        sample = (sample << 16) | (sample & 0xFFFFFFFF);
        SDL_QueueAudio(device, &sample, sizeof(sample));
        t++;
    }

    SDL_CloseAudioDevice(device);
    SDL_Quit();

    return 0;
}
