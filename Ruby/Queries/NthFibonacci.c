#include <stdio.h>
int rec_fabonacci(int i){
    if(i==0 || i==1)
        return i;
    return rec_fabonacci(i-1) + rec_fabonacci(i-2);
}

int main()
{
    printf("Enter the seq no. of where you want to find fabonacci \n");
    int num;
    scanf("%d", &num);
    printf("The num is %d",rec_fabonacci(num));

    return 0;
}