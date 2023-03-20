#include <uuid.h>

int main(int argc, char **argv) {
  uuid_t binuuid;
  uuid_generate_random(binuuid);

  char *uuid = malloc(37);
  uuid_unparse_lower(binuuid, uuid);

#ifdef CGI
  printf("Content-Type: application/json\n");
  printf("\n");
#endif

  printf("{\"data\":\"%s\",\"base64\":false,\"headers\":{},\"kv\":{},\"status\":200}\n", uuid);

  return 0;
}
