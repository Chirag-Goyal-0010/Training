#include <stdio.h>
int main(){
    printf("Enter degree in celcius \n");
    float celcius;
    scanf("%f", & celcius);
    float fahrenheit = (celcius * 9/5) + 32;
    printf("%f°F" , fahrenheit);
    return 0;
}