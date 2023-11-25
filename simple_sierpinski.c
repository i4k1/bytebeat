#include <stdio.h>
#include <SDL2/SDL.h>

#define SAMPLE_RATE 32000

int main(int argc, char **argv) {
    SDL_Init(SDL_INIT_AUDIO);

    SDL_AudioSpec want;
    want.freq = SAMPLE_RATE;
    want.format = AUDIO_U8;
    want.channels = 1;
    want.samples = 2048;
    want.callback = NULL;

    SDL_AudioDeviceID device = SDL_OpenAudioDevice(NULL, 0, &want, NULL, 0);
    SDL_PauseAudioDevice(device, 0);
    
    unsigned int t = 0;
    while (t >= 0) {
        unsigned int sample = t&t>>8;
        sample = (sample << 16) | (sample & 0xFFFFFFFF);
        SDL_QueueAudio(device, &sample, sizeof(sample));
        t++;
    }
    
    SDL_CloseAudioDevice(device);
    SDL_Quit();
    
    return 0;
}
