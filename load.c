#define _DEFAULT_SOURCE  // shut up gcc: implicit declaration of function ‘getloadavg’
#include "display.h"
#include <unistd.h>  // sysconf
#include <stdlib.h>  // getloadavg
#include <stdio.h>   // printf
#include <argp.h>
#include <string.h>  // strsep


const char* DEFAULT_INTERVAL[] = {
    "0.25,green,default",
    "0.5,colour166,default",
    "1.0,colour1,default",
};


struct load {
    double one_min;
    double five_min;
    double fifteen_min;
};


struct load
get_load(void) {
    signed int is_error;
    double loadavg[3];
    struct load load_s;

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


void
print_load(struct interval_display **entries, int count) {
    //unsigned int number_of_cpus = sysconf(_SC_NPROCESSORS_ONLN);
    //double cpus_d = (double) number_of_cpus;
    struct load loadavg = get_load();

    print_interval_item(entries, count, loadavg.one_min, " ");
    print_interval_item(entries, count, loadavg.five_min, " ");
    print_interval_item(entries, count, loadavg.fifteen_min, "");
}

/* ARGPARSE */

const char *argp_program_version =
  "load 0.1";

const char *argp_program_bug_address =
  "github.com/TomasTomecek/tmux-status-helpers/";

/* Program documentation. */
static char doc[] = "doc";

/* A description of the arguments we accept. */
static char args_doc[] = "[THRESHOLD_ENTRY [THRESHOLD_ENTRY...]]";

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

static struct argp argp = { options, parse_opt, args_doc, doc, 0, 0, 0 };

/* MAIN */

int
main(int argc, char **argv) {
    struct arguments arguments;
    arguments.args_count = 0;
    argp_parse(&argp, argc, argv, 0, 0, &arguments);

    int entries_count = arguments.args_count;
    struct interval_display **entries;
    if (entries_count == 0) {
        entries_count = (int)(sizeof(DEFAULT_INTERVAL) / sizeof(DEFAULT_INTERVAL[0]));
        entries = init_interval_list(DEFAULT_INTERVAL, entries_count);
    } else {
        entries = init_interval_list((const char**) arguments.args, entries_count);
    }
    print_load(entries, entries_count);

    return 0;
}
