package clib

/*
extern int
xmlInputMatchCallbackFunc(const char*);

extern void*
xmlInputOpenCallbackFunc(const char*);

extern int
xmlInputReadCallbackFunc(void *, char *, int);

extern int
xmlInputCloseCallbackFunc(void *);

int
xmlInputMatchCallbackFuncWrapper(const char *filename)
{
	return xmlInputMatchCallbackFunc(filename);
}

void*
xmlInputOpenCallbackFuncWrapper(const char *filename)
{
	return xmlInputOpenCallbackFunc(filename);
}

int
xmlInputReadCallbackFuncWrapper(void *context, char *buffer, int len)
{
	return xmlInputReadCallbackFunc(context, buffer, len);
}

int
xmlInputCloseCallbackFuncWrapper(void *context)
{
	return xmlInputCloseCallbackFunc(context);
}
*/
import "C"
