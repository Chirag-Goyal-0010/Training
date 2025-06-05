#include <stdio.h>
int main(){
    char str[50];
    int length = 0;
    scanf("%s",str);
    char sub_str[50];
    for(int i = 0;str[i] != '\0'; i++){
        for(int j = 0; sub_str[j] != '\0'; j++){
            if(str[i] == sub_str[j]){
                sub_str[length] = '\0';
                goto end;
            }
        }
        sub_str[length] = str[i];
        length += 1;
    }
    end :
        printf("%s", sub_str);
    return 0;
}