import { Model } from 'objection';

// Base model class for application
class Base extends Model {
  // Gets allowed attributes from jsonSchema.
  static get allowAttrs() {
    return Object.keys(this.jsonSchema.properties);
  }

  // Removes any keys in `attrs` not present in the jsonSchema.
  static sanitize(attrs = {}) {
    return Object.keys(attrs)
      .filter(k => this.allowAttrs.includes(k))
      .reduce((obj, key) => Object.assign(obj, { [key]: attrs[key] }), {});
  }

  // Sanitizes attributes & creates a POJO.
  static new(attrs = {}) {
    const cleaned = this.sanitize(attrs);
    return this.fromJson(cleaned);
  }

  // Sanitizes attributes & executes an insert.
  static async create(attrs = {}) {
    const cleaned = this.sanitize(attrs);
    return this.query().insert(cleaned).returning('*');
  }

  // Sanitizes attributes & updates model instance.
  async update(attrs = {}) {
    const cleaned = this.$modelClass.sanitize(attrs);
    Object.assign(this, cleaned);
    return this.$query().update(this).returning('*');
  }
}

export default Base;