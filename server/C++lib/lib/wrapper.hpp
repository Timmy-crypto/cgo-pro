#ifndef _WRAPPER_H_
#define _WRAPPER_H_

#ifdef __cplusplus
// extern "C" is needed so the C++ compiler exports the symbols without name
// manging.
extern "C" {
#endif

int GetVectorDataSum(int* data,int dataLen);

#ifdef __cplusplus
}
#endif

#endif