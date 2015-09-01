#include "postgres.h"
#include <string.h>
#include "fmgr.h"
#include "utils/geo_decls.h"
#include "libgo.h"

#ifdef PG_MODULE_MAGIC
PG_MODULE_MAGIC;
#endif

/* by value */
PG_FUNCTION_INFO_V1(get_jpy);

Datum
get_jpy(PG_FUNCTION_ARGS)
{
        int32   arg = PG_GETARG_INT32(0);
        float4  result = arg * jpy();
        PG_RETURN_FLOAT4(result);
}