ALTER TABLE "transactions" DROP CONSTRAINT IF EXISTS "transactions_account_id_fkey";
ALTER TABLE "transactions" DROP CONSTRAINT IF EXISTS "transactions_operation_type_id_fkey";

DROP TABLE IF EXISTS "transactions";

DROP TABLE IF EXISTS "operation_types";
DROP INDEX IF EXISTS "accounts_account_id_idx";
DROP TABLE IF EXISTS "accounts";
