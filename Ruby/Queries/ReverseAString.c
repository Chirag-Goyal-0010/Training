#include <stdio.h>
#include <string.h>
int main(){
    char str_1[50], rev_str[50];
    scanf("%s", str_1);
    int lenght = strlen(str_1);
    for(int i = 0; i<lenght; i++){
        rev_str[i] = str_1[lenght-1-i];
    }
    rev_str[lenght] = '\0';
    printf("%s",rev_str);
    return 0;
}