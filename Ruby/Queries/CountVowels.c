#include <stdio.h>
#include <string.h>
int main(){
    printf("enter the string under 100 characters \n");
    char str[100];
    int sum = 0;
    scanf("%s", str);
    for(int i = 0; i<strlen(str);i++){
        if(str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' 
        || str[i] == 'A' || str[i] == 'E' || str[i] == 'I' || str[i] == 'O' || str[i] == 'U'){
            sum = sum +1;
        }
    }
    printf("%d",sum);
    return 0;
}