#include "wrapper.hpp"
#include "cplusplus.hpp"
#include <stdio.h>
#include <vector>

int GetVectorDataSum(int* data,int dataLen){
    std::vector<int> tmpVector;
    int i;
    for(i = 0; i < dataLen; i++){
        //printf("the data is:%d\r\n",data[i]);
        //printf("the data addr is:%p\r\n",&data[i]);
        tmpVector.push_back(data[i]);
    }

    testVector m_vector(tmpVector);

    int sum;
    sum=m_vector.calcVectorValueSum();
    printf("the sum data is:%d\r\n",sum);

    return sum;
}