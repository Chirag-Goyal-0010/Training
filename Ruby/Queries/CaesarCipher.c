#include <stdio.h>
void encrypted(char str[], int shift){
    for(int i= 0; i < strlen(str); i++){
        if (str[i] > 96){
            if ((str[i] + shift) > 122){
                str[i] = str[i] + shift - 26;
            }
            else {
                str[i] = str[i] + shift;
            }
        }
        else if (str[i] < 91){
            if ((str[i] + shift) > 90){
                str[i] = str[i] + shift - 26;
            }
            else {
                str[i] = str[i] + shift;
            }
        }
    }
    return str;
}
void decryption(char str[], int shift){
    for(int i= 0; i < strlen(str); i++){
        if (str[i] > 96){
            if ((str[i] - shift) < 97){
                str[i] = str[i] - shift + 26;
            }
            else {
                str[i] = str[i] - shift;
            }
        }
        else if (str[i] < 91){
            if ((str[i] - shift) < 65){
                str[i] = str[i] - shift + 26;
            }
            else {
                str[i] = str[i] - shift;
            }
        }
    }
    return str;
}
int main(){
    char str[100];
    printf("Enter the string \n");
    scanf("%s",str);
    printf("Enter the shift value \n");
    int Shift_Value;
    scanf("%d", & Shift_Value );
    encrypted(str,Shift_Value);
    printf("the encrypted cipher code is \n %s \n",str);
    decryption(str,Shift_Value);
    printf("the decrypted cipher code is \n %s \n",str);
    return 0;
}