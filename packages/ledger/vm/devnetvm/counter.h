// NOTE: We can use https://michael-f-bryan.github.io/rust-ffi-guide/cbindgen.html to generate
// this header automatically from the Rust code.

struct CounterStation;
struct InitMessage;
struct SubmitCounterMessage;

void init_stuff();
void hello(char* name);
struct CounterStation* create_counter_station();
struct InitMessage* counter_station_init(struct CounterStation* counter_station);
struct SubmitCounterMessage* counter_station_submit(struct CounterStation* counter_station);
const char* verify_and_get_commit_init(struct InitMessage* init_msg);
const char* verify_and_get_commit(struct SubmitCounterMessage* msg);