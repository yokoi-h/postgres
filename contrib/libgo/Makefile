# contrib/libgo/Makefile

INCLUDEDIRS := -I.
INCLUDEDIRS += -I$(shell pg_config --includedir-server)
INCLUDEDIRS += -I$(shell pg_config --includedir)

LIBDIR := -L.
LIBDIR += -L$(shell pg_config --libdir)
LIBINSTALL = $(shell pg_config --pkglibdir)


callgo.so: callgo.c Makefile
                        gcc -fpic -o callgo.o -c callgo.c $(INCLUDEDIRS)
                        gcc -shared -o callgo.so callgo.o $(LIBDIR) -lpq -lm -lgo

install:
                        cp callgo.so $(LIBINSTALL)
                        cp libgo.so $(LIBINSTALL)
