#include <stdio.h>
#include <stdlib.h>
#include <uuid.h>

int main(int argc, char **argv) {
  uuid_t binuuid;
  uuid_generate_random(binuuid);

  char *uuid = malloc(37);
  uuid_unparse_lower(binuuid, uuid);

#ifdef CGI
  printf("Content-Type: text/plain\n\n");
  printf("%s\n", uuid);
#elif defined(WWS)
  printf("{\"data\":\"{\\\"uuid\\\":\\\"%s\\\"}\",\"base64\":false,\"headers\":{\"Content-Type\": \"application/json\"},\"kv\":{},\"status\":200}\n", uuid);
#else
  printf("%s\n", uuid);
#endif

  return 0;
}
