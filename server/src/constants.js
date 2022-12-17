import pkg from '../package.json' assert { type: 'json' };

export const description = pkg.description;
export const gitUrl = `https://${pkg.repository.url}`;
export const license = pkg.license;
export const version = pkg.version;
