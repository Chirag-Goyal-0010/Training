# include <stdio.h>
#include <string.h>
int main(){
    char str_1[25];
    scanf("%s",& str_1);
    // if(str.lenght >25)
    int len = strlen(str_1);
    int indicator = 0;
    for(int i = 0; i <len;i++){
        if(str_1[i] != str_1[len-i-1]){
            printf("Not palindrome");
            indicator += 1;
            break;
        }
    }
    if(indicator == 0)
    printf("Palindrome");
    return 0;
}