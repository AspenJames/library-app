import { Model } from 'objection';

// Base model class for application
class Base extends Model {
  // Gets allowed attributes from jsonSchema
  static get allowAttrs() {
    return Object.keys(this.jsonSchema.properties);
  }

  // Sanitizes input to strip out any keys not defined
  // in the jsonSchema before passing to fromJson
  static new(attrs = {}) {
    const cleanAttrs = Object.keys(attrs)
        .filter(k => this.allowAttrs.includes(k))
        .reduce((obj, key) => Object.assign(obj, {[key]: attrs[key]}), {});
    return this.fromJson(cleanAttrs);
  }

  static async create(params) {
    return this.query().insert(params).returning('*')
  }
}

export default Base;