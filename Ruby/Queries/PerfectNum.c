#include <stdio.h>
void isPerfect(int num){
    int i = 1, arr[100], lenght = 0,sum = 0;
    while (i < num){
        if(num % i == 0){
            arr[lenght] = i;
            lenght ++;
        }
        i++;
    }
    for(int i = 0; i < lenght; i++){
        sum = sum + arr[i];
    }
    if(sum == num){
        printf("Perfect num");
    }
    else {
        printf("Not Perfect");
    }
}
int main(){
    printf("Enter num \n");
    int num;
    scanf("%d", & num);
    isPerfect(num);
    return 0;
}