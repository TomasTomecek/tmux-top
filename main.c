#define _DEFAULT_SOURCE  // shut up gcc: implicit declaration of function ‘getloadavg’
#include <unistd.h>  // sysconf
#include <stdlib.h>  // getloadavg
#include <stdio.h>   // printf
#include <argp.h>


typedef struct load {
    double one_min;
    double five_min;
    double fifteen_min;
} load;


// how to display specific interval
typedef struct load_entry {
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

void
init_load_list() {
    load_entry load_list[3] = {
        { .load_d=0.05, .bg="default", .fg="green" },
        { .load_d=0.1, .bg="default", .fg="orange" },
        { .load_d=0.2, .bg="default", .fg="red" }
    };
}

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

int
main(int argc, char **argv) {
    print_load(get_load());
    argp_parse (0, argc, argv, 0, 0, 0);

    return 0;
}
