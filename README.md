# Bytebeat
Bytebeat это довольно своеобразная музыка без каких-либо партитуры и инструментов, которая была изобретена в Сентябре 2011 года. Представляет из себя просто однострочную формулу, определяющую форму сигнала как функцию времени. Если вы поместите эту формулу в программу с циклом, который каждую итерацию увеличивает переменную времени, то на выходе вы получите 8-битный беззнаковый монофонический звук с частотой 8 кГц. Bytebeat музыка довольно оптимизированна и зачастую может работать даже на самых слабых встраеваемых устройствах.
## Компиляция
`gcc main.c -pedantic -Wall -Wextra -O3 -lmingw32 -lSDL2main -lSDL2 -o bytebeat`
