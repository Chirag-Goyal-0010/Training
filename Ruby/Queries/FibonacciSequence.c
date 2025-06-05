#include <stdio.h>
int rec_fabonacci(int i){
    if(i==0 || i==1)
        return i;
    return rec_fabonacci(i-1) + rec_fabonacci(i-2);
}
int main(){
    int seq_no;
    scanf("%d",& seq_no);
    seq_no -= 1;
    while(seq_no >=0){    
        printf("%d  ",rec_fabonacci(seq_no));
        seq_no -= 1;
    }
    return 0;
}