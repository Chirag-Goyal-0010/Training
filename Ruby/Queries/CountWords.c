#include <stdio.h>
#include <string.h>
int main(){
    printf("input a string \n");
    char str[100];
    int words = 0;
    fgets(str, sizeof(str), stdin);
    printf("%s",str);
    for(int i = 0; i<sizeof(str);i++){
        if(str[i] == ' '){
            words = words + 1;
        }
    }
    printf("total words are %d",words +1);
    return 0;
}