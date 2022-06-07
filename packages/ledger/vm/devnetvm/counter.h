// NOTE: We can use https://michael-f-bryan.github.io/rust-ffi-guide/cbindgen.html to generate
// this header automatically from the Rust code.
#include <stdint.h>

struct Receipt
{
    uint8_t journal[2000];
    uint32_t seal[12];
};
struct CounterStation;
struct InitMessage;
struct SubmitCounterMessage
{
    struct Receipt receipt;
};

void init_stuff();
void hello(char *name);
struct CounterStation *create_counter_station();
struct InitMessage *counter_station_init(struct CounterStation *counter_station);
struct SubmitCounterMessage *counter_station_submit(struct CounterStation *counter_station);
const char *verify_and_get_commit_init(struct InitMessage *init_msg);
const char *verify_and_get_commit(struct SubmitCounterMessage *msg);