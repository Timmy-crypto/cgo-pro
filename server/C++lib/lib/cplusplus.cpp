#include "cplusplus.hpp"
#include <iostream>

using namespace std;

testVector::testVector(vector<int> input){
    vec = input;
}

int testVector::calcVectorValueSum(){
    int sum = 0;
    vector<int>::iterator iter;
    for (iter = vec.begin();iter != vec.end();iter++){
        sum += *iter;
    }

    cout<<"the vector sum is:"<<sum<<endl;
    return sum;
}
