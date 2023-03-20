#include <uuid.h>

int main(int argc, char **argv) {
  uuid_t binuuid;
  uuid_generate_random(binuuid);

  char *uuid = malloc(37);
  uuid_unparse_lower(binuuid, uuid);

  printf("%s\n", uuid);

  return 0;
}
