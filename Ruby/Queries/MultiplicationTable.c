#include <stdio.h>
int main(){
    printf("enter the no. \n");
    int num;
    scanf("%d", &num);
    for(int i = 1; i < 11; i++){
        printf("%d * %d = %d \n",num,i,num*i);
    }
    return 0;
}