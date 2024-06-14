#include <stdlib.h>
#include <stdio.h>

void f(void)
{
  int* x = malloc(10 * sizeof(int));
  x[10] = 0;        // problem 1: heap block overrun
}                    // problem 2: memory leak -- x not freed

int main(void)
{
  printf("start\n");
  f();
  printf("end\n");
  return 0;
}