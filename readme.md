# Go 与 C 的交互编程

## go和c之间的类型转换
```
char -->  C.char -->  byte
signed char -->  C.schar -->  int8
unsigned char -->  C.uchar -->  uint8
short int -->  C.short -->  int16
short unsigned int -->  C.ushort -->  uint16
int -->  C.int -->  int
unsigned int -->  C.uint -->  uint32
long int -->  C.long -->  int32 or int64
long unsigned int -->  C.ulong -->  uint32 or uint64
long long int -->  C.longlong -->  int64
long long unsigned int -->  C.ulonglong -->  uint64
float -->  C.float -->  float32
double -->  C.double -->  float64
wchar_t -->  C.wchar_t  -->
void * -> unsafe.Pointer

```

## 一些关键点

1. 对于C来说，可以认为没有字符串，只有数字的数组（所有的语言其实类似）

2. 传递一个指针到C函数中，需要现在Go中完成初始化，之后获取数组第一个成员的指针，将其转化为相应的指针进行传递


## 参考


ref: https://blog.csdn.net/FreeApe/article/details/51885308

