```
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include<stdio.h>
void unzip(const char  * input , char * output){
int sameCount = 0;
int count =0;
int outputCount  = 0 ;

int firstFlag = 1;
while(*input  != 0  &&  sameCount + count <=2 ){  // 如果是数字就清空标志位和重置计算

        if(*input >= '1' && *input <='9'){
                count = count * 10 +  atoi(input);
                sameCount = 0;
                //diffIndex = 0;
        }else{
                // 输出start
                if( input[-1] >= '1' && input[-1] <='9'){
                        for(int i = 0 ; i < count ;  i ++){
                                output [outputCount] = *input;
                                outputCount ++;
                        }
                }else{
                        // 输出start
                        if(firstFlag){
                                sameCount =1;
                                firstFlag = 0;
                                output[outputCount] = * input;
                                outputCount ++;
                                input ++;
                                continue;
                        }else{
                                output[outputCount] = * input;
                                outputCount ++;
                        }
                }
                // 输出end
                if(input[-1] == input[0]){
                        sameCount ++;
                        count =0;
                }else{
                        sameCount = 1 ;
                        count = 0 ;
                }

        }
        input ++;
        if(firstFlag){
                firstFlag = 0;
        }

}
if(sameCount+ count >2){

        memcpy(output,"error",6) ;
}

}
int main(){
char * input ="5a2bbcc";
char output[100];
unzip(input,output);
printf("%s",output);
return 0;

    
```
