#ifndef _CPLUSPLUS_H_
#define _CPLUSPLUS_H_

#include <vector>
using namespace std;

class testVector{
    public:
        testVector(vector<int> input);
        int calcVectorValueSum();
    private:
        vector<int> vec;
};

#endif