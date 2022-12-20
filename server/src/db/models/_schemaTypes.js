// Merges & returns any supplied `opts` with the default boolean type.
export const boolType = (opts = {}) => Object.assign(
    { type: 'boolean', },
    opts
)

// Merges & returns any supplied `opts` with the default integer type.
export const intType = (opts = {}) => Object.assign(
    { type: 'number', },
    opts
)

// Merges & returns any supplied `opts` with the default string type.
export const stringType = (opts = {}) => Object.assign(
    { type: 'string', maxLength: 255, },
    opts
)
