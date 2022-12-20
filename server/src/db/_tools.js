export const defaultCols = (table, knex, opts = { id: true }) => {
  if (opts.id) {
    table.uuid('id', { primaryKey: true }).defaultTo(knex.raw('gen_random_uuid()'));
  }
  table.timestamps({
    useTimestamps: true,
    defaultToNow: true,
    useCamelCase: false,
  });
}