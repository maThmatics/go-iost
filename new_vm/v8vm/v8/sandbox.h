#ifndef IOST_V8_SANDBOX_H
#define IOST_V8_SANDBOX_H

#include "v8.h"
#include "vm.h"

using namespace v8;

typedef struct {
  Persistent<Context> context;
  Isolate *isolate;
  const char *jsPath;
  size_t gasUsed;
  size_t gasLimit;
} Sandbox;

extern ValueTuple Execution(SandboxPtr ptr, const char *code);

#endif // IOST_V8_SANDBOX_H