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

int cmp_func(const void * a, const void * b) {
    struct load_entry *a_s = (struct load_entry *) a;
    struct load_entry *b_s = (struct load_entry *) b;
    return (a_s->load_d < b_s->load_d) - (a_s->load_d > b_s->load_d);
}

struct load_entry**
init_load_list(char *args[], int count, int *entries_count) {
    *entries_count = 3;
    //printf("count = %d, entries_count = %d\n", count, *entries_count);
    if (count > 0) {
        *entries_count = count;
    }

    struct load_entry **entries = calloc(*entries_count, sizeof(struct load_entry*));

    if (count == 0) {
        entries[0] = (struct load_entry*) malloc(sizeof(struct load_entry));
        entries[0]->load_d = 0.25;
        entries[0]->bg = "default";
        entries[0]->fg = "green";
        entries[1] = (struct load_entry*) malloc(sizeof(struct load_entry));
        entries[1]->load_d = 0.5;
        entries[1]->bg = "default";
        entries[1]->fg = "colour166";  // solarized orange
        entries[2] = (struct load_entry*) malloc(sizeof(struct load_entry));
        entries[2]->load_d = 1.0;
        entries[2]->bg = "default";
        entries[2]->fg = "colour1"; // solarzied red
    } else {
        char *token, *string;
        char *triple[3];
        int loop_iteration;
        for (int i = 0; i < *entries_count; i++) {
            string = strdup(args[i]);
            //printf("%s\n", string);
            loop_iteration = 0;
            while ((token = strsep(&string, ",")) != NULL) {
                //printf("token[%d] = %s\n", loop_iteration, token);
                triple[loop_iteration] = strdup(token);
                ++loop_iteration;
            }
            entries[i] = (struct load_entry*) malloc(sizeof(struct load_entry));
            //printf("T[0] = %s, T[1] = %s, T[2] = %s", triple[0], triple[1], triple[2]);
            entries[i]->fg = strdup(triple[1]);
            //printf("fg = %s\n", entries[i]->fg);
            sscanf(triple[0], "%lf", &entries[i]->load_d);
            //printf("load = %.2f\n", entries[i]->load_d);
            if (loop_iteration == 3) {
                entries[i]->bg = strdup(triple[2]);
            } else {
                entries[i]->bg = "";
            }
            //printf("bg = %s\n", entries[i]->bg);
            //printf("fg=%s, bg=%s, %.2f\n", entries[i]->fg, entries[i]->bg, entries[i]->load_d);
        }
    }
    // sort array
    qsort(entries, *entries_count, sizeof(struct load_entry*), cmp_func);
    /*for (int i = 0; i < *entries_count; i++) {
        printf("[%d] fg=%s, bg=%s, %.2f\n", i, entries[i]->fg, entries[i]->bg, entries[i]->load_d);
    }*/
    return entries;
}

void
print_load_item(struct load_entry **entries, int count, double cpus_d, double load_to_display, char * suffix) {
    //
    // when you load/#cpus, you get an absolute load: lets not do that
    //double absolute_load = load_to_display / cpus_d * 100.0;
    for (int i = 0; i < count; i++) {
        // printf("\n[%d] load = %.2f, threshold = %.2f\n", i, load_to_display, entries[i]->load_d);
        if (load_to_display > entries[i]->load_d || i == count - 1) {
            if (strcmp(entries[i]->bg, "") == 0) {
                printf("#[fg=%s]%.2f#[fg=default]%s",
                       entries[i]->fg, load_to_display, suffix);
            } else {
                printf("#[bg=%s,fg=%s]%.2f#[fg=default,bg=default]%s",
                       entries[i]->bg, entries[i]->fg, load_to_display, suffix);
            }
            return;
        }
    }
}

void
print_load(struct load_entry **entries, int count) {
    unsigned int number_of_cpus = sysconf(_SC_NPROCESSORS_ONLN);
    double cpus_d = (double) number_of_cpus;
    struct load loadavg = get_load();

    print_load_item(entries, count, cpus_d, loadavg.one_min, " ");
    print_load_item(entries, count, cpus_d, loadavg.five_min, " ");
    print_load_item(entries, count, cpus_d, loadavg.fifteen_min, "");
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

    int entries_count = 0;
    struct load_entry **entries = init_load_list(arguments.args, arguments.args_count, &entries_count);
    print_load(entries, entries_count);

    return 0;
}
