#include <stdio.h>
int rec_gcd(int a, int b){
    if(a % b == 0){
        return b;
    }
    int c = a % b;
    rec_gcd(b , c);
}
int main(){
    printf("enter the 2 num:  ");
    int a,b;
    scanf("%d \n %d",&a,&b);
    if(a < b){
        int temp = a;
        a = b;
        b = temp;
    }
    int gcd = rec_gcd(a , b);
    printf("%d", gcd);
    return 0;
}