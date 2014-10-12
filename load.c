#define _DEFAULT_SOURCE  // shut up gcc: implicit declaration of function ‘getloadavg’
#include <unistd.h>  // sysconf
#include <stdlib.h>  // getloadavg
#include <stdio.h>   // printf
#include <argp.h>
#include <string.h>  // strsep


typedef struct load {
    double one_min;
    double five_min;
    double fifteen_min;
} load;


// how to display specific interval
typedef struct load_entry_s {
    double load_d;
    char * bg;
    char * fg;
} load_entry;


load
get_load(void) {
    signed int is_error;
    double loadavg[3];
    load load_s;

    is_error = getloadavg(loadavg, 3);

    if (is_error < 0) {
        printf("error getting loadavg");
        exit(1);
    }

    load_s.one_min = loadavg[0];
    load_s.five_min = loadavg[1];
    load_s.fifteen_min = loadavg[2];

    return load_s;
}

load_entry*
init_load_list(char *args[], int count) {
    int to_alloc = 3;
    if (count > 0) {
        to_alloc = count;
    }

    load_entry entries[to_alloc];

    if (count == 0) {
        entries = {
            { 0.05, "default", "green" },
            { 0.05, "default", "green" },
            { 0.20, "default", "red" }
        }
        return entries;
    }

    char *token, *string;
    char *triple[3];
    int loop_iteration;
    for (int i = 0; i < count; i++) {
        string = strdup(args[i]);
        loop_iteration = 0;
        while ((token = strsep(&string, ",")) != NULL) {
            triple[i] = token;
            ++loop_iteration;
        }
        entries[i] = (load_entry) { .load_d=0.0, .bg="", .fg=triple[1] };
        sscanf(triple[0], "%lf", &entries[i].load_d);
        if (loop_iteration == 3) {
            entries[i].bg = triple[2];
        }
    }
    return entries;
}
/*
void
print_load(load loadavg) {
    unsigned int number_of_cpus = sysconf(_SC_NPROCESSORS_ONLN);
    double cpus_d = (double) number_of_cpus;

    printf("%f\n", cpus_d);
    printf("%f\n", loadavg.one_min / cpus_d * 100.0 );

    if (loadavg.one_min / cpus_d * 100.0 > LOAD_HIGH ) {
        printf("High load: %.2f\n", loadavg.one_min);
    } else if (loadavg.one_min / cpus_d * 100.0 > LOAD_MED ) {
        printf("Medium load: %.2f\n", loadavg.one_min);
    } else if (loadavg.one_min / cpus_d * 100.0 > LOAD_LOW ) {
        printf("Low load: %.2f\n", loadavg.one_min);
    } else {
        printf("Minimal load: %.2f\n", loadavg.one_min);

    }

}
*/


/* ARGPARSE */

const char *argp_program_version =
  "load 0.1";
const char *argp_program_bug_address =
  "github.com/TomasTomecek/tmux-status-helpers/";

/* Program documentation. */
static char doc[] =
  "doc";

/* A description of the arguments we accept. */
static char args_doc[] = "[THRESHOLD_ENTRY THRESHOLD_ENTRY...]";

/* The options we understand. */
static struct argp_option options[] = {
    { 0 }
};

/* Used by main to communicate with parse_opt. */
struct arguments
{
  char *args[1024];
  int args_count;
};

/* Parse a single option. */
static error_t
parse_opt (int key, char *arg, struct argp_state *state)
{
  /* Get the input argument from argp_parse, which we
     know is a pointer to our arguments structure. */
  struct arguments *arguments = state->input;

  switch (key)
    {
    case ARGP_KEY_ARG:
      arguments->args[state->arg_num] = arg;
      ++arguments->args_count;
      break;

    case ARGP_KEY_END:
      break;

    default:
      return ARGP_ERR_UNKNOWN;
    }
  return 0;
}

/* Our argp parser. */
static struct argp argp = { options, parse_opt, args_doc, doc, 0, 0, 0 };


/* MAIN */


int
main(int argc, char **argv) {
    //load_entry * entries = malloc(300 * sizeof(load_entry));
    //entries = &(load_entry) { .load_d=0.05, .bg="default", .fg="green" };
    //printf("%f", entries->load_d);
    //*(entries + 1) = (load_entry) { .load_d=0.05, .bg="default", .fg="green" };
    //*(entries + 2) = (load_entry) { .load_d=0.2, .bg="default", .fg="red" };
    //init_load_list(&entries);

    //print_load(get_load());

    struct arguments arguments;
    arguments.args_count = 0;
    argp_parse(&argp, argc, argv, 0, 0, &arguments);

    init_load_list(arguments.args, arguments.args_count);

    return 0;
}
