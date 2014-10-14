#define _DEFAULT_SOURCE  // shut up gcc: implicit declaration of function ‘getloadavg’
#include <unistd.h>  // sysconf
#include <stdlib.h>  // getloadavg
#include <stdio.h>   // printf
#include <argp.h>
#include <string.h>  // strsep


struct load {
    double one_min;
    double five_min;
    double fifteen_min;
};


// how to display specific interval
struct load_entry {
    double load_d;
    char * bg;
    char * fg;
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

struct load_entry**
init_load_list(char *args[], int count) {
    int to_alloc = 3;
    if (count > 0) {
        to_alloc = count;
    }

    struct load_entry **entries = calloc(to_alloc, sizeof(struct load_entry*));

    if (count == 0) {
        entries[0] = (struct load_entry*) malloc(sizeof(struct load_entry));
        entries[0]->load_d = 0.05;
        entries[0]->bg = "default";
        entries[0]->fg = "green";
        entries[1] = (struct load_entry*) malloc(sizeof(struct load_entry));
        entries[1]->load_d = 0.1;
        entries[1]->bg = "default";
        entries[1]->fg = "orange";
        entries[2] = (struct load_entry*) malloc(sizeof(struct load_entry));
        entries[2]->load_d = 0.5;
        entries[2]->bg = "default";
        entries[2]->fg = "red";
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
        entries[i] = (struct load_entry*) malloc(sizeof(struct load_entry));
        entries[i]->fg = triple[1];
        sscanf(triple[0], "%lf", &entries[i]->load_d);
        if (loop_iteration == 3) {
            entries[i]->bg = triple[2];
        } else {
            entries[i]->bg = "";
        }
    }
    return entries;
}

void
print_load(struct **load_entry entries, int count) {
    unsigned int number_of_cpus = sysconf(_SC_NPROCESSORS_ONLN);
    double cpus_d = (double) number_of_cpus;
    struct loadavg = get_load();

    printf("%f\n", cpus_d);
    printf("%f\n", loadavg.one_min / cpus_d * 100.0 );

    for (int i = 0; i < count; i++) {
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

}

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
    struct arguments arguments;
    arguments.args_count = 0;
    argp_parse(&argp, argc, argv, 0, 0, &arguments);

    struct **load_entry entries = init_load_list(arguments.args, arguments.args_count);
    print_load(entries, 3);

    return 0;
}
